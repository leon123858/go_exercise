package utils

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var AuthClient *auth.Client
var AuthCtx context.Context

func InitFirebase() {
	var app *firebase.App
	var err error
	if os.Getenv("GO_ENV") == "release" {
		app, err = firebase.NewApp(context.Background(), nil)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	} else {
		opt := option.WithCredentialsFile("./key.json")
		app, err = firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
		}
	}
	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	FirebaseApp = app
	AuthClient = client
	AuthCtx = ctx
}
