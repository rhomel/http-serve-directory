package server

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Serve() {
	(&server{}).Serve()
}

type server struct {
	listenAddress string
	dirPath       string
	absDirPath    string
	dirFS         fs.FS
	logAccesses   bool
}

var usageHeader string = `
Serve a directory using Go's standard http file server.

It defaults to the current directory. Use -directory to specify a custom path.

Options:
`

func (s *server) parseFlags() {
	flag.StringVar(&s.dirPath, "directory", ".", "directory path to serve")
	flag.StringVar(&s.listenAddress, "address", "0.0.0.0:3000", "address:port to listen on")
	flag.BoolVar(&s.logAccesses, "log", false, "log http requests")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usageHeader)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func (s *server) getAbsPath() error {
	absDirPath, err := filepath.Abs(s.dirPath)
	if err != nil {
		return fmt.Errorf("directory %s is not valid: %w", s.dirPath, err)
	}
	s.absDirPath = absDirPath
	return nil
}

func (s *server) getHandler() http.Handler {
	handler := http.FileServer(http.FS(s.dirFS))
	if s.logAccesses {
		handler = NewLoggingHandler(handler)
	}
	return handler
}

func (s *server) configure() error {
	s.parseFlags()
	if err := s.getAbsPath(); err != nil {
		return err
	}
	s.dirFS = os.DirFS(s.dirPath)
	return nil
}

func (s *server) Serve() {
	if err := s.configure(); err != nil {
		log.Fatal("failed to initialize:", err)
	}
	s.logConfig()
	handler := s.getHandler()
	http.ListenAndServe(s.listenAddress, handler)
}

func (s *server) logConfig() {
	s.logDirPath()
	log.Printf("listening address: http://%s", s.listenAddress)
}

func (s *server) logDirPath() {
	str := fmt.Sprintf("serving directory: %s", s.dirPath)
	if s.dirPath != s.absDirPath {
		str += fmt.Sprintf(" (%s)", s.absDirPath)
	}
	log.Printf(str)
}
