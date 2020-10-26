package godis

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

const (
	ErrorReply        = '-'
	SimpleStringReply = '+'
	IntReply          = ':'
	BulkStringReply   = '$'
	ArrayReply        = '*'
)

type parser struct{}
type ArrString []string

func (p *parser) Parse(payload []byte) (interface{}, error) {
	fmt.Printf("{buffer: %v},\n{payload: %v}\n\n", payload, string(payload))
	lines := p.parseLines(payload)

	// a empty reply was sent
	if isNilReply(lines[0]) || len(lines[0]) == 0 {
		return nil, nil
	}

	// a bulk string response was sent
	if lines[0][0] == BulkStringReply {
		size, _ := strconv.Atoi(string(lines[0][1]))
		if size > 0 {
			return p.parseBulkString(lines[1])
		}
	}

	if lines[0][0] == ArrayReply {
		return p.parseArray(payload)
	}

	if lines[0][0] == SimpleStringReply {
		return p.parseString(lines[0][1:])
	}

	if lines[0][0] == IntReply {
		return p.parseInt(lines[0][1:])
	}

	if lines[0][0] == ErrorReply {
		return nil, nil
	}

	return nil, nil
}

func (p *parser) parseLines(payload []byte) [][]byte {
	return bytes.Split(payload, []byte("\r\n"))
}

func (p *parser) parseString(payload []byte) (string, error) {
	return string(payload), nil
}

func (p *parser) parseInt(payload []byte) (int, error) {
	return strconv.Atoi(string(payload))
}

func (p *parser) parseBulkString(payload []byte) (interface{}, error) {
	matched, _ := regexp.MatchString("^[0-9]+$", string(payload))
	if matched {
		return strconv.Atoi(string(payload))
	}
	return string(payload), nil
}

func (p *parser) parseArray(payload []byte) ([]string, error) {
	fmt.Printf("{buffer: %v},\n{payload: %v}\n\n", payload, string(payload))
	lines := p.parseLines(payload)
	arraySize, _ := p.parseInt(lines[0][1:])
	var array []string

	if arraySize == 0 {
		return []string{}, nil
	}

	for i := 1; i < len(lines)-1; i += 2 {
		array = append(array, string(lines[i+1]))
	}

	return array, nil
}

func isNilReply(payload []byte) bool {
	return len(payload) == 3 && payload[0] == BulkStringReply && payload[1] == ErrorReply && payload[2] == '1'
}
