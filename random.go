package wrag

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mitchellh/mapstructure"
)

type Listing struct {
	Data
}

type Data struct {
	Children []Child
}

type Child struct {
	Kind      string
	ChildData `mapstructure:"data"`
}

type ChildData struct {
	Subreddit             string  `mapstructure:"subreddit"`
	Title                 string  `mapstructure:"title"`
	SubredditNamePrefixed string  `mapstructure:"subreddit_name_prefixed"`
	UpvoteRatio           float32 `mapstructure:"upvote_ratio"`
	Ups                   int     `mapstructure:"ups"`
	MediaURL              string  `mapstructure:"url"`
	ThreadURL             string  `mapstructure:"permalink"`
	NumComments           int     `mapstructure:"num_comments"`
}

func (l *Listing) Details() Child {
	return l.Children[0]
}

func randomURL(subreddit string) string {
	return fmt.Sprintf(endpoints.Apis["subreddit_random"], subreddit)
}

func Random(subreddit string) *Listing {
	req, err := http.NewRequest("GET", randomURL(subreddit), nil)
	req.Header.Add("Authorization", bearerToken(tokenResponse.AccessToken))
	req.Header.Add("User-Agent", config.Auth.UserAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Unable to fetch random subreddit.")
	}

	payload := []map[string]interface{}{}

	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		log.Fatal("Unable to parse listing from Reddit.")
		panic(err)
	}

	var l Listing
	err = mapstructure.Decode(payload[0], &l)
	check(err)
	return &l
}
