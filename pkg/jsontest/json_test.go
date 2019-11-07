package jsontest

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type UnmarshalFce func([]byte, interface{}) error
type MarshalFce func(interface{}) ([]byte, error)

func TestParsingFixtureStdlib(t *testing.T) {
	b, err := loadFixture("default.json")
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

func benchmarkUnmarshal(b *testing.B, unmarshal UnmarshalFce) {
	b.Run("default.json", func(b *testing.B) {
		json, err := loadFixture("default.json")
		require.NoError(b, err)
		trace := &Span{}
		for i := 0; i < b.N; i++ {
			err = unmarshal(json, trace)
		}
	})
	b.Run("default-unicode.json", func(b *testing.B) {
		json, err := loadFixture("default-unicode.json")
		require.NoError(b, err)
		trace := &Span{}
		for i := 0; i < b.N; i++ {
			err = unmarshal(json, trace)
		}
	})
}

func benchmarkMarshal(b *testing.B, marshall MarshalFce) {
	b.Run("default.json", func(b *testing.B) {
		bb, err := loadFixture("default.json")
		require.NoError(b, err)
		trace := &Span{}
		err = json.Unmarshal(bb, trace)
		require.NoError(b, err)

		for i := 0; i < b.N; i++ {
			marshall(trace)
		}
	})
	b.Run("default-unicode.json", func(b *testing.B) {
		bb, err := loadFixture("default-unicode.json")
		require.NoError(b, err)
		trace := &Span{}
		err = json.Unmarshal(bb, trace)
		require.NoError(b, err)

		for i := 0; i < b.N; i++ {
			marshall(trace)
		}
	})
}

func BenchmarkUnmarshalStdlib(b *testing.B) {
	benchmarkUnmarshal(b, json.Unmarshal)
}

func BenchmarkMarshalStdlib(b *testing.B) {
	benchmarkMarshal(b, json.Marshal)
}

func loadFixture(file string) ([]byte, error) {
	return ioutil.ReadFile("fixtures/" + file)
}

