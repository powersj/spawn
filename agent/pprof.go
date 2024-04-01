package agent

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" //nolint:blankimports // needed for pprof
	"time"
)

type Server interface {
	Start(string)
	ErrChan() <-chan error
}

type PprofServer struct {
	err chan error
}

func (p *PprofServer) Start(port string) {
	go func() {
		address := "localhost:" + port
		pprofHostPort := "http://" + address + "/debug/pprof"

		fmt.Printf("I! Starting pprof HTTP server at: %s", pprofHostPort)

		server := &http.Server{
			Addr:         address,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		if err := server.ListenAndServe(); err != nil {
			p.err <- err
		}
		close(p.err)
	}()
}

func (p *PprofServer) ErrChan() <-chan error {
	return p.err
}
