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

	req, err := http.NewRequest("DELETE", "https://www.onurix.com/api/v1/group/delete?key=AQUI_SU_KEY&client=AQUI_SU_ID&id=AQUI_ID_GROUP", nil)
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