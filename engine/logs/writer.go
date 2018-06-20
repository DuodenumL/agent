package logs

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/url"

	log "github.com/sirupsen/logrus"
	"github.com/projecteru2/agent/types"
)

type Writer struct {
	addr    string
	scheme  string
	conn    io.Writer
	stdout  bool
	encoder *json.Encoder
	Close   func() error
}

func NewWriter(addr string, stdout bool) (*Writer, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	writer := &Writer{addr: u.Host, scheme: u.Scheme}
	writer.stdout = stdout
	err = writer.CreateConn()
	if err != nil {
		return nil, err
	}
	return writer, nil
}

func (w *Writer) CreateConn() error {
	switch {
	case w.scheme == "udp":
		err := w.createUDPConn()
		return err
	case w.scheme == "tcp":
		err := w.createTCPConn()
		return err
	}
	return fmt.Errorf("Invalid scheme: %s", w.scheme)
}

func (w *Writer) Write(logline *types.Log) error {
	if w.stdout {
		log.Info(logline)
	}
	err := w.encoder.Encode(logline)
	if err != nil {
		w.Close()
		w.CreateConn()
	}
	return err
}

func (w *Writer) createUDPConn() error {
	udpAddr, err := net.ResolveUDPAddr("udp", w.addr)
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	w.conn = conn
	w.encoder = json.NewEncoder(conn)
	w.Close = conn.Close
	return nil
}

func (w *Writer) createTCPConn() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", w.addr)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	w.conn = conn
	w.encoder = json.NewEncoder(conn)
	w.Close = conn.Close
	return nil
}
