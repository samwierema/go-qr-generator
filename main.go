package main

import (
	"code.google.com/p/rsc/qr"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/", QrGenerator)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func QrGenerator(w http.ResponseWriter, r *http.Request) {
	es := r.URL.Query().Get("s")
	if es == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	s, err := url.QueryUnescape(es)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}

	log.Print(s)

	code, err := qr.Encode(s, qr.L)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

	png := code.PNG()

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(png)))

	if _, err := w.Write(png); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
