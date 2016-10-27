package main

import (
    "log"
    "fmt"
    "flag"
    "runtime"
    "net/url"

    "hello/foo"

    "github.com/valyala/fasthttp"
)

var (
    port int
)

func usage(ctx *fasthttp.RequestCtx) {
    fmt.Fprintf(ctx, "Usage:\n")
    fmt.Fprintf(ctx, "  %-20s %s\n", "/hello", "say hello")
    fmt.Fprintf(ctx, "  %-20s %s\n", "/info", "show request info")
    fmt.Fprintf(ctx, "  %-20s %s\n", "/redirect?t=<url>", "redirect to url")
}

func helloHandler(ctx *fasthttp.RequestCtx) {
    fmt.Fprintf(ctx, "hello\n")
}

func infoHandler(ctx *fasthttp.RequestCtx) {
    fmt.Fprintf(ctx, "%s", ctx.Request.Header.Header())
}

func redirectHandler(ctx *fasthttp.RequestCtx) {
    args, err := url.ParseQuery(ctx.QueryArgs().String())
    if err != nil {
	fmt.Fprintf(ctx, "parameters error: %s", err)
	return
    }

    t := args.Get("t")
    if t == "" {
	fmt.Fprintf(ctx, "miss parameter 't'")
	return
    }

    ctx.Redirect(t, 302)
}

func handler(ctx *fasthttp.RequestCtx) {
    switch string(ctx.Path()) {
    case "/info":
	infoHandler(ctx)
    case "/hello":
	helloHandler(ctx)
    case "/redirect":
	redirectHandler(ctx)
    default:
	usage(ctx)
    }
}

func init() {
    flag.IntVar(&port, "p", 9000, "server listen on")

    runtime.GOMAXPROCS(4)
}

func main() {
    flag.Parse()

    foo.Foo()

    log.Printf("listen on %d", port)

    log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
