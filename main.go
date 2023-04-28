package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	kcBaseUrl := os.Getenv("KC_BASEURL")
	kcRealm := os.Getenv("KC_REALM")
	kcUsername := os.Getenv("KC_USERNAME")
	kcPassword := os.Getenv("KC_PASSWORD")

	f, err := os.Open("users.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvR := csv.NewReader(f)

	data, err := csvR.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	client := gocloak.NewClient(kcBaseUrl)
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, kcUsername, kcPassword, kcRealm)

	if err != nil {
		log.Fatal(err)
	}

	users := parseKcUsers(data)

	for _, u := range users {

		userId, err := client.CreateUser(ctx, token.AccessToken, kcRealm, u)
		if err != nil {
			log.Fatalf("Error when creating user %s : %v", *u.Username, err)
		}
		err = client.SetPassword(ctx, token.AccessToken, userId, kcRealm, "password", false)

		if err != nil {
			log.Fatalf("Error when setting password to user %s : %v\n", userId, err)
		}
	}
}

func parseKcUsers(data [][]string) []gocloak.User {

	var users []gocloak.User

	for _, u := range data {
		users = append(users, gocloak.User{
			Username:      gocloak.StringP(u[0]),
			Email:         gocloak.StringP(u[1]),
			FirstName:     gocloak.StringP(u[2]),
			LastName:      gocloak.StringP(u[3]),
			RealmRoles:    parseRealmRole(u[4]),
			Enabled:       gocloak.BoolP(true),
			Totp:          gocloak.BoolP(true),
			EmailVerified: gocloak.BoolP(true),
		})
	}
	return users
}

func parseRealmRole(roles string) *[]string {
	realmRoles := strings.Split(roles, ",")
	return &realmRoles
}
