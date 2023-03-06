package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"serve/conv"
	"serve/fs"
	"serve/ip"
	"serve/log"
	"serve/server"
	"strings"
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
		reqPath := filepath.Join(serveDir, r.URL.Path)
		fileInfo, statErr := os.Stat(reqPath)

		if statErr != nil {
			log.E(reqPath)
			w.WriteHeader(404)
			return
		}

		log.I(reqPath)

		if fileInfo.IsDir() {
			rawPartialHTML := ""
			for _, v := range fs.ReadDir(reqPath) {
				icon := "file"
				if v.IsDir {
					icon = "folder"
				}
				rawPartialHTML += fmt.Sprintf("<a class='file-card icon %s' href='%s'>%s</a>", icon, url.PathEscape(v.Name), v.Name)
			}
			w.Write([]byte(CreateDocument(reqPath, rawPartialHTML)))
		} else {
			http.FileServer(http.Dir(serveDir)).ServeHTTP(w, r)
		}
	})

	server.Listen(":"+conv.IntToString(port), func() {
		log.I("Listening on: http://localhost:" + conv.IntToString(port))
		log.I("              http://" + ip.GetIPv4Addr() + ":" + conv.IntToString(port))
		log.I(fmt.Sprintf("Serving dir `%s`", serveDir))
	})
}

func CreateDocument(heading string, rawHTML string) string {
	doc := "<!DOCTYPE html><html lang='en'><head> <meta charset='UTF-8'><meta http-equiv='X-UA-Compatible' content='IE=edge'> <meta name='viewport' content='width=device-width, initial-scale=1.0'><title>{{title}}</title><style>*{margin: 0; border: 0; padding: 0; outline: 0; line-height: 1.1; box-sizing: border-box;}body{width: 100%; height: 100vh; color: #eee; overflow: hidden; background-color: #111; font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;}.root{gap: 2em; width: 100%; height: 100%; overflow-y: scroll; display: flex; padding: 4em 12em; flex-direction: column;}.heading{color: #fff; font-size: 36px; font-weight: 700;}.body{gap: 0.5em; width: 100%; display: grid; grid-template-columns: repeat(3, 1fr);}.file-card{gap: 0.6em; color: #ccc; display: flex; padding: 1.3em; cursor: pointer; transition: 100ms; border-radius: 4px; align-items: center; text-decoration: none; background-color: #1e1e1e;}.file-card:hover{background-color: #262626;}.file-card > svg{width: 22px; height: 22px; fill: #ccc;}</style></head><body> <div class='root'> <div class='top-container'> <span class='heading'>{{title}}</span></div><div class='body'>{{rawHTML}}</div></div><script>document.querySelectorAll('.icon').forEach(e=>{if (e.classList.contains('file')) e.innerHTML=`<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 384 512'><path d='M0 64C0 28.7 28.7 0 64 0H224V128c0 17.7 14.3 32 32 32H384V448c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V64zm384 64H256V0L384 128z'/></svg>` + e.innerHTML; if (e.classList.contains('folder')) e.innerHTML=`<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 512 512'><path d='M64 480H448c35.3 0 64-28.7 64-64V160c0-35.3-28.7-64-64-64H288c-10.1 0-19.6-4.7-25.6-12.8L243.2 57.6C231.1 41.5 212.1 32 192 32H64C28.7 32 0 60.7 0 96V416c0 35.3 28.7 64 64 64z'/></svg>` + e.innerHTML;}) </script></body></html>"
	doc = strings.ReplaceAll(doc, "{{title}}", heading)
	doc = strings.ReplaceAll(doc, "{{rawHTML}}", rawHTML)
	return doc
}
