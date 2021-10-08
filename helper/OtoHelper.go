package helper

import (
	"encoding/json"
	"net/http"
)

func ParseFromRaw(r *http.Request, model interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(model)
}
