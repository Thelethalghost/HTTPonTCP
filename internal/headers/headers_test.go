package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaderParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	res, ok := headers.Get("Host")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069", res)
	assert.Equal(t, 25, n)
	assert.True(t, done)

	// Test: Valid double header
	headers = NewHeaders()
	data = []byte("Host: localhost:42069\r\nFooFoo: barbar\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	res, ok = headers.Get("Host")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069", res)
	res, ok = headers.Get("FooFoo")
	assert.True(t, ok)
	assert.Equal(t, "barbar", res)
	res, ok = headers.Get("MissingKey")
	assert.False(t, ok)
	assert.Equal(t, 41, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Valid with Capital header
	headers = NewHeaders()
	data = []byte("HOST: localhost:42069\r\nFoOfOO: barbar\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	res, ok = headers.Get("HOST")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069", res)
	res, ok = headers.Get("FoofOO")
	assert.True(t, ok)
	assert.Equal(t, "barbar", res)
	res, ok = headers.Get("MissingKey")
	assert.False(t, ok)
	assert.Equal(t, 41, n)
	assert.True(t, done)

	// Test: Inbalid Header Name with invalid Characters
	headers = NewHeaders()
	data = []byte("H©st: localhost:42069\r\nFooFoo: barbar\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	require.Equal(t, 0, n)
	require.False(t, done)

	// Test: Valid Header with 1 key multiple values
	headers = NewHeaders()
	data = []byte("HOST: localhost:42069\r\nHost: barbar\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	res, ok = headers.Get("HOST")
	assert.True(t, ok)
	assert.Equal(t, "localhost:42069, barbar", res)
	res, ok = headers.Get("MissingKey")
	assert.False(t, ok)
	assert.Equal(t, 39, n)
	assert.True(t, done)
}
