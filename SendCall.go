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

	data := strings.NewReader("client=AQUI_SU_CLIENT&key=AQUI_SU_KEY&phone=AQUI_EL_NUMERO_DE_CELULAR&message=AQUI_EL_MENSAJE_A_ENVIAR&voice=AQUI_TIPO_DE_VOZ&retries=AQUI_NUMERO_DE_INTENTOS&leave-voicemail=false&audio-code=7b9765ef-48b0-494c-8671-89ddb8f75fa4&country-code=CO")
	req, err := http.NewRequest("POST", "https://www.onurix.com/api/v1/call/send", data)
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
