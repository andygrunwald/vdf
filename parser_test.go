package vdf

import (
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
		{
			s: `"Root"
{
 "map"
 {
   "attr1" "hello"
 }
 "map"
 {
   "attr2" "world"
 }
}`,
			m: map[string]interface{}{
				"Root": map[string]interface{}{
					"map": map[string]interface{}{
						"attr1": "hello",
						"attr2": "world",
					},
				},
			},
		},
	}

	for i, tt := range tests {
		m, err := NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, (err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == nil && !reflect.DeepEqual(tt.m, m) {
			t.Errorf("%d. %q\n\nparse mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.m, m)
		}
	}
}

func TestParser_Parse(t *testing.T) {
	getReader := func(fileName string) (io.Reader, error) {
		return os.Open(fileName)
	}

	tests := []struct {
		name     string
		fileName string
		want     func(got map[string]interface{}, err error)
	}{
		{
			name:     "empty file",
			fileName: "testdata/empty.vdf",
			want: func(got map[string]interface{}, err error) {
				require.Error(t, err)
			},
		},
		{
			name:     "corrupted file",
			fileName: "testdata/corrupted.vdf",
			want: func(got map[string]interface{}, err error) {
				require.True(t, errors.Is(err, ErrNotValidFormat))
			},
		},
		{
			name:     "ok",
			fileName: "testdata/ok.vdf",
			want: func(got map[string]interface{}, err error) {
				require.NoError(t, err)
				saveFile := got["SaveFile"].(map[string]interface{})
				require.Equal(t, "ciccio", saveFile["team1"])
				require.Equal(t, "pasticcio", saveFile["team2"])
			},
		},
		// Full example sample file from
		// https://developer.valvesoftware.com/wiki/Counter-Strike:_Global_Offensive_Game_State_Integration#Sample_Configuration_File
		{
			name:     "gamestate_integration_consolesample",
			fileName: "testdata/gamestate_integration_consolesample.cfg",
			want: func(got map[string]interface{}, err error) {
				require.NoError(t, err)
				expected := map[string]interface{}{
					"Console Sample v.1": map[string]interface{}{
						"uri":       "http://127.0.0.1:3000",
						"timeout":   "5.0",
						"buffer":    "0.1",
						"throttle":  "0.5",
						"heartbeat": "60.0",
						"auth": map[string]interface{}{
							"token": "CCWJu64ZV3JHDT8hZc",
						},
						"output": map[string]interface{}{
							"precision_time":     "3",
							"precision_position": "1",
							"precision_vector":   "3",
						},
						"data": map[string]interface{}{
							"provider":           "1",
							"map":                "1",
							"round":              "1",
							"player_id":          "1",
							"player_state":       "1",
							"player_weapons":     "1",
							"player_match_stats": "1",
						},
					},
				}
				require.Equal(t, expected, got)
			},
		},
		// Comments that start with "// my comment" are supported
		// Broken comments like "/ my comment" are not.
		// The file will be partially parsed.
		{
			name:     "broken and unsupported comment",
			fileName: "testdata/broken_comment.cfg",
			want: func(got map[string]interface{}, err error) {
				require.NoError(t, err)
				expected := map[string]interface{}{
					"Broken Comment Sample v.1": map[string]interface{}{
						"uri":     "http://127.0.0.1:3456",
						"timeout": "8.0",
					},
				}
				require.Equal(t, expected, got)
			},
		},
		// Each level, even the first one, requires
		// surrounding curly braces. Intending is not enough.
		{
			name:     "no curly brace",
			fileName: "testdata/no_curly_brace.vdf",
			want: func(got map[string]interface{}, err error) {
				require.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, err := getReader(tt.fileName)
			require.NoError(t, err)
			p := NewParser(reader)
			got, err := p.Parse()
			tt.want(got, err)
		})
	}
}
