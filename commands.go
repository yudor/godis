package godis

import "fmt"

func (g *godis) Keys() ([]string, error) {
	if _, err := g.writeBytes(fmt.Sprintf("KEYS %v", "*")); err == nil {
		resp, _ := g.readBytes(64)
		_, parseErr := g.p.Parse(resp)
		if parseErr == nil {
			return []string{}, nil
		}
	}
	return []string{}, ConnectionError
}

func (g *godis) Flush() (interface{}, error) {
	if _, err := g.writeBytes("FLUSHALL"); err == nil {
		resp, _ := g.readBytes(5)
		opResult, parseErr := g.p.Parse(resp)
		if parseErr == nil {
			return opResult == "OK", nil
		}
	}
	return false, ConnectionError
}

func (g *godis) Ping() (interface{}, error) {
	if _, err := g.writeBytes("PING"); err == nil {
		resp, _ := g.readBytes(5)
		return g.p.Parse(resp)
	}
	return "", ConnectionError
}

func (g *godis) Get(key string) (interface{}, error) {
	byteSize, _ := g.getMemoryBytes(key)
	if _, err := g.writeBytes(fmt.Sprintf("GET %v", key)); err == nil {
		resp, _ := g.readBytes(byteSize)
		return g.p.Parse(resp)
	}
	return "", ConnectionError
}

func (g *godis) Set(key string, value interface{}) (bool, error) {
	_, err := g.writeBytes(fmt.Sprintf("SET %v \"%v\"", key, value))
	if err == nil {
		resp, _ := g.readBytes(5)
		opResult, parseErr := g.p.Parse(resp)
		if parseErr == nil {
			return opResult == "OK", nil
		}
	}
	return false, ConnectionError
}
