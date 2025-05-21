package confjson_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sv-tools/conf"

	confjson "github.com/sv-tools/conf-parser-json"
)

func TestParser(t *testing.T) {
	data := strings.NewReader(`{"foo": 42, "bar": "test"}`)
	c := conf.New().WithReaders(conf.NewStreamParser(data).WithParser(confjson.Parser))
	require.NoError(t, c.Load(t.Context()))

	require.Equal(t, 42, c.GetInt("foo"))
	require.Equal(t, "test", c.Get("bar"))
}

var errFake = errors.New("fake error")

type testReader struct{}

func (t *testReader) Read(_ []byte) (int, error) {
	return 0, errFake
}

func TestParserErrors(t *testing.T) {
	c := conf.New().WithReaders(conf.NewStreamParser(&testReader{}).WithParser(confjson.Parser))
	require.ErrorIs(t, c.Load(t.Context()), errFake)

	data := strings.NewReader(`{"foo": 42, "bar": "test"`)
	c = conf.New().WithReaders(conf.NewStreamParser(data).WithParser(confjson.Parser))
	require.ErrorContains(t, c.Load(t.Context()), "unexpected end of JSON input")
}

func ExampleParser() {
	data := strings.NewReader(`{"foo": 42, "bar": "test"}`)
	c := conf.New().WithReaders(conf.NewStreamParser(data).WithParser(confjson.Parser))
	if err := c.Load(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println(c.GetInt("foo"))
	// Output: 42
}
