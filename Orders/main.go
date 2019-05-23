package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
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

func burgerOrderStatus(w http.ResponseWriter, req *http.Request) {

	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err)
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	c := session.DB(mongodb_database).C(mongodb_collection)
	params := mux.Vars(req)
	var uuid string = params["orderId"]
	if uuid == "" {
		var orders_array []BurgerOrder
		err = c.Find(bson.M{}).All(&orders_array)
		fmt.Println("Burger Orders:", orders_array)
		respBody, err := json.MarshalIndent(orders_array, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	} else {
		fmt.Println("orderID: ", uuid)
		var result BurgerOrder
		err = c.Find(bson.M{"orderId": uuid}).One(&result)
		if err != nil {
			ErrorWithJSON(w, "Order not found", http.StatusNotFound)
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Burger Order: ", result)
		respBody, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	}

}

func burgerOrderStatusByUser(w http.ResponseWriter, req *http.Request) {

	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	c := session.DB(mongodb_database).C(mongodb_collection)
	params := mux.Vars(req)
	var uuid string = params["userId"]
	fmt.Println("userId: ", uuid)
	var result []BurgerOrder
	err = c.Find(bson.M{"userId": uuid}).All(&result)
	if err != nil {
		ErrorWithJSON(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	if len(result) == 0 {
		ErrorWithJSON(w, "No order for this user", http.StatusNotFound)
		return
	}
	_ = json.NewDecoder(req.Body).Decode(&result)
	fmt.Println("Burger Order: ", result)
	respBody, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)

}

// API Create New Burger Order
func burgerOrderHandler(w http.ResponseWriter, req *http.Request) {

	var orderdetail RequiredPayload
	_ = json.NewDecoder(req.Body).Decode(&orderdetail)
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		ErrorWithJSON(w, "Could not connect to database", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	var order BurgerOrder
	var newitem Items
	err = c.Find(bson.M{"orderId": orderdetail.OrderId}).One(&order)
	newitem.ItemId = orderdetail.ItemId
	newitem.ItemName = orderdetail.ItemName
	newitem.Price = orderdetail.Price
	newitem.Description = orderdetail.Description
	if err == nil {
		if order.OrderStatus == "Paid" {
			ErrorWithJSON(w, "Please create new order", http.StatusNotFound)
			return
		}
		order.Cart = append(order.Cart, newitem)
		order.TotalAmount = (order.TotalAmount + newitem.Price)
		fmt.Println("Orders: ", "Orders found")
		c.Update(bson.M{"orderId": orderdetail.OrderId}, bson.M{"$set": bson.M{"items": order.Cart, "totalAmount": order.TotalAmount}})
	} else {
		fmt.Println("Orders: ", "Orders not found")
		order = BurgerOrder{
			OrderId:     orderdetail.OrderId,
			UserId:      orderdetail.UserId,
			OrderStatus: "Placed",
			TotalAmount: newitem.Price,
			Cart: []Items{
				newitem,
			},
		}
		_ = json.NewDecoder(req.Body).Decode(&order)
		err = c.Insert(order)
		if err != nil {
			ErrorWithJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	respBody, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)

}

// API Paid Order
func burgerOrderPaid(w http.ResponseWriter, req *http.Request) {

	var paymentdetail RequiredPayload
	_ = json.NewDecoder(req.Body).Decode(&paymentdetail)
	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		ErrorWithJSON(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	c := session.DB(mongodb_database).C(mongodb_collection)
	params := mux.Vars(req)
	var uuid string = params["orderId"]
	fmt.Println(uuid)
	var orderpaid BurgerOrder
	err = c.Find(bson.M{"orderId": uuid}).One(&orderpaid)
	if err != nil {
		fmt.Println("Order not found")
		ErrorWithJSON(w, "Order not found", http.StatusNotFound)
		return
	}
	orderpaid.OrderStatus = "Paid"
	orderpaid.UserId = paymentdetail.UserId
	c.Update(bson.M{"orderId": uuid}, bson.M{"$set": bson.M{"orderStatus": orderpaid.OrderStatus, "userId": orderpaid.UserId}})
	fmt.Println("Order:", uuid, "paid")
	respBody, err := json.MarshalIndent(orderpaid, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)
}

// API Delete Item from Order
func burgerItemDelete(w http.ResponseWriter, req *http.Request) {

	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		ErrorWithJSON(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	var orderdetail RequiredPayload
	_ = json.NewDecoder(req.Body).Decode(&orderdetail)
	params := mux.Vars(req)
	var uuid string = params["orderId"]
	var result BurgerOrder
	fmt.Println("order ID: ", uuid)
	err = c.Find(bson.M{"orderId": uuid}).One(&result)
	if err != nil {
		fmt.Println("order not found")
		ErrorWithJSON(w, "Order not found", http.StatusNotFound)
		return
	}
	for i := 0; i < len(result.Cart); i++ {
		if result.Cart[i].ItemId == orderdetail.ItemId {
			result.TotalAmount = result.TotalAmount - result.Cart[i].Price
			result.Cart = append(result.Cart[0:i], result.Cart[i+1:]...)
			break
		}
	}
	c.Update(bson.M{"orderId": uuid}, bson.M{"$set": bson.M{"items": result.Cart, "totalAmount": result.TotalAmount}})
	fmt.Println("Delete Item: ", orderdetail.ItemId, "from order", uuid)
	respBody, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ResponseWithJSON(w, respBody, http.StatusOK)

}

// API Delete Burger Order
func burgerOrderDelete(w http.ResponseWriter, req *http.Request) {

	info := &mgo.DialInfo{
		Addrs:    []string{mongodb_server},
		Timeout:  60 * time.Second,
		Database: mongodb_database,
		Username: mongo_username,
		Password: mongo_password,
	}

	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	if err != nil {
		ErrorWithJSON(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodb_database).C(mongodb_collection)
	var orderdetail RequiredPayload
	_ = json.NewDecoder(req.Body).Decode(&orderdetail)
	fmt.Println("order ID: ", orderdetail.OrderId)
	err = c.Remove(bson.M{"orderId": orderdetail.OrderId})
	if err != nil {
		fmt.Println("order not found")
		ErrorWithJSON(w, "Order Not Found", http.StatusNotFound)
		return
	}
	ErrorWithJSON(w, "Order "+orderdetail.OrderId+" deleted", http.StatusInternalServerError)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/order/ping", pingHandler).Methods("GET")
	router.HandleFunc("/order", burgerOrderStatus).Methods("GET")
	router.HandleFunc("/order/{orderId}", burgerOrderStatus).Methods("GET")
	router.HandleFunc("/orders/{userId}", burgerOrderStatusByUser).Methods("GET")
	router.HandleFunc("/order", burgerOrderHandler).Methods("POST")
	router.HandleFunc("/order/{orderId}", burgerOrderPaid).Methods("PUT")
	router.HandleFunc("/order/{orderId}", burgerItemDelete).Methods("DELETE")
	router.HandleFunc("/order", burgerOrderDelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8003", router))
}
