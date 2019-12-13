package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// enviar requests al servidor a razón de 1 request cada 10 milisegundos
	// (aproximado por el scheduler del runtime)
	ticker := time.NewTicker(10 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				doRequest()
			}
		}
	}()

	// ejecutar la aplicación por 13 segundos (< server)
	time.Sleep(13 * time.Second)
	ticker.Stop()
	done <- true

	fmt.Println("Stopping client")
}

func doRequest() {
	resp, err := http.Get("http://localhost:5050/ping")

	if err != nil {
		fmt.Sprintf("error: %+v", err.Error())
	}
	if resp != nil {
		if resp.Body != nil {
			defer resp.Body.Close()
			body, readErr := ioutil.ReadAll(resp.Body)
			if readErr != nil {
				fmt.Sprintf("response: %s", string(body))
			}
		}
	}
}
