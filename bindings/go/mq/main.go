package main

//dependencies
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)
type Order struct {
	OrderId string `json:"orderId"`
}
//code
func getCheckout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	log.Println("Received Message: ", order.OrderId)
	time.Sleep(time.Second * 10)
	log.Println("Received Message Done: ", order.OrderId)
	if err != nil {
		log.Printf("error parsing checkout input binding payload: %s", err)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func main() {
	env, _ := os.LookupEnv("port")
	var port int
	port, _ = strconv.Atoi(env)
	r := mux.NewRouter()
	//here is important
	r.HandleFunc("/bind-route", getCheckout).Methods("POST", "OPTIONS")
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
