package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	headers := map[string][]string{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"Accept":       []string{"application/json"},
	}

	req, err := http.NewRequest("GET", "https://www.onurix.com/api/v1/balance?client=a3f3f9ea627d496c57e2cd354bb55280354b5be26491dcc974e18&key=3", nil)
	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
