package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RequestObj struct {
	TotalPage int            `json:"total"`
	Data      []CustomerResp `json:"data"`
}
type CustomerResp struct {
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func DataCustomerInit() ([]CustomerResp, error) {
	var custResps []CustomerResp
	var i = 1
	var totalPages = 1

	for i <= totalPages {
		var custResp RequestObj
		url := "https://reqres.in/api/users?page=" + strconv.Itoa(i)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		//fmt.Println(string(body))
		if err := json.Unmarshal(body, &custResp); err != nil {
			return nil, err
		}
		totalPages = custResp.TotalPage
		i += 1
		custResps = append(custResps, custResp.Data...)
	}
	return custResps, nil
}
