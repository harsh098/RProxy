package configs

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *configuration
		wantErr bool
	}{
		{
			name: "Check YAML Destructuring",
			want: &configuration{
				Gateway: struct {
					Host        string `yaml:"host"`
					Listen_port string `yaml:"listen_port"`
					Scheme      string `yaml:"scheme"`
				}(struct {
					Host        string
					Listen_port string
					Scheme      string
				}{
					Host:        "localhost",
					Listen_port: "8080",
					Scheme:      "http",
				}),
				Resources: []resource{
					{
						Name:         "Serv1",
						Endpoint:     "/server1",
						Upstream_URL: "http://localhost:9001",
					},
					{
						Name:         "Serv2",
						Endpoint:     "/server2",
						Upstream_URL: "http://localhost:9002",
					},
					{
						Name:         "Serv3",
						Endpoint:     "/server3",
						Upstream_URL: "http://localhost:9003",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
