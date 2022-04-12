package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	_ "embed"
)

//go:embed index.html
var indexpage []byte
var geometry []byte

func index(w http.ResponseWriter, r *http.Request) {
	w.Write(indexpage)
}

func screenshot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png") // normal header
	cmd := exec.Command("grim", "-g", string(geometry), "-")
	cmd.Stdout = w
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

func main() {
	var err error
	geometry, err = io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	geometry = bytes.TrimSpace(geometry)

	http.HandleFunc("/", index)
	http.HandleFunc("/screenshot", screenshot)

	exec.Command("xdg-open", "http://127.0.0.1:5050/").Start()
	if err := http.ListenAndServe("127.0.0.1:5050", nil); err != nil {
		log.Fatal(err)
	}
}
