package data

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	PhoneNo string `json:"phone_number" xml:"phoneNo"`
	Address string `json:"address" xml:"address"`
}
