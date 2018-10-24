package indent

import (
	"bytes"
	"encoding/json"
)

// JSON - use for indenting json message from mq
func JSON(input []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, input, "", "\t")
	return out.Bytes(), err
}
