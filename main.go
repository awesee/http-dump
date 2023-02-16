package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

var port uint

func init() {
	flag.UintVar(&port, "port", 8090, "TCP address for the server to listen on")
	flag.Parse()
}

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
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("http://127.0.0.1%s\n", addr)
	log.Println(http.ListenAndServe(addr, nil))
}
