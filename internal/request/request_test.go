package request

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type chunkreader struct {
	data            string
	numbytesperread int
	pos             int
}

// read reads up to len(p) or numbytesperread bytes from the string per call
// its useful for simulating reading a variable number of bytes per chunk from a network connection
func (cr *chunkreader) read(p []byte) (n int, err error) {
	if cr.pos >= len(cr.data) {
		return 0, io.eof
	}
	endindex := cr.pos + cr.numbytesperread
	if endindex > len(cr.data) {
		endindex = len(cr.data)
	}
	n = copy(p, cr.data[cr.pos:endindex])
	cr.pos += n

	return n, nil
}
func testrequestlineparse(t *testing.t) {
	// test: good get request line
	reader := &chunkreader{
		data:            "get / http/1.1\r\nhost: localhost:42069\r\nuser-agent: curl/7.81.0\r\naccept: */*\r\n\r\n",
		numbytesperread: 3,
	}
	r, err := requestfromreader(reader)
	require.noerror(t, err)
	require.notnil(t, r)
	assert.equal(t, "get", r.requestline.method)
	assert.equal(t, "/", r.requestline.requesttarget)
	assert.equal(t, "1.1", r.requestline.httpversion)

	// test: good get request line with path
	reader = &chunkreader{
		data:            "get /coffee http/1.1\r\nhost: localhost:42069\r\nuser-agent: curl/7.81.0\r\naccept: */*\r\n\r\n",
		numbytesperread: 1,
	}
	r, err = requestfromreader(reader)
	require.noerror(t, err)
	require.notnil(t, r)
	assert.equal(t, "get", r.requestline.method)
	assert.equal(t, "/coffee", r.requestline.requesttarget)
	assert.equal(t, "1.1", r.requestline.httpversion)
}

func testheaderparse(t *testing.t) {
	// test: standard headers
	reader := &chunkreader{
		data:            "get / http/1.1\r\nhost: localhost:42069\r\nuser-agent: curl/7.81.0\r\naccept: */*\r\n\r\n",
		numbytesperread: 3,
	}
	r, err := requestfromreader(reader)
	require.noerror(t, err)
	require.notnil(t, r)
	res, ok := r.headers.get("host")
	assert.true(t, ok)
	assert.equal(t, "localhost:42069", res)
	res, ok = r.headers.get("user-agent")
	assert.true(t, ok)
	assert.equal(t, "curl/7.81.0", res)
	res, ok = r.headers.get("accept")
	assert.true(t, ok)
	assert.equal(t, "*/*", res)
	// test: malformed header
	reader = &chunkreader{
		data:            "get / http/1.1\r\nhost localhost:42069\r\n\r\n",
		numbytesperread: 3,
	}
	r, err = requestfromreader(reader)
	require.error(t, err)
}

func testbodyparse(t *testing.t) {
	// Test: Standard Body
	reader := &chunkReader{
		data: "POST /submit HTTP/1.1\r\n" +
			"Host: localhost:42069\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"hello world!\n",
		numBytesPerRead: 3,
	}
	r, err := RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "hello world!\n", string(r.Body))

	// Test: Body shorter than reported content length
	reader = &chunkReader{
		data: "POST /submit HTTP/1.1\r\n" +
			"Host: localhost:42069\r\n" +
			"Content-Length: 20\r\n" +
			"\r\n" +
			"partial content",
		numBytesPerRead: 3,
	}
	r, err = RequestFromReader(reader)
	require.Error(t, err)
}
