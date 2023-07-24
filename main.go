package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	viaCep := make(chan string)
	apiCep := make(chan string)

	go func() {
		viaCep <- getResponse("http://viacep.com.br/ws/08330380/json/")
	}()

	go func() {
		apiCep <- getResponse("https://cdn.apicep.com/file/apicep/08330-380.json")
	}()

	select {
	case msg1 := <-viaCep:
		fmt.Printf("received from viaCep: %s\n", msg1)
	case msg2 := <-apiCep:
		fmt.Printf("received from apiCep: %s\n", msg2)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

func getResponse(url string) string {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return "ocorreu um erro na chamada da api " + url
	}
	defer resp.Body.Close()
	responseString, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Erro ao decodificar a resposta JSON da api " + url
	}
	return string(responseString)
}
