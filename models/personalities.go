package models

type Personality struct {
	Id      int    `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	History string `json:"history" bson:"history"`
}

var Personalities []Personality
