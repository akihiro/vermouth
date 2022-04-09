package vermouth

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	listen   string
	monitor  string
	shutdown time.Duration
)

func init() {
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
