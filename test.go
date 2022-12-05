package main

import (
	"fmt"
	"os"

	"github.com/devans10/pugo/pure1"
)

func main() {
	appID := os.Getenv("PURE1_APPID")
	privateKey := []byte(os.Getenv("PURE1_PRIVATEKEY"))
	restVersion := ""

	client := pure1.NewClient(appID, privateKey, restVersion)

	array, _ := client.Array.GetArray(nil)
	fmt.Println("ID: ", array.ID)
	fmt.Println("Name: ", array.Name)
	fmt.Println("Model: ", array.Model)
	fmt.Println("OS: ", array.OS)
	fmt.Println("Version: ", array.Version)
	fmt.Println("AsOf: ", array.AsOf)
}
