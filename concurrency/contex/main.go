package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// a prog or any third-part call might take more time to execute.
// to crub that we have to use contex
// we can set context.WithValue(context.Background(), "username", "daftPunk")
// we can fetch the uname := ctx.Value("username")
func main() {

	start := time.Now()
	resp, err := getUserResponse()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response : %v \ntime taken: %v\n", resp, time.Since(start))
}

func getUserResponse() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel() // once time limit exceeds use it to cancel execution

	type result struct {
		userResp string
		err      error
	}
	resultch := make(chan result, 1)

	go func() {
		resp, err := thirdpartHTTPCall()
		resultch <- result{
			userResp: resp,
			err:      err,
		}
	}()

	select {
	// <- the conext timout is exceeded
	// <- the context has manually canceled -> Cancel()
	case <-ctx.Done(): // check if the channel execution is done
		return "", ctx.Err() // return error returned
	case resp := <-resultch:
		return resp.userResp, nil
	}
}

func thirdpartHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 200)
	return "hello from the user 1", nil
}
