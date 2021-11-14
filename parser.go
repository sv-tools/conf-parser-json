package confjson

import (
	"context"
	"encoding/json"
	"io"
)

// Parser is a conf.ParseFunc to parse the given json
func Parser(_ context.Context, r io.Reader) (interface{}, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var res interface{}
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}
