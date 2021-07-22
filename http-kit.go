package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	return cfg.Build()
}

func main_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s | %s", "/dump", "Dump request and echo back.")
}

func dump_handler(w http.ResponseWriter, r *http.Request) {
	headers, err := httputil.DumpRequest(r, false)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reading body: " + err.Error())
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	logger.Info("dump request",
		zap.String("header", string(headers)),
		zap.String("body", string(body)),
		zap.String("remote_addr", r.RemoteAddr),
	)

	fmt.Fprintf(w, "%s%s", headers, body)
}

func main() {
	logger, _ = NewLogger()

	http.HandleFunc("/", main_handler)
	http.HandleFunc("/dump", dump_handler)
	listen := os.Getenv("HTTP_PORT")
	if listen == "" {
		listen = ":80"
	}
	logger.Info("Listening on " + listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}
