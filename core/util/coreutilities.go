package util

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/sonyjop/goweave/core"
)

func ParseNodeURI_deprecated(uri string) (core.Node, error) {
	var componentName, path string
	parts := strings.Split(uri, "://")

	validateUriPart := func(input, pattern string) bool {
		match, err := regexp.MatchString(pattern, input)
		if err != nil {
			fmt.Println("Error compiling regex:", err)
			return false
		}
		return match
	}
	if len(parts) < 2 || strings.TrimSpace(parts[0]) == "" {
		return core.Node{}, errors.New("invalid uri format, unable to identify component name")
	}

	componentName = parts[0]
	remainder := parts[1]
	queryMap := make(map[string]string)

	if !strings.Contains(remainder, "?") {
		path = remainder

	} else {
		parts = strings.Split(remainder, "?")
		if strings.TrimSpace(parts[0]) == "" || strings.TrimSpace(parts[1]) == "" {
			return core.Node{}, errors.New("invalid uri format, path/query is empty")
		}
		path = parts[0]
		remainder = parts[1]

		parts = strings.Split(remainder, "&")

		for _, val := range parts {
			keyValArr := strings.Split(val, "=")
			if len(keyValArr) < 2 {
				return core.Node{}, errors.New("invalid uri format, query format invalid")
			}
			if !validateUriPart(keyValArr[0], "[a-zA-Z][A-Za-z0-9]+") {
				return core.Node{}, errors.New("invalid key in query")
			}
			queryMap[keyValArr[0]] = keyValArr[1]
		}
	}
	if !validateUriPart(componentName, "[a-zA-Z][A-Za-z0-9]+") {
		return core.Node{}, errors.New("invalid component name")
	}
	if !validateUriPart(path, "[a-zA-Z][A-Za-z0-9]+") {
		return core.Node{}, errors.New("invalid path")
	}
	return core.Node{
		ComponentName: componentName,
		Path:          path,
		Query:         queryMap,
	}, nil
}
func ParseNodeURI(uri string) (core.Node, error) {
	var componentName, path, query string
	var queryMap map[string]string
	//parts := strings.Split(uri, "://")

	schemePattern := "([A-Za-z][A-Za-z0-9]{1,}:\\/\\/)"
	pathPattern := "([A-Za-z][A-Za-z0-9@\\/]{1,}(:[0-9]{1,})?)[\\/A-Za-z0-9]*"
	queryPattern := "(\\?[A-Za-z][A-Za-z0-9]{0,}=.*)?"

	urlPattern := schemePattern + pathPattern + queryPattern

	urlRegex := regexp.MustCompile(urlPattern)
	if !urlRegex.MatchString(uri) {
		return core.Node{}, errors.New("The given URL is not valid")
	}
	matchGroups := urlRegex.FindAllStringSubmatch(uri, -1)
	//matchGroups first element will have 5 groups
	// [0] - have the full uri
	// [1] - have the scheme
	// [2] - have path
	// [3] - port, if any
	// [4] - query part if any, this includes the ? also
	componentName, _, _ = strings.Cut(matchGroups[0][1], ":")
	path = matchGroups[0][2]
	query = matchGroups[0][4]
	if query != "" {
		queryMap = make(map[string]string)
		query = strings.Replace(query, "?", "", 1)
		keyValuePairs := strings.Split(query, "&")
		for _, keyValuePair := range keyValuePairs {
			parts := strings.Split(keyValuePair, "=")

			if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
				return core.Node{}, errors.New("query parameters malformed. key or value missing : " + parts[0] + " : " + parts[1])
			}
			queryMap[parts[0]] = parts[1]
		}

	}
	return core.Node{
		ComponentName: componentName,
		Path:          path,
		Query:         queryMap,
	}, nil
}
