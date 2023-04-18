package main

import (
	"fmt"
	"log"
	"sync"
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
	userNamech := make(chan string, 1)
	tweetsch := make(chan []string, 1)
	followingch := make(chan []string, 1)
	followersch := make(chan []string, 1)

	wg := sync.WaitGroup{}

	// getting data using go routine only may cause issue during execution
	// if func returns soon the userprofile might not be populated
	// need waight group to be implemented
	go getUserName(userNamech, &wg)
	go getTweets(tweetsch, &wg)
	go getFollowers(followersch, &wg)
	go getFollowing(followingch, &wg)

	wg.Add(4) // adds 4 to the wg. after that wg will be 0

	wg.Wait() // check if the wait group is done with execution
	userProfile := &UserProfile{
		UserName:  <-userNamech,
		Tweets:    <-tweetsch,
		Followers: <-followersch,
		Following: <-followingch,
	}

	return userProfile, nil
}

func getUserName(userNamech chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	userNamech <- "FooBar"
	wg.Done() // wg -= 1
}

func getTweets(tweetsch chan []string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	tweetsch <- []string{
		"i am pain",
		"how are you doing",
		"AI will take over the world!!",
	}
	wg.Done() // wg -= 1
}

func getFollowing(followingch chan []string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	followingch <- []string{
		"Naruto",
		"Madara",
		"Sasori",
		"Daidara",
	}
	wg.Done() // wg -= 1
}

func getFollowers(followersch chan []string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 175)
	followersch <- []string{
		"gorillaKing",
		"monkeyKing",
		"testme",
		"Nagato",
	}
	wg.Done() // wg -= 1
}
