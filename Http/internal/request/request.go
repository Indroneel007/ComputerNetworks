package request

import (
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

var ERROR_BAD_START_LINE = fmt.Errorf("bad start line")
var SEPARATOR = "\r\n"

func ParseRequestLine(b string) (*RequestLine, string, error) {
	idx := strings.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, "", nil
	}

	startLine := b[:idx]
	restOfMesg := b[idx+len(SEPARATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, "", ERROR_BAD_START_LINE
	}

	return &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}, restOfMesg, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	str := string(data)
	reqLine, _, err := ParseRequestLine(str)
	if err != nil {
		return nil, err
	}
	//Lets see
	return &Request{
		RequestLine: *reqLine,
	}, nil
}
