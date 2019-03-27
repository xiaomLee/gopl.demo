package main

import (
	"time"
	"errors"
	"bufio"
	"os"
	"fmt"
	"strings"
)

func Send(c chan string, msg string) (error) {
	select {
	case c<- msg:
		return nil
	case <- time.After(2 * time.Second):
		return errors.New("请稍后再试！")
	}
}

func Receive(c chan string)  {
	for {
		select {
		case msg := <-c:
			println("received " + msg)
		}
	}
}

func ReceiveOne(c chan string)  {
	select {
	case msg := <- c:
		println(msg)
	case <- time.After(1 * time.Second):

	}

}

func main() {
	c := make(chan string, 2)
	//go Receive(c)
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("commond-> ")
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line," ")
		switch tokens[0] {
		case "close":
			close(c)
		case "pop":
			ReceiveOne(c)
		case "consume":
			go Receive(c)
		case "send":
			err := Send(c, strings.Join(tokens[1:], " "))
			if err != nil {
				println(err.Error())
			}
		default:
			println("commond error")
		}
		time.Sleep(10 * time.Millisecond)
	}
}

