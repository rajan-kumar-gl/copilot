package testhelpers

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	copilotsnapshot "code.cloudfoundry.org/copilot/snapshot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	mcp "istio.io/api/mcp/v1alpha1"
	"istio.io/api/networking/v1alpha3"
	"istio.io/istio/pkg/mcp/sink"
)

type MockMCPUpdater struct {
	changesMux sync.Mutex
	objects    map[string][]*sink.Object
}

func (m *MockMCPUpdater) Apply(c *sink.Change) error {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	m.objects[c.Collection] = c.Objects
	return nil
}

func (m *MockMCPUpdater) GetAllObjectNames() map[string][]string {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	allObjectNames := map[string][]string{}
	for obType, objects := range m.objects {
		for _, object := range objects {
			allObjectNames[obType] = append(allObjectNames[obType], object.Metadata.Name)
		}
	}
	return allObjectNames
}

func (m *MockMCPUpdater) GetAllVirtualServices() []*v1alpha3.VirtualService {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	allVirtualServices := []*v1alpha3.VirtualService{}
	for _, vs := range m.objects[copilotsnapshot.VirtualServiceTypeURL] {
		r := vs.Body.(interface{}).(*v1alpha3.VirtualService)
		allVirtualServices = append(allVirtualServices, r)
	}

	return allVirtualServices
}

func (m *MockMCPUpdater) GetAllDestinationRules() []*v1alpha3.DestinationRule {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	allDestinationRules := []*v1alpha3.DestinationRule{}
	for _, o := range m.objects[copilotsnapshot.DestinationRuleTypeURL] {
		r := o.Body.(interface{}).(*v1alpha3.DestinationRule)
		allDestinationRules = append(allDestinationRules, r)
	}

	return allDestinationRules
}

func (m *MockMCPUpdater) GetAllGateways() []*v1alpha3.Gateway {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	allGateways := []*v1alpha3.Gateway{}
	for _, o := range m.objects[copilotsnapshot.GatewayTypeURL] {
		r := o.Body.(interface{}).(*v1alpha3.Gateway)
		allGateways = append(allGateways, r)
	}

	return allGateways
}

func (m *MockMCPUpdater) GetAllSidecars() []*v1alpha3.Sidecar {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	allSidecars := []*v1alpha3.Sidecar{}
	for _, o := range m.objects[copilotsnapshot.SidecarTypeURL] {
		r := o.Body.(interface{}).(*v1alpha3.Sidecar)
		allSidecars = append(allSidecars, r)
	}

	return allSidecars
}

func (m *MockMCPUpdater) GetAllServiceEntries() []*v1alpha3.ServiceEntry {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	allServiceEntries := []*v1alpha3.ServiceEntry{}
	for _, o := range m.objects[copilotsnapshot.ServiceEntryTypeURL] {
		r := o.Body.(interface{}).(*v1alpha3.ServiceEntry)
		allServiceEntries = append(allServiceEntries, r)
	}

	return allServiceEntries
}

func (m *MockMCPUpdater) GetAllMessageNames() []string {
	m.changesMux.Lock()
	defer m.changesMux.Unlock()
	typeURLs := []string{}
	for k, _ := range m.objects {
		typeURLs = append(typeURLs, k)
	}
	return typeURLs
}

// MetricReporter is used to report metrics for an MCP client.
type MockMetricReporter struct{}

func (m *MockMetricReporter) Close() error                                                      { return nil }
func (m *MockMetricReporter) RecordSendError(err error, code codes.Code)                        {}
func (m *MockMetricReporter) RecordRecvError(err error, code codes.Code)                        {}
func (m *MockMetricReporter) RecordRequestSize(collection string, connectionID int64, size int) {}
func (m *MockMetricReporter) RecordRequestAck(collection string, connectionID int64)            {}
func (m *MockMetricReporter) RecordRequestNack(collection string, connectionID int64, code codes.Code) {
}

func (m *MockMetricReporter) SetStreamCount(clients int64) {}
func (m *MockMetricReporter) RecordStreamCreateSuccess()   {}

type MockPilotMCPClient struct {
	ctx        context.Context
	client     *sink.Client
	cancelFunc func()
	conn       *grpc.ClientConn
	*MockMCPUpdater
}

func (m *MockPilotMCPClient) Close() error {
	m.cancelFunc()
	err := m.conn.Close()
	return err
}

func NewMockPilotMCPClient(tlsConfig *tls.Config, serverAddr string) (*MockPilotMCPClient, error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		grpc.WithTimeout(5 * time.Second),
	}

	conn, err := grpc.DialContext(context.Background(), serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	svcClient := mcp.NewResourceSourceClient(conn)
	mockUpdater := &MockMCPUpdater{objects: make(map[string][]*sink.Object)}
	mockReporter := &MockMetricReporter{}

	typeURLs := []string{
		copilotsnapshot.AuthenticationPolicyTypeURL,
		copilotsnapshot.AuthenticationMeshPolicyTypeURL,
		copilotsnapshot.AuthorizationPolicyTypeURL,
		copilotsnapshot.GatewayTypeURL,
		copilotsnapshot.VirtualServiceTypeURL,
		copilotsnapshot.DestinationRuleTypeURL,
		copilotsnapshot.ServiceEntryTypeURL,
		copilotsnapshot.SidecarTypeURL,
		copilotsnapshot.EnvoyFilterTypeURL,
		copilotsnapshot.HTTPAPISpecTypeURL,
		copilotsnapshot.HTTPAPISpecBindingTypeURL,
		copilotsnapshot.QuotaSpecTypeURL,
		copilotsnapshot.QuotaSpecBindingTypeURL,
		copilotsnapshot.PolicyTypeURL,
		copilotsnapshot.MeshPolicyTypeURL,
		copilotsnapshot.ServiceRoleTypeURL,
		copilotsnapshot.ServiceRoleBindingTypeURL,
		copilotsnapshot.RbacConfigTypeURL,
	}
	collectionOptions := sink.CollectionOptionsFromSlice(typeURLs)

	sinkOptions := &sink.Options{
		CollectionOptions: collectionOptions,
		Updater:           mockUpdater,
		ID:                "",
		Metadata:          nil,
		Reporter:          mockReporter,
	}
	cl := sink.NewClient(svcClient, sinkOptions)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go cl.Run(ctx)

	return &MockPilotMCPClient{
		ctx:            ctx,
		client:         cl,
		cancelFunc:     cancelFunc,
		conn:           conn,
		MockMCPUpdater: mockUpdater,
	}, nil
}
