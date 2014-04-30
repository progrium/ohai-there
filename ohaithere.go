package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func ohai(attribute string) ([]byte, error) {
	buff := new(bytes.Buffer)
	cmd := exec.Command("ohai", attribute)
	var stdout io.ReadCloser
	var err error
	if stdout, err = cmd.StdoutPipe(); err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	if _, err := io.Copy(buff, stdout); err != nil {
		return nil, err
	}
	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func main() {
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8000"
	}
	addr = ":" + addr

	log.Println("ohai there! serving ohai data on " + addr)
	log.Fatal(http.ListenAndServe(addr,
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			switch req.Method {
			case "GET":
				log.Println(req.Method, req.RequestURI)
				output, err := ohai(req.RequestURI[1:])
				if err != nil {
					log.Println(err)
				}
				w.Header().Set("Content-Length", strconv.Itoa(len(output)))
				w.Header().Set("Content-Type", "application/json")
				w.Write(output)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		}),
	))
}
