package logparser

import (
	"github.com/mssola/user_agent"
	"regexp"
	"errors"
)

func ParseClientInfo(clientInfoAsString string) (map[string]string, error) {
	re := regexp.MustCompile(REGEX_STRING)
	names := re.SubexpNames()
	values := re.FindAllStringSubmatch(clientInfoAsString, -1)
	if len(values) == 0 {
		return map[string]string{}, errors.New("Invalid log string")
	}

	mappedValues := map[string]string{}
	for key, val := range values[0] {
		mappedValues[names[key]] = val
	}

	return mappedValues, nil
}

func ParseUserAgent(useragent string) (string, string) {
	ua := user_agent.New(useragent)
	return ua.Browser()
}
