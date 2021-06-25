package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

func main() {
	addr := os.Getenv("PLUGIN_ADDR")
	c, cancel, err := NewPluginClient(addr)
	if err != nil {
		panic(err)
	}
	defer cancel()

	fmt.Println("Making request")
	_, err = c.QueryData(context.Background(), &backend.QueryDataRequest{
		PluginContext: backend.PluginContext{
			DataSourceInstanceSettings: &backend.DataSourceInstanceSettings{
				ID:               0,
				UID:              "",
				Name:             "",
				URL:              "",
				User:             "",
				Database:         "",
				BasicAuthEnabled: false,
				BasicAuthUser:    "",
				JSONData:         []byte{},
				DecryptedSecureJSONData: map[string]string{
					"": "",
				},
				Updated: time.Time{},
			},
		},
		Headers: map[string]string{
			"": "",
		},
		Queries: []backend.DataQuery{
			{
				JSON: []byte(fmt.Sprintf(`"%s"`, os.Args[1])),
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
