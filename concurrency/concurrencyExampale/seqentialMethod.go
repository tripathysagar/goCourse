package main

import (
	"fmt"
	"log"
	"time"
)

type UserProfile struct {
	UserName  string
	Tweets    []string
	Followers []string
	Following []string
}

func main() {
	start := time.Now()

	userProfile, err := handleGetUserProfile()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user profile details are :\n %+v\n", userProfile)

	fmt.Println("timetaken for user profile is ", time.Since(start))
}

func handleGetUserProfile() (*UserProfile, error) {
	UserName, err := getUserName()
	if err != nil {
		return &UserProfile{}, err
	}

	Tweets, err := getTweets()
	if err != nil {
		return &UserProfile{}, err
	}

	Followers, err := getFollowers()
	if err != nil {
		return &UserProfile{}, err
	}

	Following, err := getFollowing()
	if err != nil {
		return &UserProfile{}, err
	}

	userProfile := &UserProfile{
		UserName:  UserName,
		Tweets:    Tweets,
		Followers: Followers,
		Following: Following,
	}

	return userProfile, nil
}

func getUserName() (string, error) {
	time.Sleep(time.Millisecond * 200)
	return "FooBar", nil
}

func getTweets() ([]string, error) {
	time.Sleep(time.Millisecond * 100)
	return []string{
		"i am pain",
		"how are you doing",
		"AI will take over the world!!",
	}, nil
}

func getFollowing() ([]string, error) {
	time.Sleep(time.Millisecond * 150)
	return []string{
		"Naruto",
		"Madara",
		"Sasori",
		"Daidara",
	}, nil
}

func getFollowers() ([]string, error) {
	time.Sleep(time.Millisecond * 175)
	return []string{
		"gorillaKing",
		"monkeyKing",
		"testme",
		"Nagato",
	}, nil
}
