package logparser

import (
	"github.com/ivan-uskov/go-course/log_parser/datastruct"
	"regexp"
	"errors"
)

func ParseClientInfo(clientInfoAsString string) (datastruct.ClientInfo, error) {
	re := regexp.MustCompile(REGEX_STRING)
	names := re.SubexpNames()
	values := re.FindAllStringSubmatch(clientInfoAsString, -1)
	if len(values) == 0 {
		return datastruct.ClientInfo{}, errors.New("Invalid log string")
	}

	mappedValues := map[string]string{}
	for key, val := range values[0] {
		mappedValues[names[key]] = val
	}

	return createClientInfoFromMap(mappedValues), nil
}

func createClientInfoFromMap(mappedValues map[string]string) datastruct.ClientInfo {
	return  datastruct.ClientInfo{
		Ip: mappedValues[IP_ADDRESS],
		Date: mappedValues[DATETIME],
		Browser: mappedValues[USERAGENT],
	}
}
