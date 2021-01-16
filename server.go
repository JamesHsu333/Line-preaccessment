package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

type SelfIntro struct {
	bot         *linebot.Client
	appBaseURL  string
	downloadDir string
}

func NewSelfIntro(channelSecret, channelToken string) (*SelfIntro, error) {
	bot, err := linebot.New(
		channelSecret,
		channelToken,
	)
	if err != nil {
		return nil, err
	}
	return &SelfIntro{
		bot:         bot,
		appBaseURL:  "nil",
		downloadDir: "nil",
	}, nil
}

func (s *SelfIntro) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := s.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		fmt.Printf("Got event %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println(message.Text)
				if err := s.handleText(message, event.ReplyToken, event.Source); err != nil {
					fmt.Println(err)
				}
			/*case *linebot.ImageMessage:
				if err := app.handleImage(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.VideoMessage:
				if err := app.handleVideo(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.AudioMessage:
				if err := app.handleAudio(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.FileMessage:
				if err := app.handleFile(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				if err := app.handleLocation(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				if err := app.handleSticker(message, event.ReplyToken); err != nil {
					log.Print(err)
				}*/
			default:
				log.Printf("Unknown message: %v", message)
			}
		/*case linebot.EventTypeFollow:
			if err := app.replyText(event.ReplyToken, "Got followed event"); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeUnfollow:
			log.Printf("Unfollowed this bot: %v", event)
		case linebot.EventTypeJoin:
			if err := app.replyText(event.ReplyToken, "Joined "+string(event.Source.Type)); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		case linebot.EventTypePostback:
			data := event.Postback.Data
			if data == "DATE" || data == "TIME" || data == "DATETIME" {
				data += fmt.Sprintf("(%v)", *event.Postback.Params)
			}
			if err := app.replyText(event.ReplyToken, "Got postback: "+data); err != nil {
				log.Print(err)
			}
		case linebot.EventTypeBeacon:
			if err := app.replyText(event.ReplyToken, "Got beacon: "+event.Beacon.Hwid); err != nil {
				log.Print(err)
			}*/
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}

func (s *SelfIntro) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	case "profile":
		if source.UserID != "" {
			profile, err := s.bot.GetProfile(source.UserID).Do()
			if err != nil {
				return s.replyText(replyToken, err.Error())
			}
			if _, err := s.bot.ReplyMessage(
				replyToken,
				linebot.NewTextMessage("Display name: "+profile.DisplayName),
				linebot.NewTextMessage("Status message: "+profile.StatusMessage),
			).Do(); err != nil {
				return err
			}
		} else {
			return s.replyText(replyToken, "Bot can't use profile API without user ID")
		}
	default:
		log.Printf("Echo message to %s: %s", replyToken, message.Text)
		if _, err := s.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(message.Text),
		).Do(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SelfIntro) replyText(replyToken, text string) error {
	if _, err := s.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app, err := NewSelfIntro(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	http.HandleFunc("/callback", app.Callback)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
