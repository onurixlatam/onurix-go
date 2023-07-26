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

	data := strings.NewReader("client=AQUI_SU_CLIENT&key=AQUI_SU_KEY&id=AQUI_ID_CONTACTO&name=OPCIONAL_AQUI_NOMBRE_CONTACTO&lastname=OPCIONAL_AQUI_APELLIDO_CONTACTO&email=OPCIONAL_AQUI_EMAIL_CONTACTO&phone=OPCIONAL_AQUI_TELEFONO_CONTACTO")
	req, err := http.NewRequest("POST", "https//www.onurix.com/api/v1/contacts/update", data)
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
