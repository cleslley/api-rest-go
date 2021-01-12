package model

import "gopkg.in/mgo.v2/bson"

//Movie armazena os dados do filme
type Movie struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Title string        `bson:"title" json:"title"`
	Year  string        `bson:"year" json:"year"`
}
