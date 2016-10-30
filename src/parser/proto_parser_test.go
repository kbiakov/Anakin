package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// TestParseProto tests ParseProto with valid path.
func TestParseProto(t *testing.T) {
	examplePath := "test.proto"
	exampleRpc := createExampleRpc()
	exampleProto := Proto{
		Services: []Service{{
			Name: "Greeter",
			Rpcs: []Rpc{
				exampleRpc,
			},
		}},
		Messages: []string{
			exampleRpc.Req,
			exampleRpc.Res,
		},
	}

	proto, err := ParseProto(examplePath)
	if err != nil {
		t.Fatalf("Error reading proto-file by path %s: %s", examplePath, err.Error())
	}

	if !reflect.DeepEqual(*proto, exampleProto) {
		protoJson, _ := json.Marshal(&proto)
		exampleProtoJson, _ := json.Marshal(&exampleProto)
		t.Errorf("Expected %s, found %s", string(exampleProtoJson), string(protoJson))
	}
}

// TestParseProtoWithInvalidPath tests ParseProto with invalid path.
func TestParseProtoWithInvalidPath(t *testing.T) {
	if _, err := ParseProto("abc.proto"); err == nil {
		t.Fail()
	}
}

// TestSearchSurrounded tests searchSurrounded when surrounded string found.
func TestSearchSurrounded(t *testing.T) {
	ok, ss := searchSurrounded("(", ")", "(abc)")
	if !ok || ss != "abc" {
		t.Fail()
	}
}

// TestSearchSurroundedNotFound tests searchSurrounded when surrounded string not found.
func TestSearchSurroundedNotFound(t *testing.T) {
	ok1, ss1 := searchSurrounded("(", ")", "(abc")
	if ok1 || ss1 != "" {
		t.Fail()
	}

	ok2, ss2 := searchSurrounded("(", ")", "abc)")
	if ok2 || ss2 != "" {
		t.Fail()
	}

	ok3, ss3 := searchSurrounded("(", ")", "abc")
	if ok3 || ss3 != "" {
		t.Fail()
	}
}

// TestSearchService tests searchService.
// (Other cases covered by previous tests.)
func TestSearchService(t *testing.T) {
	ok, sName := searchService("service abc {")
	if !ok || sName != "abc" {
		t.Fail()
	}
}

// TestSearchMessage tests searchMessage.
// (Other cases covered by previous tests.)
func TestSearchMessage(t *testing.T) {
	ok, sName := searchMessage("message abc {")
	if !ok || sName != "abc" {
		t.Fail()
	}
}

// TestSelectSurrounded tests selectSurrounded.
// (Other cases covered by previous tests.)
func TestSelectSurrounded(t *testing.T) {
	ss := selectSurrounded("(", ")", "(abc)")
	if ss != "abc" {
		t.Fail()
	}
}

// TestIsFoundRpc tests isFoundRpc.
func TestIsFoundRpc(t *testing.T) {
	if !isFoundRpc("rpc abc") {
		t.Fail()
	}
	if isFoundRpc("rpcabc") {
		t.Fail()
	}
}

// TestIsStopParseEntity tests isStopParseEntity.
func TestIsStopParseEntity(t *testing.T) {
	if !isStopParseEntity("}") {
		t.Fail()
	}
	if !isStopParseEntity("  }  ") {
		t.Fail()
	}
	if isStopParseEntity("rpcabc") {
		t.Fail()
	}
}

// TestParseRpc tests parseRpc.
func TestParseRpc(t *testing.T) {
	exampleRpc := createExampleRpc()

	rpc := parseRpc(fmt.Sprintf("rpc %s (%s) returns (%s) {}",
		exampleRpc.Name, exampleRpc.Req, exampleRpc.Res))

	if !reflect.DeepEqual(*rpc, exampleRpc) {
		protoJson, _ := json.Marshal(&rpc)
		exampleProtoJson, _ := json.Marshal(&exampleRpc)
		t.Errorf("Expected %s, found %s", string(exampleProtoJson), string(protoJson))
	}
}

func createExampleRpc() Rpc {
	return Rpc{
		Name: "SayHello",
		Req:  "HelloRequest",
		Res:  "HelloReply",
	}
}
