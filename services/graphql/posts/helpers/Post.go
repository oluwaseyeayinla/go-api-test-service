package helpers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zalbiraw/go-api-test-service/services/graphql/posts/graph/model"
)

var posts []*model.Post

func LoadPosts() error {
	byteValue, err := ioutil.ReadFile("/go/src/github.com/zalbiraw/go-api-test-service/helpers/posts-db.json")

	if nil != err {
		return err
	}

	json.Unmarshal(byteValue, &posts)

	return nil
}

func GetPosts() []*model.Post {
	return posts
}
