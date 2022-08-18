package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	headers := map[string][]string{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"Accept":       []string{"application/json"},
	}

	data := strings.NewReader("client=AQUI_SU_CLIENT&key=AQUI_SU_KEY&phone=AQUI_EL_NUMERO_DE_CELULAR&name-url=AQUI_NOMBRE_DE_URL&since=Fecha inicial YYYY-MM-DD&until=Fecha final YYYY-MM-DD")
	req, err := http.NewRequest("POST", "https://www.onurix.com/api/v1/url/short-statistic", data)
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
