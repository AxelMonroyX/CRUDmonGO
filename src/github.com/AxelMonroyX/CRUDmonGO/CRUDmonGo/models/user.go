package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	User struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Name   string        `json:"name" bson:"name"`
		Adress string       `json:"adress" bson:"adress"`
		Money  float32      `json:"money" bson:"money"`
		Age    int         `json:"age" bson:"age"`
	}
)