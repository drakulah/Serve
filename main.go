package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"serve/conv"
	"serve/ip"
	"serve/log"
	"serve/server"
)

func main() {

	var port int
	var serveDirPath string

	flag.IntVar(&port, "p", -1, "Port to start server")
	flag.StringVar(&serveDirPath, "dir", "./", "Directory to server")

	flag.Parse()

	cwd, cwdErr := os.Getwd()

	if cwdErr != nil {
		log.E("Couldn't access current working directory")
		return
	}

	if port < 0 {
		log.E("Invalid server port")
		return
	}

	serveDir := filepath.Join(cwd, serveDirPath)
	server := server.New()

	server.Get("^/", func(r *http.Request, w http.ResponseWriter) {
		http.FileServer(http.Dir(serveDir)).ServeHTTP(w, r)
	})

	server.Listen(":"+conv.IntToString(port), func() {
		log.I("Listening on: http://localhost:" + conv.IntToString(port))
		log.I("              http://" + ip.GetIPv4Addr() + ":" + conv.IntToString(port))
		log.I(fmt.Sprintf("Serving dir `%s`\n", serveDir))
	})
}
