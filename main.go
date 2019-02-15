package main

import (
	"net/http"
	"os/exec"
)

func main() {
	http.Handle("/api", http.HandlerFunc(imageHandler))
	staticHandler := http.FileServer(http.Dir("tmpl"))
	http.Handle("/", staticHandler)
	http.ListenAndServe(":8000", nil)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	in, header, _ := r.FormFile("file")
	mimeType := header.Header.Get("Content-Type")
	if mimeType != "image/png" && mimeType != "image/jpeg" {
		response := "Unexpected MIME type: " + mimeType
		w.WriteHeader(400)
		w.Write([]byte(response))
		return
	}
	// Run what amounts to convert | base64 | tr. The convert arguments are stolen from https://tex.stackexchange.com/a/94541/88992
	// TODO: Figure out if there's a way to exploit passing the POST data directly to convert's STDIN
	c1 := exec.Command("convert", "-", "-colorspace", "gray", "(", "+clone", "-blur", "0x1", ")", "+swap", "-compose", "divide", "-composite", "-linear-stretch", "5%x0%", "-rotate", "1.5", "-")
	// Note: base64 --wrap=0 would allow us to skip tr, but the version of base64 shipped with
	// alpine doesn't have this option.
	c2 := exec.Command("base64")
        c3 := exec.Command("tr", "-d", "\n")
	c1.Stdin = in
        c2.Stdin, _ = c1.StdoutPipe()
        c3.Stdin, _ = c2.StdoutPipe()
        c3.Stdout = w
        c3.Start()
        c2.Start()
        c1.Start()
        c3.Wait()
}
