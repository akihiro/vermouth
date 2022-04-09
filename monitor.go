package vermouth

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func startMonitor() error {
	if len(monitor) == 0 {
		return nil
	}

	log.Printf("Monitor Listen: %s", monitor)
	ln, err := net.Listen("tcp", monitor)
	if err != nil {
		return err
	}

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Print("Start Monitor")
		http.Serve(ln, nil)
	}()
	return nil
}
