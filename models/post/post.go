package post

import (
	"encoding/json"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Post model
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

// EmptyPost is used for comparing empty Post values
var EmptyPost = Post{}

// IsValidPost validates Post coming from client
func IsValidPost(p Post) bool {
	if p.ID == "" || p.Title == "" {
		return false
	}
	return true
}

// FromJSON creates a Post out of JSON
func FromJSON(jsonPost []byte) (Post, error) {
	var newPost Post
	err := json.Unmarshal(jsonPost, &newPost)
	if err != nil {
		errMessage := "models.PostFromJSON > unmarshal error: " + err.Error()
		return Post{}, errors.New(errMessage)
	}

	return newPost, nil
}

func GeneratePost(partialPost Post) (Post, error) {
	partialPost.ID = uuid.Must(uuid.NewV4(), nil).String()
	partialPost.CreatedAt = time.Now()

	return partialPost, nil
}