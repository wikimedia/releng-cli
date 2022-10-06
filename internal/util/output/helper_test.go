package output

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

func provideMap(name string) map[interface{}]interface{} {
	switch name {
	case "empty":
		return make(map[interface{}]interface{})
	case "test1.json":
		return mapFromFile("test1.json")
	case "simpleTable":
		objects := make(map[interface{}]interface{}, 2)
		objects["k1"] = "v1"
		objects["k2"] = "v2"
		return objects
	}
	panic("Unexpected map name")
}

func mapFromFile(file string) map[interface{}]interface{} {
	type ComplexObject struct {
		TopLevelString         string `json:"TopLevelString"`
		TopLevelMatchingInt    int    `json:"TopLevelMatchingInt"`
		TopLevelMatchingString string `json:"TopLevelMatchingString"`
		TopLevelStruct         struct {
			SecondLevelString     string `json:"SecondLevelString"`
			SecondLevelStructList []struct {
				ThirdLevelString string   `json:"ThirdLevelString"`
				ThirdLevelInt    int      `json:"ThirdLevelInt"`
				ThirdLevelList   []string `json:"ThirdLevelList"`
			} `json:"SecondLevelStructList"`
		} `json:"TopLevelStruct"`
	}
	type ComplexObjects map[string]ComplexObject

	// Get a filterable map from the 1 test file
	filterable := ComplexObjects{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = json.NewDecoder(strings.NewReader(string(bytes))).Decode(&filterable); err != nil {
		logrus.Fatal(err)
	}
	filterableMap := make(map[interface{}]interface{}, len(filterable))
	for key, value := range filterable {
		filterableMap[key] = value
	}

	return filterableMap
}

func emptyMap() map[interface{}]interface{} {
	return provideMap("empty")
}
