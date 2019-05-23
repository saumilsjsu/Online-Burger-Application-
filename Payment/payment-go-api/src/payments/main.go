/*

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
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = os.Getenv("MONGO_SERVER")
var mongodb_database = os.Getenv("MONGO_DATABASE")
var mongodb_collection = os.Getenv("MONGO_COLLECTION")
var mongodb_username = os.Getenv("MONGO_USERNAME")
var mongodb_password = os.Getenv("MONGO_PASS")

type Payments []Payment

var payments []Payment

// NewServer configures and returns a Server.

func init() {
	fmt.Println("Mongo Server: ", mongodb_server)
	fmt.Println("Mongo DB :", mongodb_database)
	fmt.Println("Mongo Collection:", mongodb_collection)
	fmt.Println("Mongo User:", mongodb_username)

}

func pingHandler(w http.ResponseWriter, req *http.Request) {
	log.Print("hello")
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

func handleRequest() {

}

func getAllPayments(w http.ResponseWriter, req *http.Request) {
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongodb_username,
		Password: mongodb_password,
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
	var result []bson.M
	err = c.Find(nil).All(&result)
	if err != nil {
		ErrorWithJSON(w, "Get All Payment Error", http.StatusNotFound)
		return
	}
	fmt.Println("getAllPayments:", result)
	respBody, err := json.MarshalIndent(result, "", "  ")
	ResponseWithJSON(w, respBody, http.StatusOK)
}

func getPaymentByID(w http.ResponseWriter, req *http.Request) {
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongodb_username,
		Password: mongodb_password,
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
	var result bson.M
	params := mux.Vars(req)

	fmt.Printf("params[id]=%s \n", params["id"])

	err = c.Find(bson.M{"paymentid": params["id"]}).One(&result)
	if err != nil {
		ErrorWithJSON(w, "Get Payment by ID Error", http.StatusNotFound)
		return
	}
	fmt.Println("getPaymentByID:", result)
	respBody, err := json.MarshalIndent(result, "", "  ")
	ResponseWithJSON(w, respBody, http.StatusOK)
}

func createPayments(w http.ResponseWriter, req *http.Request) {
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongodb_username,
		Password: mongodb_password,
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

	var payment Payment
	_ = json.NewDecoder(req.Body).Decode(&payment)

	uuid, _ := uuid.NewV4()
	payment.PaymentID = uuid.String()
	t := time.Now()
	payment.PaymentDate = t.Format("2006-01-02 15:04:05")
	payment.Status = true

	err = c.Insert(payment)
	if err != nil {
		ErrorWithJSON(w, "Create Payment Error", http.StatusNotFound)
		return
	}
	fmt.Println("Create new payment:", payment)
	respBody, err := json.MarshalIndent(payment, "", "  ")
	ResponseWithJSON(w, respBody, http.StatusOK)
}

func main() {
	log.Print("hello")
	mx := mux.NewRouter()
	mx.HandleFunc("/payments/ping", pingHandler).Methods("GET")
	mx.HandleFunc("/payments", getAllPayments).Methods("GET")
	mx.HandleFunc("/payments/{id}", getPaymentByID).Methods("GET")
	mx.HandleFunc("/payments", createPayments).Methods("POST")
	log.Fatal(http.ListenAndServe(":8002", mx))
}
