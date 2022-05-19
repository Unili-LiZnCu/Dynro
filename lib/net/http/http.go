package http

import (
	//"net"
	"io/ioutil"
	"os"
	//"time"

	"log"
	"net/http"
	//"os"
)

func InitLogger(){
	file , err := os.Create("Dynro.log")
	if file == nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	
	log.SetPrefix("[info]")
	log.Println("Logger Started.")
}

func HttpsGetAndSend(w http.ResponseWriter , r *http.Request) () {
	InitLogger()

	log.Println("Http Handler Started")
	if r.URL == nil{
		log.Panicln("empty pointer.")
	}
	req , err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	resp := req.Response
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	w.Write(body)
}

func Proxy(){
	//go http.ListenAndServe(":8081", http.HandlerFunc(http.FileServer(http.Dir("/")).ServeHTTP))
	err := http.ListenAndServeTLS(":443", "cacert/DYNRO.crt","cacert/DYNRO.key", http.HandlerFunc(HttpsGetAndSend))
	if err != nil {
		panic(err)
	}
	//http.ListenAndServe(":80", http.HandlerFunc(HttpsGetAndSend))
}