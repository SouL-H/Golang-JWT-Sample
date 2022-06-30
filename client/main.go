package main

import (
	"fmt"
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
	fmt.Fprintf(w, validToken)
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
func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {


handleRequest()

}
