package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// BindParam bind to request parameter to struct
func BindParam(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return decoder.Decode(dst, r.Form)
}

// BindJSONParam bind to request body to struct
func BindJSONParam(r *http.Request, dst interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return err
	}

	return json.Unmarshal(b, dst)
}
