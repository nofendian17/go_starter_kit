package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Config
	}{
		{
			name: "success",
			want: &Config{
				Application: Application{
					Name:    "MyApp",
					Version: "1.0.0",
					Address: "localhost",
					Port:    3000,
				},
				Database: Database{
					Driver:        "mysql",
					Host:          "localhost",
					Port:          3306,
					Username:      "root",
					Password:      "Password123",
					Database:      "test",
					Charset:       "utf8mb4",
					Timezone:      "Asia/Jakarta",
					MaxIdleConns:  10,
					MaxOpenConns:  100,
					SSLMode:       false,
					Debug:         true,
					DisableMetric: true,
					DisableTrace:  true,
				},
				Cache: Cache{
					Host:     "localhost",
					Port:     6379,
					Username: "",
					Password: "",
					Database: 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			assert.Equal(t, tt.want, got)
		})
	}
}
