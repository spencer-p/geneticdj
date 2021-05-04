package main

import (
	"bytes"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/spencer-p/geneticdj/pkg/sexpr"
)

var (
	templates = template.Must(template.ParseFiles("generate.html", "play.html",
		"error.html", "index.html"))
	generator = sexpr.Preset{
		MaxDepth:      10,
		NumRange:      [2]int{1, 10},
		ArrayLenRange: [2]int{3, 7},
		OddsBranch:    0.75,
		OddsVar:       0.5,
	}
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", interesting_fns)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	fn := generator.Generate()
	templates.ExecuteTemplate(w, "generate.html", fn.String())
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	fn_text := r.URL.Path[len("/play/"):]
	templates.ExecuteTemplate(w, "play.html", fn_text)
}

func audioHandler(w http.ResponseWriter, r *http.Request) {
	fn_text := r.URL.Path[len("/api/audio/"):]

	// Parse the function
	fn, err := sexpr.Parse(fn_text)
	if err != nil {
		log.Println("Recvd unparseable function:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}

	// Write the audio to a buffer, making sure the fn is valid
	var wavbuf bytes.Buffer
	err = sexpr.WriteWav(fn, &wavbuf, 30)
	if err != nil {
		log.Println("Function could not be evaluated:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}

	// Write the audio
	w.Header().Set("Content-Type", "audio/x-wav")
	w.WriteHeader(http.StatusOK)
	if _, err := wavbuf.WriteTo(w); err != nil {
		log.Println("Error writing bytes of audio to http.ResponseWriter:", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/play/", playHandler)
	http.HandleFunc("/api/generate/", generateHandler)
	http.HandleFunc("/api/audio/", audioHandler)
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}
