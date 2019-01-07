package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/azraeljack/algorithm-practice/p2p"
)

type Server struct {
	peers    map[string]*p2p.Peer
	listener *http.Server
	p2p      *net.TCPListener
	engine   p2p.Consensus
	table    map[string]string
	stop     chan struct{}

	tableMu sync.RWMutex
	peerMu  sync.RWMutex
}

func NewServer(listenAddr string, engine p2p.Consensus) *Server {
	server := &Server{
		listener: &http.Server{Addr: listenAddr, Handler: http.DefaultServeMux},
		p2p:      &net.TCPListener{},
		engine:   engine,
		stop:     make(chan struct{}, 1),
		table:    make(map[string]string),
		peers:    make(map[string]*p2p.Peer),
	}
	http.HandleFunc("/", server.requestHandler)
	return server
}

func (s *Server) Run() {
	go s.startHttp()
	go s.startP2P()
}

func (s *Server) startHttp() {
	if err := s.listener.ListenAndServe(); err != nil {
		log.Fatal("failed to start p2p server")
	}
}

func (s *Server) startP2P() {
	defer s.p2p.Close()

	for {
		conn, err := s.p2p.AcceptTCP()
		if err != nil {
			log.Print("failed to accept connection")
			continue
		}
		go s.p2pHandler(conn)
	}
}

func (s *Server) Shutdown() {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	_ = s.listener.Shutdown(ctx)
}

func (s *Server) requestHandler(w http.ResponseWriter, req *http.Request) {

}

func (s *Server) p2pHandler(conn *net.TCPConn) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	for {
		select {
		case _, _ = <-s.stop:
			return
		default:
			var msg p2p.Message
			err := decoder.Decode(msg)
			if err != nil {
				log.Println("failed to decode message, reason: ", err.Error())
				return
			}
		}
	}
}

func main() {
	log.SetOutput(os.Stdout)
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Kill, os.Interrupt)

	<-interrupt
	log.Println("server exited")
}
