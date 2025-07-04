package vdf

import (
	"errors"
	"fmt"
	"strings"
)

func newDumper(vdfMap map[string]interface{}) (string, error) {
	var outBuilder strings.Builder
	err := recursiveMap(vdfMap, 0, &outBuilder)
	if err != nil {
		return "", err
	}
	return outBuilder.String(), nil
}

func recursiveMap(m map[string]interface{}, depth int, outBuilder *strings.Builder) error {
	for key, value := range m {
		switch valueType := value.(type) {
		case map[string]interface{}:
			outBuilder.WriteString(fmt.Sprintf("%s\"%s\"\n%s{\n", strings.Repeat("\t", depth), key, strings.Repeat("\t", depth)))
			err := recursiveMap(valueType, depth+1, outBuilder)
			if err != nil {
				return err
			}
			outBuilder.WriteString(fmt.Sprintf("%s}\n", strings.Repeat("\t", depth)))
		case string:
			outBuilder.WriteString(fmt.Sprintf("%s\"%s\"\t\t\"%s\"\n", strings.Repeat("\t", depth), key, value))
		default:
			return errors.New("Unsupported value type")
		}
	}
	return nil
}
