/*
	Burger Menu Item API
*/

package main

type MenuItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type RestaurantMenu struct {
	RestaurantId string     `json:"resId"`
	Items        []MenuItem `json:"items"`
}

// type menuItem struct {
// 	Id          string
// 	Name        string
// 	Price       int
// 	Description string
// 	Calories    int
// }

type restaurantReqBody struct {
	RestaurantId string   `json:"resId"`
	Item         MenuItem `json:"item"`
}

type deleteReqBody struct {
	RestaurantId string `json:"resId"`
	ItemId       string `json:"itemId`
}
