# httpserver
HTTP server routers

## Usage:

```go
r := mux.NewRouter()
// ...
http.Handle("/", httpserver.PanicRecoveryMiddleware(httpserver.LogMiddleware(r)))
s := &http.Server{
	Addr:           listenAddress,
	Handler:        nil,
	ReadTimeout:    1000 * time.Second,
	WriteTimeout:   1000 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
```
