package jsontest

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wI2L/jettison"
	"testing"
)

func TestParsingFixtureJettison(t *testing.T) {
	b, err := loadFixture("default-unicode.json")
	require.NoError(t, err)

	spanStd := &Span{}
	err = json.Unmarshal(b, spanStd)
	require.NoError(t, err)

	jettisonBytes, err := jettison.Marshal(spanStd)
	require.NoError(t, err)
	spanStd2 := &Span{}
	err = json.Unmarshal(jettisonBytes, spanStd2)
	require.NoError(t, err)
	assert.Equal(t, spanStd, spanStd2)
}

func BenchmarkMarshalJettison(b *testing.B) {
	benchmarkMarshal(b, jettison.Marshal)
}
