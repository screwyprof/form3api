package form3api

import (
	"bytes"
	"encoding/json"
	"io"
)

const contentTypeJSON = "application/json"

type jsonSerializer struct{}

func (s *jsonSerializer) Serialize(req interface{}) (io.Reader, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	return buf, enc.Encode(req)
}

func (s *jsonSerializer) Deserialize(reader io.Reader, resp interface{}) error {
	return json.NewDecoder(reader).Decode(resp)
}

func (s *jsonSerializer) ContentType() string {
	return contentTypeJSON
}
