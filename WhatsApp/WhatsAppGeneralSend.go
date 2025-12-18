package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://www.onurix.com/api/v1/whatsapp/send?key=AQUI_SU_KEY&client=AQUI_SU_CLIENT&template_id=AQUI_EL_ID_DE_LA_PLANTILLA&phone_sender_id=AQUI_EL_ID_DEL_NUMERO_DE_TELEFONO_REMITENTE"
	method := "POST"

	payload := strings.NewReader(`AQUI_EL_JSON_CON_LOS_VALORES_PARA_LA_PLANTILLA`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
