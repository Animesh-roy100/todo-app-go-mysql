package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken creates a new token with the given id
// id: the id of the user
// returns: the token string and an error if any
func CreateToken(id uint32) (string, error) {
 claims := jwt.MapClaims{}
 claims["authorized"] = true
 claims["id"] = id
 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// TokenValid checks if the token is valid
// r: the http request
// returns: an error if any
func TokenValid(r *http.Request) error {
 tokenString := ExtractToken(r)
 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
   return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
  }
  return []byte(os.Getenv("API_SECRET")), nil
 })
 if err != nil {
  return err
 }
 if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
  Pretty(claims)
 }
 return nil
}

// ExtractToken extracts the token from the http request
// r: the http request
// returns: the token string
func ExtractToken(r *http.Request) string {
 keys := r.URL.Query()
 token := keys.Get("token")
 if token != "" {
  return token
 }
 bearerToken := r.Header.Get("Authorization")
 if len(strings.Split(bearerToken, " ")) == 2 {
  return strings.Split(bearerToken, " ")[1]
 }
 return ""
}

// ExtractTokenID extracts the token id from the http request
// r: the http request
// returns: the token id and an error if any
func ExtractTokenID(r *http.Request) (uint32, error) {
 tokenString := ExtractToken(r)
 token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
  if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
   return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
  }
  return []byte(os.Getenv("API_SECRET")), nil
 })
 if err != nil {
  return 0, err
 }
 claims, ok := token.Claims.(jwt.MapClaims)
 if ok && token.Valid {
  uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["id"]), 10, 32)
  if err != nil {
   return 0, err
  }
  return uint32(uid), nil
 }
 return 0, nil
}

// Pretty display the claims licely in the terminal
// data: the data to be displayed
func Pretty(data interface{}) {
 b, err := json.MarshalIndent(data, "", " ")
 if err != nil {
  log.Println(err)
  return
 }
 fmt.Println(string(b))
}