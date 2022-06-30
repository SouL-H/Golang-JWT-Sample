package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySignigcKey = []byte("testSecret")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8081/", nil)
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(body))
}
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "testUser"
	claims["exp"] = time.Now().Add(time.Minute * 24).Unix()
	tokenString, err := token.SignedString(mySignigcKey)

	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
		return "", err

	}
	return tokenString, nil
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {

	handleRequest()

}
