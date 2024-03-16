package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestEncoderDecoder(t *testing.T) {
	newUser := &User{
		Id:       4,
		Login:    "Alice",
		Password: "198328",
		Access:   0,
	}

	jsonData, err := json.Marshal(newUser)
	if err != nil {
		log.Fatalf("Error during marshal json %v\n", err)
	}
	fmt.Println("JSON  DATA:", string(jsonData))

	token := base64Encode(jsonData)
	fmt.Println("TOKEN IS: ", token)

	newJsonData, err := base64Decode(token)
	if err != nil {
		log.Fatalf("Error during decode base64 %v\n", err)
	}
	fmt.Println("JSON  DATA:", string(newJsonData))
}
