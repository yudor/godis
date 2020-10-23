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
}

var ConnectionError = errors.New("failed to establish connection to server")

func New(address string) (*godis, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &godis{connection: conn}, nil
}

func (g *godis) Ping() (string, error) {
	if _, err := g.writeBytes("PING"); err == nil {
		resp, err2 := g.readBytes(5)
		return resp, err2
	}
	return "", ConnectionError
}

func (g *godis) Get(key string) (string, error) {
	byteSize, _ := g.getMemoryBytes(key)
	if _, err := g.writeBytes(fmt.Sprintf("GET %v", key)); err == nil {
		resp, err2 := g.readBytes(byteSize)
		return resp, err2
	}
	return "", ConnectionError
}

func (g *godis) readBytes(length int) (string, error) {
	resp := make([]byte, length)
	_, err := g.connection.Read(resp)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func (g *godis) writeBytes(command string) (string, error) {
	_, err := g.connection.Write([]byte(command + "\r\n"))
	if err != nil {
		return "", err
	}
	return "", nil
}

func (g *godis) getMemoryBytes(key string) (int, error) {
	_, err := g.writeBytes(fmt.Sprintf("MEMORY USAGE %v", key))
	if err != nil {
		return 0, err
	}

	resp := make([]byte, 100)
	_, err = g.connection.Read(resp)
	if err != nil {
		return 0, err
	}
	formattedResp := strings.ReplaceAll(string(bytes.Trim(resp, "\x00")), "\r\n", "")

	if formattedResp == "$-1" {
		return 0, nil
	}

	return strconv.Atoi(formattedResp[1:])
}
