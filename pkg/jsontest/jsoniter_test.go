package jsontest

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParsingFixtureJsoniter(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := loadFixture("default-unicode.json")
	require.NoError(t, err)
	assert.NotNil(t, b)

	span := &Span{}
	err = json.Unmarshal(b, span)
	require.NoError(t, err)

	bb, err := json.Marshal(span)
	require.NoError(t, err)

	span2 := &Span{}
	err = json.Unmarshal(bb, span2)
	require.NoError(t, err)
	assert.Equal(t, span, span2)
}

func BenchmarkUnmarshalJsoninter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	benchmarkUnmarshal(b, json.Unmarshal)
}

func BenchmarkMarshalJsoniter(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	benchmarkMarshal(b, json.Marshal)
}
