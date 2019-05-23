package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = os.Getenv("MONGO_SERVER")
var mongodb_database = os.Getenv("MONGO_DATABASE")
var mongodb_collection = os.Getenv("MONGO_COLLECTION")
var mongo_user = os.Getenv("MONGO_USERNAME")
var mongo_pass = os.Getenv("MONGO_PASS")

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/carts/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/carts/{cartid}", cartHandler(formatter)).Methods("GET")
	mx.HandleFunc("/carts", createCartHandler(formatter)).Methods("POST")
	mx.HandleFunc("/carts", updateCartHandler(formatter)).Methods("PUT")

}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"CART API version 1.0 alive!"})
	}
}

// API CART INFO
func cartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		cartid := vars["cartId"]

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var result Cart
		if err = c.FindId(bson.ObjectIdHex(cartid)).One(&result); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		formatter.JSON(w, http.StatusOK, result)
	}
}

// API REGISTER NEW CART
func createCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var newCart Cart
		if err := json.NewDecoder(req.Body).Decode(&newCart); err != nil {
			formatter.JSON(w, http.StatusBadRequest, "Invalid Request")
			return
		}

		fmt.Println(newCart)

		newCart.CartID = bson.NewObjectId()

		if err := c.Insert(&newCart); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		var result Cart
		if err = c.FindId(newCart.CartID).One(&result); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Println(result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

//API UPDATE CART

func updateCartHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)

		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		var newCart Cart
		if err := json.NewDecoder(req.Body).Decode(&newCart); err != nil {
			formatter.JSON(w, http.StatusBadRequest, "Invalid Request")
			return
		}

		if err := c.UpdateId(newCart.CartID, &newCart); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		var updatedCart Cart
		if err := c.FindId(newCart.CartID).One(&updatedCart); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		formatter.JSON(w, http.StatusOK, updatedCart)

	}
}
