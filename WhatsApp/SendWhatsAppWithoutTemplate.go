package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://www.onurix.com/api/v1/whatsapp/send/no-template?key=AQUI_SU_SECRET_KEY&client=AQUI_SU_CLIENT_ID&phone_sender_id=AQUI_EL_ID_DEL_NUMERO_DE_TELEFONO_REMITENTE"
	method := "POST"

	payload := strings.NewReader(`{
    "from_phone_meta_id": "AQUI_EL_META_ID_DEL_TELEFONO",
    "phone": "AQUI_EL_TELEFONO_DESTINO",
    "message": {
        "type" : "text",
        "value": "AQUI_EL_MENSAJE"
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

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
