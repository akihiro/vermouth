package vermouth

import (
	"log"
	"net"
	"net/http"
)

func startServer(h http.Handler) (*http.Server, error) {
	log.Printf("Listen Addr: %s", listen)
	ln, err := net.Listen("tcp", listen)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{Handler: h}
	go func() {
		log.Print("Start Server")
		log.Print(srv.Serve(ln))
	}()

	return srv, nil
}
