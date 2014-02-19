package main

import (
	"fmt"
	"github.com/ryaninhust/beanstalk"
	"time"
)

func Example() {
	pool := beanstalk.NewPool([]string{
		"localhost:11300", "localhost:11300", "localhost:11300"})
	id, err := pool.Put([]byte("hello"), 1, 0, 120*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println("job", id)
}

func main() {
	Example()
}
