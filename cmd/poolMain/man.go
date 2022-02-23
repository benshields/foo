package main

import "github.com/benshields/foo/pool"

func main() {
	TestCheckOutMany()
}

func TestCheckOutMany() {
	p := pool.NewPool("localhost")
	for i := 0; i < 10; i++ {
		_, _ = p.CheckOut()
	}
}
