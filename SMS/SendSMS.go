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

	// is-transactional es opcional. 1 (por defecto): mensaje transaccional. 0: mensaje comercial (publicidad, promociones);
	// los numeros inscritos en el Registro de Numeros Excluidos (RNE) que no aceptan SMS se rechazan con el error 1010.
	data := strings.NewReader("client=AQUI_SU_CLIENT&key=AQUI_SU_KEY&phone=AQUI_EL_NUMERO_DE_CELULAR&sms=AQUI_EL_SMS_A_ENVIAR&groups=AQUI_ID_GRUPOS&is-transactional=1")
	req, err := http.NewRequest("POST", "https://www.onurix.com/api/v1/sms/send", data)
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
