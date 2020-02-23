package firebase

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

// App firebase.
var App *firebase.App

// Client firebase.
var Client *auth.Client

func init() {
	opt := option.WithCredentialsFile("service_account.json")
	App, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing app: %v", err))
	}

	// Access auth service from the default app
	Client, err = App.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
}
