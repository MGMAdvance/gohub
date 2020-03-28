package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Github username's info
type Github struct {
	Username string `json:"login"`
	ID       int    `json:"id"`
	Created  string `json:"created_at"`
	Updated  string `json:"updated_at"`
}

// Search Search and save username's info
func (g Github) Search(name string) {
	url := fmt.Sprintf("https://api.github.com/users/" + name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record Github

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("ID          = ", record.ID)
	fmt.Println("Username    = ", record.Username)
	fmt.Println("Created At  = ", record.Created)
	fmt.Println("Updated At  = ", record.Updated)

	g.insertInfos(record)
}

func (g Github) insertInfos(git Github) Github {
	g.ID = git.ID
	g.Username = git.Username
	g.Created = git.Created
	g.Updated = git.Updated

	return g
}

// GetID Get the username's ID
func (g Github) GetID() int {
	return g.ID
}

// GetUsername Get the username's Username
func (g Github) GetUsername() string {
	return g.Username
}

// GetCreated Get the username's Created
func (g Github) GetCreated() string {
	return g.Created
}

// GetUpdated Get the username's Uṕdated
func (g Github) GetUpdated() string {
	return g.Updated
}
