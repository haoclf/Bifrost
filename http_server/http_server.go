package main

import (
	"net/http"
	"log"
)

func handel_data(w http.ResponseWriter,req *http.Request){
	switch req.Method {
	case "GET":
		check_uri()
		break
	case "POST":
		post(w,req)
		break
	default:
		log.Println("Methon:",req.Method," not supported ")
		break
	}
	w.Write([]byte("success"))
}

func check_uri()  {
	log.Println("check uri success")
	return
}

func post(w http.ResponseWriter,req *http.Request)  {
	req.ParseForm()
	log.Println("EventType",req.Form.Get("EventType"))
	log.Println("SchemaName",req.Form.Get("SchemaName"))
	log.Println("TableName",req.Form.Get("TableName"))
	log.Println("Data",req.Form.Get("data"))
	return
}

func main()  {
	http.HandleFunc("/bifrost_http_api_test",handel_data)
	http.ListenAndServe("0.0.0.0:3332", nil)

}
