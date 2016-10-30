package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	entityService = "service "
	entityMessage = "message "
	entityRpc     = "rpc "
	entityOpen    = " {"
	entityClose   = "}"

	featureBrOpen  = " ("
	featureBrClose = ") "
	featureReturns = "returns"
)

// Proto describes structure of proto-file (servers & messages).
type Proto struct {
	Services []Service `json:"services"`
	Messages []string  `json:"messages"`
}

// Service has name & rpc-methods.
type Service struct {
	Name string `json:"name"`
	Rpcs []Rpc  `json:"rpcs"`
}

// Rpc is combination of name, request & response.
type Rpc struct {
	Name string `json:"name"`
	Req  string `json:"req"`
	Res  string `json:"res"`
}

// ParseProto parses proto-file by presented path.
func ParseProto(path string) (*Proto, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	proto := new(Proto)
	isParseService := false

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()

		if isParseService {
			if isStopParseEntity(line) {
				isParseService = false
			} else if isFoundRpc(line) {
				rpc := parseRpc(line)
				last := len(proto.Services) - 1
				rpcs := proto.Services[last].Rpcs
				proto.Services[last].Rpcs = append(rpcs, *rpc)
			}
		} else if ok, serviceName := searchService(line); ok {
			service := Service{Name: serviceName}
			proto.Services = append(proto.Services, service)
			isParseService = true
		} else if ok, message := searchMessage(line); ok {
			proto.Messages = append(proto.Messages, message)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return proto, nil
}

func searchSurrounded(openFeature string, closeFeature string, str string) (bool, string) {
	fmtOpenFeature := screenBrackets(openFeature)
	fmtCloseFeature := screenBrackets(closeFeature)
	regex := fmt.Sprintf("(?:%s)(.*)(?:%s)", fmtOpenFeature, fmtCloseFeature)
	trimmedStr := strings.TrimSpace(str)
	if re := regexp.MustCompile(regex); re.MatchString(trimmedStr) {
		return true, re.FindStringSubmatch(trimmedStr)[1]
	}
	return false, ""
}

func searchService(str string) (bool, string) {
	return searchSurrounded(entityService, entityOpen, str)
}

func searchMessage(str string) (bool, string) {
	return searchSurrounded(entityMessage, entityOpen, str)
}

func selectSurrounded(openFeature string, closeFeature string, str string) string {
	_, surrounded := searchSurrounded(openFeature, closeFeature, str)
	return surrounded
}

func isFoundRpc(str string) bool {
	return strings.HasPrefix(strings.TrimSpace(str), entityRpc)
}

func isStopParseEntity(str string) bool {
	return strings.TrimSpace(str) == entityClose
}

func parseRpc(str string) *Rpc {
	req := selectSurrounded(featureBrOpen, featureBrClose+featureReturns, str)
	res := selectSurrounded(featureReturns+featureBrOpen, featureBrClose, str)
	name := selectSurrounded(entityRpc, featureBrOpen+req+featureBrClose+featureReturns, str)

	return &Rpc{
		Name: name,
		Req:  req,
		Res:  res,
	}
}

func screenBrackets(str string) string {
	return strings.Replace(strings.Replace(str, "(", "\\(", -1), ")", "\\)", -1)
}
