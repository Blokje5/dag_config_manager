package task

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	simpleCase := `
name: test
provider: test
`

	dateCase := `
name: test
provider: test
file: test
`
	tests := []struct {
		name    string
		data []byte
		task    Task
	}{
		{
			"Simple case without data",
			[]byte(simpleCase),
			Task{
				Name: "test",
				Provider: "test",
			},
		},
		{
			"Basic extra keys",
			[]byte(dateCase),
			Task{
				Name: "test",
				Provider: "test",
				Data: map[string]interface{}{
					"file": "test",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Parse(tc.data)
			if err != nil {
				t.Fatalf("Unexpected Parse error: %v", err)
			}
			if !reflect.DeepEqual(got, tc.task) {
				t.Errorf("Parse() = %v, want %v", got, tc.task)
			}
		})
	}
}
