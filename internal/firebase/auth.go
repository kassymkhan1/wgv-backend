package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"path/filepath"
	"strings"
)

type FIREBASE interface {
	AuthMiddleware() string
	SetupFirebase() *auth.Client
}
type Fire struct {
	Token string
	Auth  string
}

func (f *Fire) AuthMiddleware() string {
	firebaseAuth := f.SetupFirebase()
	idToken := strings.TrimSpace(strings.Replace(f.Token, "Bearer", "", 1))

	if idToken == "" {
		return ""
	}
	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return ""
	}
	return token.UID
}
func (f *Fire) SetupFirebase() *auth.Client {

	serviceAccountKeyFilePath, err := filepath.Abs(f.Auth)
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic("Firebase load error")
	}
	return auth
}
