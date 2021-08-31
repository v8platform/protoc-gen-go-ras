package tests

import (
	"github.com/k0kubun/pp"
	v1 "github.com/v8platform/protoc-gen-go-ras/tests/gen/ras/client/v1"
	v12 "github.com/v8platform/protoc-gen-go-ras/tests/gen/ras/protocol/v1"
	"testing"
)

func TestNewClientService(t *testing.T) {
	client := v1.NewClientService("srv-uk-app10:1545")
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

	pp.Println(connectAck)
	EndpointOpenAck, err := client.EndpointOpen(&v12.EndpointOpen{
		Service: "v8.service.Admin.Cluster",
		Version: "10.0",
	})
	if err != nil {
		pp.Fatal(err)
	}

	pp.Println(EndpointOpenAck.String())
	pp.Println(EndpointOpenAck)
}
