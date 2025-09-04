package headers

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var rn = []byte("\r\n")

func parseHeader(fieldLines []byte) (string, string, error) {
	parts := bytes.SplitN(fieldLines, []byte(":"), 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Malformed Field Line")
	}

	regex := regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.\^_` + "`" + `|~]+$`)

	name := parts[0]
	value := bytes.TrimSpace(parts[1])

	if bytes.HasSuffix(name, []byte(" ")) || !regex.Match(name) {
		return "", "", fmt.Errorf("Malformed Field Name")
	}

	return string(name), string(value), nil
}

type Headers struct {
	headers map[string]string
}

func (h *Headers) Get(name string) string {
	return h.headers[strings.ToLower(name)]
}

func (h *Headers) Set(name, value string) {

	name = strings.ToLower(name)
	if v, ok := h.headers[strings.ToLower(name)]; ok {
		h.headers[name] = v + ", " + value
	} else {
		h.headers[name] = value
	}
}

func (h *Headers) Parse(data []byte) (int, bool, error) {
	read := 0
	done := false
	for {
		idx := bytes.Index(data[read:], rn)
		if idx == -1 {
			break
		}

		// Empty Headers
		if idx == 0 {
			read += len(rn)
			done = true
			break
		}

		name, value, err := parseHeader(data[read : read+idx])

		read += idx + len(rn)
		if err != nil {
			return 0, false, err
		}

		h.Set(name, value)
	}

	return read, done, nil
}

func NewHeaders() *Headers {
	return &Headers{
		headers: map[string]string{},
	}
}
