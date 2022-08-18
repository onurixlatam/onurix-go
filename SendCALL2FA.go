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

	data := strings.NewReader("client=AQUI_SU_CLIENT&key=AQUI_SU_KEY&phone=AQUI_EL_NUMERO_DE_CELULAR&app-name=AQUI_NOMBRE_APP&voice=AQUI_TIPO_DE_VOZ&retries=AQUI_NUMERO_DE_INTENTOS&country-code=CO")
	req, err := http.NewRequest("POST", "https://www.onurix.com/api/v1/call/2fa/send-call", data)
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
