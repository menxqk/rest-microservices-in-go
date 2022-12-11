package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func NewLocalStaticHandler(dir string) http.Handler {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	baseDir := filepath.Join(workDir, dir)
	return &localStaticHandler{baseDir}
}

type localStaticHandler struct {
	baseDir string
}

func (h *localStaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("not allowed"))
		return
	}

	upath := path.Clean(r.URL.Path)

	if upath == "/" || upath == "/index.html" {
		h.sendFile("index.html", w)
		return
	}

	f, err := os.Open(filepath.Join(h.baseDir, upath))
	if err != nil {
		if os.IsNotExist(err) {
			h.send404(w)
			return
		}
		fmt.Println(err)
		h.send500(w)
		return
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		h.send500(w)
		return
	}

	if info.IsDir() {
		h.send404(w)
		return
	}

	http.ServeFile(w, r, filepath.Join(h.baseDir, upath))
}

func (h *localStaticHandler) sendFile(name string, w http.ResponseWriter) {
	f, err := os.Open(filepath.Join(h.baseDir, name))
	if err != nil {
		if os.IsNotExist(err) {
			h.send404(w)
			return
		}
		fmt.Println(err)
		h.send500(w)
		return
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Println(err)
		h.send500(w)
		return
	}
}

func (h *localStaticHandler) send404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func (h *localStaticHandler) send500(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}
