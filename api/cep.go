package http

import (
	"fmt"
	"io"
	"net/http"
)

func Cep(cep string) (string, error) {

	url := "https://viacep.com.br/ws/" + cep + "/json/"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	bodyBites, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(bodyBites)

	return bodyString, err

}
