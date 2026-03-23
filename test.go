package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request){
	date:=r.URL.Query().Get("date_req")
	errFlag:= r.URL.Query().Get("error")
	usd:= r.URL.Query().Get("usd")

	if date == ""{
		date = time.Now().Format("02/01/2006")
	}

	if errFlag == "true"{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(w)
	}

	if usd == ""{
		usd = "70.20"
	}

	w.Header().Set("Content-Type", "application/xml; charset=windows-1251")
	response:= fmt.Sprintf(`<?xml version="1.0" encoding="windows-1251"?>
<ValCurs Date="%s" name="Foreign Currency Market">
  <Valute ID="R01235">
    <NumCode>840</NumCode>
    <CharCode>USD</CharCode>
    <Nominal>1</Nominal>
    <Name>Доллар США</Name>
    <Value>%s</Value>
  </Valute>
  <Valute ID="R01239">
    <NumCode>978</NumCode>
    <CharCode>EUR</CharCode>
    <Nominal>1</Nominal>
    <Name>Евро</Name>
    <Value>80.50</Value>
  </Valute>
</ValCurs>`, date, usd)
fmt.Fprintln(w, response)
}
func main(){
	fmt.Println("Запуск сервера")

	http.HandleFunc("/scripts/XML_daily.asp", handler)
	err:= http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Конец работы сервера")
}