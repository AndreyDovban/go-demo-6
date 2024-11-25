package jwt_test

import (
	"go-demo-6/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJWT("SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Is not valid")
	}

	if data.Email != email {
		t.Fatalf("Email %s not equal %s ", email, data.Email)
	}
}
