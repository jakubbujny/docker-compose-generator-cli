package yml_operator

import (
	"gopkg.in/yaml.v2"
	"strings"
)

func get(m map[interface{}]interface{}, path string) map[interface{}]interface{} {
	pathSplitted := strings.Split(path, ".")
	iterate := m
	for _,el := range pathSplitted {
		iterate = iterate[el].(map[interface{}]interface{})
	}
	return iterate
}

func set(m map[interface{}]interface{}, path string, set map[interface{}]interface{}) map[interface{}]interface{} {
	pathSplitted := strings.Split(path, ".")
	iterate := m
	for index,el := range pathSplitted {
		if index == len(pathSplitted) -1 {
			iterate[el] = set
		} else {
			iterate = iterate[el].(map[interface{}]interface{})
		}
	}
	return iterate
}

func AppendToYmlInSection(toAppend string, sourceYml string, appendPath string) (string, error) {
	sourceYmlParsed := make(map[interface{}]interface{})
	yaml.Unmarshal([]byte(sourceYml), &sourceYmlParsed)
	toAppendYmlParsed := make(map[interface{}]interface{})
	yaml.Unmarshal([]byte(toAppend), &toAppendYmlParsed)
	merged := get(sourceYmlParsed, appendPath)
	for k, v := range toAppendYmlParsed {
		merged[k] = v
	}
	set(sourceYmlParsed, appendPath, merged)
	toReturn,_ := yaml.Marshal(sourceYmlParsed)
	return string(toReturn), nil
}
