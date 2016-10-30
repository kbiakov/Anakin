package main

import (
	"fmt"
	"os"

	g "./generator"
	p "./parser"
)

var protoPath, host, port string

func main() {
	protoPath = os.Args[1]
	host = os.Args[2]
	port = os.Args[3]

	tplsPath := "src/templates/"
	tplServerMain := g.LoadTemplate(tplsPath + "server_main")
	tplServerMethod := g.LoadTemplate(tplsPath + "server_method")
	tplClientMain := g.LoadTemplate(tplsPath + "client_main")
	tplClientMethod := g.LoadTemplate(tplsPath + "client_method")

	proto, err := p.ParseProto(protoPath)
	if err != nil {
		panic(err)
	}
	if len(proto.Services) != 1 {
		panic(fmt.Sprintf("Found %d services, expected one", len(proto.Services)))
	}

	service := proto.Services[0]
	g.GenerateCode("server.go", &service, host, port, tplServerMain, tplServerMethod)
	g.GenerateCode("client.go", &service, host, port, tplClientMain, tplClientMethod)
}
