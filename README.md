# vermouth
Web server skelton for kubernetes

## startup error handling

## graceful shutdown

## monitoring

- Split service port and monitoring port 
  - `LISTEN` service port
  - `MONITOR` monitoring port
- Use [prometheus](https://github.com/prometheus/client_golang)
  - monitor port `/metrics`
  - GO runtime monitoring
- Use [pprof](https://pkg.go.dev/net/http/pprof)
  - monitor port `/debug/pprof/`
