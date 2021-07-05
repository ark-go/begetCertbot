package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/antonholmquist/jason"
)

// запрос на сервер
func sendRequest(req string) (*jason.Object, error) {

	log.Println("Запрос:", req)

	if req == "" {
		//log.Println("Ошибка при составлении запроса")
		return nil, fmt.Errorf("SendRequest: %s", "Ошибка при составлении запроса")
	}
	resp, err := http.Get(req)
	if err != nil {
		//log.Println("Error-1, Ошибка связи:", err) // ошибки передачи но не ошибки сервера например 500 здесь не будет
		return nil, fmt.Errorf("SendRequest: %s: %v", "ошибка связи", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Println("Error-2", err)
		return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа", err)
	}

	v, _ := jason.NewObjectFromBytes(body)
	var otvOb *jason.Object
	var otv string
	if otv, err = v.GetString("status"); err != nil {
		return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа", err)
	}
	if otv != "success" {
		if otv, err = v.GetString("error_text"); err != nil {
			return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа", err)
		}
		return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа", otv)

	}
	if otv, err = v.GetString("answer", "status"); err != nil {
		return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа2", err)
	}
	if otv != "success" {
		if otv, err = v.GetString("error_text"); err != nil {
			return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа", err)
		}
		return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа", otv)

	}
	if otvOb, err = v.GetObject("answer", "result"); err != nil {
		return nil, fmt.Errorf("SendRequest: %s: %v", "Ошибка ответа3", err)
	}
	log.Println(otvOb)
	//printPrettier(otvOb)
	return otvOb, nil

	// printPrettier2(result)
}
