package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"bibit/handler"

	"github.com/rs/cors"
)

func init() {
	prometheus.MustRegister(handler.PingCounter)
	prometheus.MustRegister(handler.TrxCounter)
	prometheus.MustRegister(handler.TrxStatusCodeCounter)
	prometheus.MustRegister(handler.TrxTypeCounter)
	prometheus.MustRegister(handler.TrxHistogram)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/api/trx", handler.Trx)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", handler.HelloServer)

	handler := cors.Default().Handler(mux)

	var host string
	e := os.Getenv("TRX_HOST")
	if e == "" {
		host = ":8080"
	}

	http.ListenAndServe(host, handler)

}
