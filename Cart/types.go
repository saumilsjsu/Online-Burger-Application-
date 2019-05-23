package main

import "gopkg.in/mgo.v2/bson"

type Items struct {
	ItemId      string `bson:"itemId" json:"itemId"`
	ItemName    string `bson:"itemName" json:"itemName"`
	Price       string `bson:"price" json:"price"`
	Description string `bson:"description" json:"description"`
}

type Cart struct {
	CartID bson.ObjectId `bson:"_id" json:"cartId"`
	Total  string        `bson:"Total" json:"Total"`
	Items  []Items       `bson:"Products" json:"Products"`
}
