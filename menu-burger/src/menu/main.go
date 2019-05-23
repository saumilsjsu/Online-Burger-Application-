/*
	Burger Menu Item API
*/

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

func createMenuItemHandler(w http.ResponseWriter, req *http.Request) {
	var reqPayload restaurantReqBody
	_ = json.NewDecoder(req.Body).Decode(&reqPayload)
	fmt.Println("Menu ItemPayload ", reqPayload.Item)
	uuid, _ := uuid.NewV4()
	reqPayload.Item.Id = uuid.String()
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
	mongo_collection := session.DB(mongodb_database).C(mongodb_collection)

	var menu RestaurantMenu
	err = mongo_collection.Find(bson.M{"restaurantid": reqPayload.RestaurantId}).One(&menu)
	if err != nil {
		fmt.Println("error: ", err)

		menu.RestaurantId = reqPayload.RestaurantId
		//menu.RestaurantName = reqPayload.RestaurantName
		menu.Items = append(menu.Items, reqPayload.Item)

		error := mongo_collection.Insert(menu)
		fmt.Println("error: ", error)
		if error != nil {
			ErrorWithJSON(w, "Internla server error", http.StatusInternalServerError)
			return
		}

	} else {
		menu.Items = append(menu.Items, reqPayload.Item)
		error := mongo_collection.Update(bson.M{"restaurantid": menu.RestaurantId}, bson.M{"$set": bson.M{"items": menu.Items}})
		if error != nil {
			fmt.Println("error: ", error)
			ErrorWithJSON(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	respBody, err := json.MarshalIndent(menu, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)

}

// API to find an item in the menu
func findRestaurantMenu(w http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	var restaurantId string = params["restaurantId"]
	fmt.Println("restaurant ID: ", restaurantId)
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	// session, err := mgo.Dial(mongodb_server)
	if err != nil {
		ErrorWithJSON(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer session.Close()

	mongo_collection := session.DB(mongodb_database).C(mongodb_collection)
	var result bson.M
	err = mongo_collection.Find(bson.M{"restaurantid": restaurantId}).One(&result)
	if err != nil {
		ErrorWithJSON(w, "Menu not found", http.StatusNotFound)
		return
	}
	respBody, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	ResponseWithJSON(w, respBody, http.StatusOK)
}

// API to update an items in the menu
func updateMenuItemHandler(w http.ResponseWriter, request *http.Request) {

	var reqPayload restaurantReqBody
	_ = json.NewDecoder(request.Body).Decode(&reqPayload)
	fmt.Println("Menu ItemPayload ", reqPayload.Item)
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		ErrorWithJSON(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	mongo_collection := session.DB(mongodb_database).C(mongodb_collection)

	var menu RestaurantMenu
	err = mongo_collection.Find(bson.M{"restaurantid": reqPayload.RestaurantId}).One(&menu)
	if err != nil {
		fmt.Println("error: ", err)
		ErrorWithJSON(w, "Restaurant not found", http.StatusNotFound)
		return
	} else {
		for i := 0; i < len(menu.Items); i++ {
			if menu.Items[i].Id == reqPayload.Item.Id {
				menu.Items[i].Name = reqPayload.Item.Name
				menu.Items[i].Price = reqPayload.Item.Price
				menu.Items[i].Description = reqPayload.Item.Description
				break
			}
		}
		error := mongo_collection.Update(bson.M{"restaurantid": menu.RestaurantId}, bson.M{"$set": bson.M{"items": menu.Items}})
		if error != nil {
			ErrorWithJSON(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	respBody, err := json.MarshalIndent(menu, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	ResponseWithJSON(w, respBody, http.StatusOK)

}

// API to delete an items in the menu
func deleteMenuItemHandler(w http.ResponseWriter, request *http.Request) {

	var reqPayload deleteReqBody
	_ = json.NewDecoder(request.Body).Decode(&reqPayload)
	fmt.Println("Menu ItemPayload ", reqPayload)
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		ErrorWithJSON(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	mongo_collection := session.DB(mongodb_database).C(mongodb_collection)

	var menu RestaurantMenu
	err = mongo_collection.Find(bson.M{"restaurantid": reqPayload.RestaurantId}).One(&menu)
	if err != nil {
		fmt.Println("error: ", err)
		ErrorWithJSON(w, "Restaurant not found", http.StatusNotFound)
		return
	} else {
		for i := 0; i < len(menu.Items); i++ {
			if menu.Items[i].Id == reqPayload.ItemId {
				menu.Items = append(menu.Items[0:i], menu.Items[i+1:]...)
				break
			}
		}
		error := mongo_collection.Update(bson.M{"restaurantid": menu.RestaurantId}, bson.M{"$set": bson.M{"items": menu.Items}})
		if error != nil {
			ErrorWithJSON(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	respBody, err := json.MarshalIndent(menu, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	ResponseWithJSON(w, respBody, http.StatusOK)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/menu/ping", pingHandler).Methods("GET")
	router.HandleFunc("/menu", createMenuItemHandler).Methods("POST")
	router.HandleFunc("/menu/{restaurantId}", findRestaurantMenu).Methods("GET")
	router.HandleFunc("/menu", updateMenuItemHandler).Methods("PUT")
	router.HandleFunc("/menu", deleteMenuItemHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8004", router))
}
