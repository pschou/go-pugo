package main

import (
	"os"

	"github.com/devans10/pugo/pure1"
)

func main() {
	appID := os.Getenv("PURE1_APPID")
	privateKey := []byte(os.Getenv("PURE1_PRIVATEKEY"))
	restVersion := ""

	client := pure1.NewClient(appID, privateKey, restVersion)
}
