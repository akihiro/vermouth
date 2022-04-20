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
		if cert != "" {
			log.Print("Start TLS Server")
			log.Print(srv.ServeTLS(ln, cert, key))
		} else {
			log.Print("Start Plaintext Server")
			log.Print(srv.Serve(ln))
		}
	}()

	return srv, nil
}
