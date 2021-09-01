package tests

import (
	"github.com/k0kubun/pp"
	v1 "github.com/v8platform/protoc-gen-go-ras/tests/gen/ras/client/v1"
	v13 "github.com/v8platform/protoc-gen-go-ras/tests/gen/ras/messages/v1"
	v12 "github.com/v8platform/protoc-gen-go-ras/tests/gen/ras/protocol/v1"
	"testing"
)

func TestNewClientService(t *testing.T) {
	client := v1.NewClientService("srv-uk-app31:1545")
	_, err := client.Negotiate(&v12.NegotiateMessage{
		Magic:    475223888,
		Protocol: 256,
		Version:  256,
	})
	if err != nil {
		pp.Fatal(err)
	}

	connectAck, err := client.Connect(&v12.ConnectMessage{})
	if err != nil {
		pp.Fatal(err)
	}

	pp.Println(connectAck.String())
	EndpointOpenAck, err := client.EndpointOpen(&v12.EndpointOpen{
		Service: "v8.service.Admin.Cluster",
		Version: "10.0",
	})
	if err != nil {
		pp.Fatal(err)
	}

	pp.Println(EndpointOpenAck.String())

	endpoint, err := client.NewEndpoint(EndpointOpenAck)
	if err != nil {
		pp.Fatal(err)
	}

	pp.Println(endpoint)

	endpointService := v1.NewEndpointService(client, endpoint)

	clustersService := v1.NewClustersService(endpointService)
	clusters, err := clustersService.GetClusters(&v13.GetClustersRequest{})
	if err != nil {
		pp.Fatal(err)
	}

	pp.Println(clusters.String())

	clusterInfo, err := clustersService.GetClusterInfo(&v13.GetClusterInfoRequest{ClusterId: clusters.Clusters[0].Uuid})
	if err != nil {
		pp.Fatal(err)
	}

	pp.Println(clusterInfo.ClusterInfo.String())
	auth := v1.NewAuthService(endpointService)
	_, err = auth.AuthenticateCluster(&v13.ClusterAuthenticateRequest{ClusterId: clusters.Clusters[0].Uuid})
	if err != nil {
		pp.Fatalln("auth", err)
	}

	infobasesService := v1.NewInfobasesService(endpointService)
	resp, err := infobasesService.GetShortInfobases(&v13.GetInfobasesShortRequest{ClusterId: clusters.Clusters[0].Uuid})
	if err != nil {
		pp.Fatalln("sessions", err)
	}
	pp.Println(resp.Infobases)

	// sessionService := v1.NewSessionsService(endpointService)
	// sessions, err := sessionService.GetSessions(&v13.GetSessionsRequest{ClusterId: clusters.Clusters[0].Uuid})
	// if err != nil {
	// 	pp.Fatalln("sessions", err)
	// }
	//
	// pp.Println(sessions.Sessions)

}
