package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LineOfLog struct {
    RemoteAddr  string
    ContentType string
    Path        string
    Query       string
    Method      string
    Body        string
}

type ReqestJson struct {
	Request string `json:"request"`
}

type ResponsJson struct {
	Respons string `json:"respons"`
}


func handler(w http.ResponseWriter, r *http.Request) {
    bufbody := new(bytes.Buffer)
    bufbody.ReadFrom(r.Body)
	body := bufbody.String()

	jsonBytes := ([]byte)(body)
	data := new(ReqestJson)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
        fmt.Println("JSON Unmarshal error:", err)
        return
    }

	respons := ResponsJson{Respons: data.Request}

	result, err := json.Marshal(respons)

	if err != nil {
        fmt.Println("JSON Marshal error:", err)
        return
    }

	w.Write([]byte(result))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}

