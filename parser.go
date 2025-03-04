package confjson

import (
	"context"
	"encoding/json"
	"io"
)

// Parser is a conf.ParseFunc to parse the given json
func Parser(_ context.Context, r io.Reader) (any, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var res any
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}
