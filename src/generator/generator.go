package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	p "../parser"
)

// LoadTemplate loads template by path and returns it.
func LoadTemplate(path string) *template.Template {
	tpl, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot load template %s: %s", path, err))
	}
	return template.Must(template.New(path).Parse(string(tpl)))
}

// MethodsToPlaintext executes method template by inner rpc methods and returns as plaintext.
func MethodsToPlaintext(serviceName string, rpcs []p.Rpc, methodTpl *template.Template) string {
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

// GenerateCode generates code by specified path using provided main & method templates.
func GenerateCode(filename string, service *p.Service, host string, port string, mainTpl *template.Template, methodTpl *template.Template) {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Could not create %s: %s", filename, err))
	}
	defer file.Close()

	methodsPlaintext := MethodsToPlaintext(service.Name, service.Rpcs, methodTpl)

	mainTpl.Execute(file, map[string]string{
		"Service": service.Name,
		"Methods": methodsPlaintext,
		"Host":    host,
		"Port":    port,
	})
}
