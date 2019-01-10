package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	lineChannelSecret := os.Getenv("CHANNEL_SECRET")
	lineChannelToken := os.Getenv("CHANNEL_TOKEN")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`
		CREATE TABLE IF NOT EXISTS destination (
			id TEXT NOT NULL PRIMARY KEY
		);
	`)

	bot, err := linebot.New(lineChannelSecret, lineChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			db.Exec(`INSERT INTO destination VALUES ($1)`, event.Source.RoomID)
		}
	})

	type sendMessageRequest struct {
		Message string `json:"message"`
	}

	http.HandleFunc("/send-message", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var reqJSON sendMessageRequest
		if err := json.NewDecoder(req.Body).Decode(&reqJSON); err != nil {
			log.Println(err)
			return
		}

		message := reqJSON.Message

		rows, err := db.Query(`SELECT id FROM destination`)
		if err != nil {
			log.Println(err)
			return
		}

		for rows.Next() {
			var id string
			if err := rows.Scan(&id); err != nil {
				log.Println(err)
				continue
			}

			bot.PushMessage(id, linebot.NewTextMessage(message))
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
