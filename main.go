package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s\n", data)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		_ = r.ParseForm()
		data, err = json.MarshalIndent(map[string]interface{}{
			"Method":     r.Method,
			"URL":        r.URL.String(),
			"Proto":      r.Proto,
			"Header":     r.Header,
			"Body":       string(body),
			"Host":       r.Host,
			"Form":       r.Form,
			"Query":      r.URL.Query(),
			"PostForm":   r.PostForm,
			"RemoteAddr": r.RemoteAddr,
			"RequestURI": r.RequestURI,
		}, "", "\t")
		if err != nil {
			log.Println(err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write(data)
	})
	log.Println(http.ListenAndServe(":8090", nil))
}
