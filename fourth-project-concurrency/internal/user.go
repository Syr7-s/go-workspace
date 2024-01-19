package user

import (
	"encoding/json"
	"io"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetRandomUser() User {
	var user User

	resp, _ := http.Get("https://random-data-api.com/api/users/random_user")

	defer resp.Body.Close() // close the connection. memory leak if not closed
	bytes, _ := io.ReadAll(resp.Body)

	json.Unmarshal(bytes, &user)
	return user
}
