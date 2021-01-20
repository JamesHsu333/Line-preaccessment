package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
	"github.com/joho/godotenv"
)

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter
)

func ParseMessage(sentence string) string {
	hmm := seg.CutSearch(sentence, true)
	for _, str := range hmm {
		if value, exist := keywords[str]; exist {
			return value
		}
	}
	return "default"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	app, err := NewSelfIntro(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	seg.LoadDict()

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", app.Callback)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
