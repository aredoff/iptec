package blacklist

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func serialization(lines []string) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&lines)
	if err != nil {
		return nil, fmt.Errorf("cant serialize list, err=%s", err)
	}
	return buffer.Bytes(), nil
}

func deserialization(data []byte) ([]string, error) {
	var lines []string
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&lines)
	if err != nil {
		return nil, fmt.Errorf("cant deserialize list, err=%s", err)
	}
	return lines, nil
}
