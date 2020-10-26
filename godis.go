package godis

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type godis struct {
	connection net.Conn
	p          parser
}

var ConnectionError = errors.New("failed to establish connection to server")

func New(address string) (*godis, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &godis{connection: conn}, nil
}

func (g *godis) readBytes(length int) ([]byte, error) {
	resp := make([]byte, length)
	_, err := g.connection.Read(resp)
	if err != nil {
		return []byte{}, err
	}
	return resp, nil
}

func (g *godis) writeBytes(command string) (int, error) {
	fmt.Printf("command: %v\n", command)
	length, err := g.connection.Write([]byte(command + "\r\n"))
	if err != nil {
		return 0, err
	}
	return length, nil
}

func (g *godis) getMemoryBytes(key string) (int, error) {
	_, err := g.writeBytes(fmt.Sprintf("MEMORY USAGE %v", key))
	if err != nil {
		return 0, err
	}

	resp := make([]byte, 8)
	_, err = g.connection.Read(resp)
	if err != nil {
		return 0, err
	}
	fmt.Printf("{buffer: %v},\n{payload: %v}\n\n", resp, string(resp))
	formattedResp := strings.ReplaceAll(string(bytes.Trim(resp, "\x00")), "\r\n", "")

	if formattedResp == "$-1" {
		return 0, nil
	}

	size, convertErr := strconv.Atoi(formattedResp[1:])

	return size / 2, convertErr
}
