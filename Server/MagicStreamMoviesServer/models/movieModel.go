package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct {
	ID   bson.ObjectID `json:"id"`
	Name string        `json:"name"`
}

type Movie struct {
	ID          bson.ObjectID `json:"id"`
	ImdbID      string        `json:"imdb_id"`
	Title       string        `json:"title"`
	PosterPath  string        `json:"poster_path"`
	YoutubeID   string        `json:"youtube_id"`
	Genres      []Genre       `json:"genres"`
	Description string        `json:"description"`
	Year        int           `json:"year"`
}
