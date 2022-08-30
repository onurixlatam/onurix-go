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

	data := strings.NewReader("client=AQUI_SU_CLIENT&key=AQUI_SU_KEY&template=AQUI_EL_NOMBRE_DE_LA_PLANTILLA&content=AQUI_EL_JSON_CON_LOS_VALORES_PARA_LA_PLANTILLA")
	req, err := http.NewRequest("POST", "https://www.onurix.com/api/v1/whatsapp/send", data)
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
