package vdf_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/andygrunwald/vdf"
)

// Ensure the parser can parse strings into Statement ASTs.
func TestParser_ParseStatement(t *testing.T) {
	var tests = []struct {
		s   string
		m   map[string]interface{}
		err error
	}{
		// Single field statement
		{
			s: `"Example"
{
	"TimeNextStatsReport"      "1234567890"
	"ContentStatsID"           "-7123456789012345678"
}`,
			m: map[string]interface{}{
				"Example": map[string]interface{}{
					"TimeNextStatsReport": "1234567890",
					"ContentStatsID":      "-7123456789012345678",
				},
			},
		},
		{
			s: `"Root"
{
 "attr1"       "hey-ho"
 "attr2"       "ho-hey"
 "map1"
 {
   "foo"       "Q79v5tbar"
 }
 "data"
 {
   "val"       "1"
   "map"       "2"
   "player"    "3"
 }
}`,
			m: map[string]interface{}{
				"Root": map[string]interface{}{
					"attr1": "hey-ho",
					"attr2": "ho-hey",
					"map1": map[string]interface{}{
						"foo": "Q79v5tbar",
					},
					"data": map[string]interface{}{
						"val":    "1",
						"map":    "2",
						"player": "3",
					},
				},
			},
		},
		{
			s: `"Root"
{
 attr1       "hey-ho"
 "attr2"       ho
 "map1"
 {
   "foo"       "Q79v5tbar"
 }
 "data"
 {
   "v\\al"       "1"
   "map"       "2"
   "pl\"ayer"    "3"
 }
}`,
			m: map[string]interface{}{
				"Root": map[string]interface{}{
					"attr1": "hey-ho",
					"attr2": "ho",
					"map1": map[string]interface{}{
						"foo": "Q79v5tbar",
					},
					"data": map[string]interface{}{
						"v\\al":    "1",
						"map":      "2",
						"pl\"ayer": "3",
					},
				},
			},
		},
		{
			s: `"Root"
{
 attr1       "hey-ho"
 "attr2"       ho
 "map1"
 {
   // This is a comment
   "foo"       "Q79v5tbar"
 }
 "data"
 {
   "v\\al"       "1"
   "map"       "2"
   "pl\"ayer"    "3"
 }
}`,
			m: map[string]interface{}{
				"Root": map[string]interface{}{
					"attr1": "hey-ho",
					"attr2": "ho",
					"map1": map[string]interface{}{
						"foo": "Q79v5tbar",
					},
					"data": map[string]interface{}{
						"v\\al":    "1",
						"map":      "2",
						"pl\"ayer": "3",
					},
				},
			},
		},
		{
			s: `"Root"
{
 attr1       "hey-ho"
 "attr2"       ho
 "map1"
 {
   // Comment line first
   // Comment line second
   "foo"       "Q79v5tbar"
   // Comment line third
 }
 "data"
 {
   "v\\al"       "1"
   "map"       "2"
   "pl\"ayer"    "3"
 }
}`,
			m: map[string]interface{}{
				"Root": map[string]interface{}{
					"attr1": "hey-ho",
					"attr2": "ho",
					"map1": map[string]interface{}{
						"foo": "Q79v5tbar",
					},
					"data": map[string]interface{}{
						"v\\al":    "1",
						"map":      "2",
						"pl\"ayer": "3",
					},
				},
			},
		},
		{
			s: `// Root comment line
"Root"
{
 attr1       "hey-ho"
 "attr2"       ho
}`,
			m: map[string]interface{}{
				"Root": map[string]interface{}{
					"attr1": "hey-ho",
					"attr2": "ho",
				},
			},
		},
	}

	for i, tt := range tests {
		m, err := vdf.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, (err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == nil && !reflect.DeepEqual(tt.m, m) {
			t.Errorf("%d. %q\n\nparse mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.m, m)
		}
	}
}
