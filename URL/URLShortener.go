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

	data := strings.NewReader("client=AQUI_SU_CLIENT&key=AQUI_SU_KEY&name=AQUI_NOMBRE_DE_URL&url-long=AQUI_URL_LARGA&alias=OPCIONAL_ALIAS_URL&is-premium=OPCIONAL_AQUI_TRUE_OR_FALSE_DEFAULT_FALSE&group-name=OPCIONAL_AQUI_NOMBRE_DE_GRUPO&expiration-time-statistics=OPCIONAL_AQUI_TIEMPO_ALMACENAMIENTO_ESTADITICAS")
	req, err := http.NewRequest("POST", "https://www.onurix.com/api/v1/url/short", data)
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
