package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, obj interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error 1")

	}
	error := json.Unmarshal([]byte(body), obj)
	if error != nil {
		fmt.Println("Error 2")
	}
}
