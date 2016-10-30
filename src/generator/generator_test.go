package generator

import (
	"fmt"
	"os"
	"testing"

	p "../parser"
	t "text/template"
)

// TestLoadTemplate tests LoadTemplate by name for existing template.
func TestLoadTemplate(t *testing.T) {
	testLoadTemplateWithPath := func(t *testing.T, tplName string) {
		if loadTestTemplate(tplName) == nil {
			t.Errorf("Could not load %s template", tplName)
		}
	}

	testLoadTemplateWithPath(t, "server_main")
	testLoadTemplateWithPath(t, "server_method")
	testLoadTemplateWithPath(t, "client_main")
	testLoadTemplateWithPath(t, "client_method")
}

func loadTestTemplate(tplName string) *t.Template {
	return LoadTemplate(fmt.Sprintf("%s%s", "../templates/", tplName))
}

// TestLoadTemplateNotExisting tests LoadTemplate for not existing template.
func TestLoadTemplateNotExisting(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Cannot recover after trying to load not exisitng tmeplate")
		}
	}()

	LoadTemplate("not_existing")

	t.Fail()
}

// TestMethodsToPlaintext tests MethodsToPlaintext.
func TestMethodsToPlaintext(t *testing.T) {
	serviceName := "Greeter"
	rpc := createExampleRpc()
	testMethodsToPlaintextForSide := func(side string, expectedMethodPlaintext string) {
		tpl := loadTestTemplate(fmt.Sprintf("%s%s", side, "_method"))
		methodPlaintext := MethodsToPlaintext(serviceName, []p.Rpc{rpc}, tpl)
		if methodPlaintext != expectedMethodPlaintext {
			t.Errorf("Expected %+q, found %+q", expectedMethodPlaintext, methodPlaintext)
		}
	}

	testMethodsToPlaintextForSide("server", fmt.Sprintf(
		"// TODO: ADD IMPLEMENTATION BELOW\n"+
			"func (s *%sServer) %s(ctx context.Context, in *pb.%s) (*pb.%s, error) {\n"+
			"    return &pb.%s{}, nil\n"+
			"}",
		serviceName, rpc.Name, rpc.Req, rpc.Res, rpc.Res))

	testMethodsToPlaintextForSide("client", fmt.Sprintf(
		"// TODO: ADD ADDITIONAL INFO BELOW\n"+
			"func %s(req *pb.%s) (*pb.%s, error) {\n"+
			"    c := Get%sClientInstance()\n"+
			"    res := c.%s(context.Background(), &req)\n"+
			"    return res\n"+
			"}",
		rpc.Name, rpc.Req, rpc.Res, serviceName, rpc.Name))
}

// TestGenerateCode tests GenerateCode.
func TestGenerateCode(t *testing.T) {
	tplServerMain := loadTestTemplate("server_main")
	tplServerMethod := loadTestTemplate("server_method")
	tplClientMain := loadTestTemplate("client_main")
	tplClientMethod := loadTestTemplate("client_method")

	proto, err := p.ParseProto("../parser/test.proto")
	if err != nil {
		t.Errorf("Error parsing proto-file: %s", err.Error())
	}
	if len(proto.Services) != 1 {
		t.Errorf("Found %d services, expected one", len(proto.Services))
	}

	service := proto.Services[0]
	host := "localhost"
	port := "50051"

	exampleServerPath := "example_server.go"
	GenerateCode(exampleServerPath, &service, host, port, tplServerMain, tplServerMethod)
	if _, err := os.Stat(exampleServerPath); os.IsNotExist(err) {
		t.Errorf("File %s not found", exampleServerPath)
	}
	os.Remove(exampleServerPath)

	exampleClientPath := "example_client.go"
	GenerateCode(exampleClientPath, &service, host, port, tplClientMain, tplClientMethod)
	if _, err := os.Stat(exampleClientPath); os.IsNotExist(err) {
		t.Errorf("File %s not found", exampleClientPath)
	}
	os.Remove(exampleClientPath)
}

func createExampleRpc() p.Rpc {
	return p.Rpc{
		Name: "SayHello",
		Req:  "HelloRequest",
		Res:  "HelloReply",
	}
}
