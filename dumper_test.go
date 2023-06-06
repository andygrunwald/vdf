package vdf

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewDumper(t *testing.T) {
	testMap1 := map[string]interface{}{"map1": map[string]interface{}{"key1": "test1", "key2": "test2",
		"map2": map[string]interface{}{"key3": "test3", "key4": "test4"}},
		"map3": map[string]interface{}{"key5": "test5", "key6": "test6"}}
	testMap2 := map[string]interface{}{"map1": map[string]interface{}{"key1": "test1", "key2": "test2",
		"map2": []string{"key5", "key6"}},
		"map3": map[string]interface{}{"key5": "test5", "key6": "test6"}}
	vdfString, err := newDumper(testMap1)
	fmt.Println(vdfString, err)
	reader := strings.NewReader(vdfString)
	p := NewParser(reader)
	m, err := p.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
	vdfString, err = newDumper(testMap2)
	fmt.Println(vdfString, err)
}
