package datasource

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

// SampleDatasource is an example datasource used to scaffold
// new datasource plugins with an backend.
type SampleDatasource struct{}

// New returns a new SampleDatasource
func New() *SampleDatasource {
	return &SampleDatasource{}
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifer).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (td *SampleDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	log.DefaultLogger.Info("QueryData", "request", req)

	// create response struct
	response := backend.NewQueryDataResponse()

	// loop over queries and execute them individually.
	for _, q := range req.Queries {
		res, err := td.query(ctx, q)
		if err != nil {
			return nil, err
		}

		// save the response in a hashmap
		// based on with RefID as identifier
		response.Responses[q.RefID] = res
	}

	return response, nil
}

type queryModel struct {
	Format string `json:"format"`
}

func (td *SampleDatasource) query(ctx context.Context, query backend.DataQuery) (backend.DataResponse, error) {
	// Unmarshal the json into our queryModel
	var qm queryModel
	response := backend.DataResponse{}
	err := json.Unmarshal(query.JSON, &qm)
	if err != nil {
		return response, err
	}

	// Return an error is `Format` is empty. Returning an error on the `DataResponse`
	// will allow others queries to be executed. If we return an error as the second
	// param we expect to halt all queries.
	if qm.Format == "" {
		response.Error = errors.New("format cannot be empty")
		return response, nil
	}

	// create data frame response
	frame := data.NewFrame("response")
	frame.Fields = append(frame.Fields, data.NewField("countries", nil, []string{"Sweden", "Belgium", "Germany"}))

	// add the frames to the response
	response.Frames = append(response.Frames, frame)

	return response, nil
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (td *SampleDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	var status = backend.HealthStatusOk
	var message = ""

	if rand.Int()%2 == 0 {
		status = backend.HealthStatusError
		message = "randomized error"
	}

	return &backend.CheckHealthResult{
		Status:  status,
		Message: message,
	}, nil
}