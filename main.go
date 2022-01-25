package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func returnMD5Hash(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: md5")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Submitted plaintext: %+v\n", string(reqBody))

	h := md5.Sum([]byte(reqBody))
	fmt.Fprintf(w, "%x\n", string(h[:]))
}

func returnSHA1Hash(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: sha1")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Submitted plaintext: %+v\n", string(reqBody))

	h := sha1.New()
	h.Write([]byte(reqBody))
	text := h.Sum(nil)
	fmt.Fprintf(w, "%x\n", text)
}

func returnSHA256Hash(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: sha256")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Submitted plaintext: %+v\n", string(reqBody))

	h := sha256.New()
	h.Write([]byte(reqBody))
	text := h.Sum(nil)
	//	json.NewEncoder(w).Encode(fmt.Sprintf("%x", text))
	fmt.Fprintf(w, "%x\n", text)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoints available: /md5, /sha1, /sha256\n")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(port int) {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/md5", returnMD5Hash).Methods("POST")
	router.HandleFunc("/sha1", returnSHA1Hash).Methods("POST")
	router.HandleFunc("/sha256", returnSHA256Hash).Methods("POST")

	fmt.Printf("waiting for connections on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), router); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

}

func main() {

	port := 9000
	handleRequests(port)

}
