package handlers

import (
	"banking/data"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")

}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []data.Customer{
		{Name: "Rohan", PhoneNo: "+91-93245435435", Address: "Karnataka"},
		{Name: "Rohan1", PhoneNo: "+91-93245412321", Address: "Karnataka"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
