package connection

import (
	"reflect"
	"testing"
)

func TestAuthMetadataFromYaml(t *testing.T) {
	var simplePasswordAuth = `
simplePasswordAuth:
  user: user
  password: password
`
	tests := []struct {
		name string
		data []byte
		want *AuthMetadata
	}{
		{
			"Expect correct serialization of simple auth",
			[]byte(simplePasswordAuth),
			&AuthMetadata{
				&SSHSimplePasswordAuth{
					UserName: "user",
					Password: "password",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := AuthMetadataFromYaml(tc.data)
			if err != nil {
				t.Fatalf("Unexpected serialization error %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("AuthMetadataFromYaml() = %v, want %v", got, tc.want)
			}
		})
	}
}
