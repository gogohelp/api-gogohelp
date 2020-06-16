package firebase

import (
	"log"
	"os"
	"path/filepath"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var DB *firestore.Client

func Init() *firestore.Client {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("WFT %v\n", err)
	}
	filepath := filepath.Join(dir, "/serviceAccountKey.json")
	opt := option.WithCredentialsFile(filepath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	DB, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return DB
}
