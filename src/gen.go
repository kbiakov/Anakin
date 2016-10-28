package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	p "./parser"
)

var protoPath, host, port string

func loadTemplate(tplName string) *template.Template {
	tpl, err := ioutil.ReadFile(fmt.Sprintf("%s%s", "src/templates/", tplName))
	if err != nil {
		panic(fmt.Sprintf("Cannot load template %s: %s", tplName, err))
	}
	return template.Must(template.New(tplName).Parse(string(tpl)))
}

func methodsToPlaintext(serviceName string, rpcs []p.Rpc, methodTpl *template.Template) string {
	var methods bytes.Buffer

	for i, rpc := range rpcs {
		var method bytes.Buffer

		methodTpl.Execute(&method, map[string]string{
			"Service":  serviceName,
			"Method":   rpc.Name,
			"Request":  rpc.Req,
			"Response": rpc.Res,
		})

		methods.Write(method.Bytes())

		if i < len(rpcs)-1 {
			methods.Write([]byte("\n"))
		}
	}
	return methods.String()
}

func execTemplate(filename string, service *p.Service, mainTpl *template.Template, methodTpl *template.Template) {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not create %s: %s", filename, err))
	}
	defer file.Close()

	methodsPlaintext := methodsToPlaintext(service.Name, service.Rpcs, methodTpl)

	mainTpl.Execute(file, map[string]string{
		"Service": service.Name,
		"Methods": methodsPlaintext,
		"Host":    host,
		"Port":    port,
	})
}

func main() {
	protoPath = os.Args[1]
	host = os.Args[2]
	port = os.Args[3]

	tplServerMain := loadTemplate("server_main")
	tplServerMethod := loadTemplate("server_method")
	tplClientMain := loadTemplate("client_main")
	tplClientMethod := loadTemplate("client_method")

	proto, err := p.ParseProto(protoPath)
	if err != nil {
		panic(err)
	}
	if len(proto.Services) != 1 {
		panic(fmt.Sprintf("Found %d services, expected one", len(proto.Services)))
	}

	service := proto.Services[0]
	execTemplate("server.go", &service, tplServerMain, tplServerMethod)
	execTemplate("client.go", &service, tplClientMain, tplClientMethod)
}