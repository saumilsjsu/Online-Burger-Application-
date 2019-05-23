package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = os.Getenv("SERVER")
var mongodb_database = os.Getenv("DATABASE")
var mongodb_collection = os.Getenv("COLLECTION")
var mongo_admin_database = os.Getenv("ADMIN_DATABASE")
var mongo_username = os.Getenv("USERNAME")
var mongo_password = os.Getenv("PASSWORD")

func init() {
	fmt.Println("Mongo Server: ", mongodb_server)
	fmt.Println("Mongo DB :", mongodb_database)
	fmt.Println("Mongo Collection:", mongodb_collection)
	fmt.Println("Mongo User:", mongo_username)

}

func pingHandler(w http.ResponseWriter, req *http.Request) {
	mapD := map[string]string{"message": "API Working"}
	mapB, _ := json.Marshal(mapD)
	ResponseWithJSON(w, mapB, http.StatusOK)
	return
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	mapD := map[string]string{"message": message}
	mapB, _ := json.Marshal(mapD)
	ResponseWithJSON(w, mapB, code)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

/*
Handler method for getting all restaurants based on a ziplocation
*/
func getRestaurantHandler(w http.ResponseWriter, req *http.Request) {

	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(mongodb_database).C(mongodb_collection)

	params := mux.Vars(req)
	var zipcode string = params["zipcode"]
	fmt.Println(zipcode)

	var res []restaurant
	err = collection.Find(bson.M{"zipcode": zipcode}).All(&res)

	if res == nil || len(res) <= 0 || err != nil {
		ErrorWithJSON(w, "Cannot find any restaurants for that zipcode", http.StatusNotFound)
		return
	} else {
		fmt.Println("Result: ", res)
		respBody, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	}

}

/*
Handler to get all restaurant
*/
func getAllRestaurantHandler(w http.ResponseWriter, req *http.Request) {
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(mongodb_database).C(mongodb_collection)

	var res []restaurant
	err = collection.Find(bson.M{}).All(&res)

	if res == nil || len(res) <= 0 || err != nil {
		ErrorWithJSON(w, "Cannot find any restaurants for that zipcode", http.StatusNotFound)
		return
	} else {
		fmt.Println("Result: ", res)
		respBody, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	}

}

func addRestaurantHandler(w http.ResponseWriter, req *http.Request) {

	uuidForRestaurant, _ := uuid.NewV4()
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(mongodb_database).C(mongodb_collection)

	var res restaurant
	_ = json.NewDecoder(req.Body).Decode(&res)
	res.RestaurantId = uuidForRestaurant.String()
	fmt.Println("Restaurants: ", res)
	err = collection.Insert(res)
	if err != nil {
		ErrorWithJSON(w, "Cannot find any restaurants for that zipcode", http.StatusNotFound)
		return
	}

	respBody, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	ResponseWithJSON(w, respBody, http.StatusOK)
}

/*
Handler method for getting restaurant based on a Id
*/
func getRestaurantByIDHandler(w http.ResponseWriter, req *http.Request) {
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(mongodb_database).C(mongodb_collection)

	params := mux.Vars(req)
	var restaurantId string = params["restaurantId"]
	fmt.Println("restaurant id is : ", restaurantId)

	var res restaurant
	err = collection.Find(bson.M{"restaurantid": restaurantId}).One(&res)

	if err != nil {
		ErrorWithJSON(w, "Cannot find any restaurants for that id", http.StatusNotFound)
		return
	} else {
		fmt.Println("Result: ", res)
		respBody, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	}

}
func deleteRestaurantHandler(w http.ResponseWriter, req *http.Request) {

	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)

	params := mux.Vars(req)

	var result restaurant
	err = c.Find(bson.M{"restaurantid": params["restaurantId"]}).One(&result)
	if err == nil {
		c.Remove(bson.M{"restaurantid": params["restaurantId"]})
		respBody, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	} else {
		ErrorWithJSON(w, "Cannot find any restaurants for that id for deletion", http.StatusNotFound)
		return
	}

}

func main() {
	mx := mux.NewRouter()
	mx.HandleFunc("/restaurant/ping", pingHandler).Methods("GET")
	mx.HandleFunc("/restaurant", getAllRestaurantHandler).Methods("GET")
	mx.HandleFunc("/restaurant", addRestaurantHandler).Methods("POST")
	mx.HandleFunc("/restaurant/{restaurantId}", getRestaurantByIDHandler).Methods("GET")
	mx.HandleFunc("/restaurant/{restaurantId}", deleteRestaurantHandler).Methods("DELETE")
	mx.HandleFunc("/restaurant/zipcode/{zipcode}", getRestaurantHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", mx))
}
