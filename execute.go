package vermouth

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	cert     string
	key      string
	listen   string
	monitor  string
	shutdown time.Duration
)

func init() {
	flag.StringVar(&cert, "cert", "", "TLS Cetificate file")
	flag.StringVar(&key, "key", "", "TLS private key")
	flag.StringVar(&listen, "listen", "[::]:5000", "listen addr")
	flag.StringVar(&monitor, "monitor", "", "monitor addr")
	flag.DurationVar(&shutdown, "shutdown", time.Second, "graceful shutdown timeout")
}

func Parse() {
	flag.VisitAll(func(f *flag.Flag) {
		key := strings.ToUpper(f.Name)
		if v, ok := os.LookupEnv(key); ok {
			f.Value.Set(v)
		}
	})
	flag.Parse()

	if cert != "" {
		if _, err := os.Stat(cert); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				log.Fatalf("cert not found: %s", cert)
			}
		}
		if key == "" {
			log.Fatal("key perameter required")
		}
		if _, err := os.Stat(key); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				log.Fatalf("key not found: %s", key)
			}
		}
	}
}

func Execute(h http.Handler) error {
	if err := startMonitor(); err != nil {
		return err
	}

	srv, err := startServer(h)
	if err != nil {
		return err
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	log.Printf("Signal Recieved: %s", <-sigCh)

	ctx, stop := context.WithTimeout(context.Background(), shutdown)
	defer stop()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	log.Print("Finished")
	return nil
}
