package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type httpSrv struct {
	*http.Server
	ctx context.Context
}

func (h *httpSrv) Start(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	err = h.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (h *httpSrv) Stop() error {
	return h.Shutdown(h.ctx)
}

func newServer(mux *http.ServeMux, ctx context.Context) *httpSrv {
	debugHttp := httpSrv{Server: &http.Server{Handler: mux}, ctx: ctx}
	return &debugHttp
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, errCtx := errgroup.WithContext(ctx)
	//第一个http
	debugMux := http.NewServeMux()
	debugMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "debug goruntine")
	})
	debugHttp := newServer(debugMux, errCtx)
	g.Go(func() error {
		return debugHttp.Start("127.0.0.1:8080")
	})
	g.Go(func() error {
		<-errCtx.Done()
		fmt.Println("debug 退出")
		return debugHttp.Stop()
	})
	//第二个http
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		select {
		case <-time.After(1 * time.Second):
			fmt.Fprintf(writer, "时间到期")
			cancel()
		}
	})
	muxHttp := newServer(mux, errCtx)
	g.Go(func() error {
		return muxHttp.Start("127.0.0.1:8081")
	})
	g.Go(func() error {
		<-errCtx.Done()
		fmt.Println("mux 退出")
		return muxHttp.Stop()
	})

	//命令
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT}
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	g.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-c:
				cancel()
			}
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("退出 %+v", err)
	}
}
