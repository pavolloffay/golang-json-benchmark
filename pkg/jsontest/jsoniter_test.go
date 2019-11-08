package jsontest

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsingFixtureJsoniter(t *testing.T) {
	var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := loadFixture("default-unicode.json")
	require.NoError(t, err)

	span := &Span{}
	err = jsoniter.Unmarshal(b, span)
	require.NoError(t, err)
	spanStd := &Span{}
	err = json.Unmarshal(b, spanStd)
	require.NoError(t, err)
	// unmarshal is the same as stdlib
	assert.Equal(t, spanStd, span)

	bb, err := jsoniter.Marshal(span)
	require.NoError(t, err)

	spanStd = &Span{}
	err = json.Unmarshal(bb, spanStd)
	require.NoError(t, err)
	assert.Equal(t, spanStd, span)
}

func BenchmarkUnmarshalJsoninter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	benchmarkUnmarshal(b, json.Unmarshal)
}

func BenchmarkMarshalJsoniter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	benchmarkMarshal(b, json.Marshal)
}
