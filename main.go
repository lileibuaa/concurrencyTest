package main

import (
	"net"
	"fmt"
	"strconv"
	"time"
)

const MAX_CONNECT = 200
const DELAY = false

func main() {
	chanList := make([]chan int, MAX_CONNECT)
	for i := 0; i < MAX_CONNECT; i++ {
		chanList[i] = make(chan int, 1)
		go connectSocket(chanList[i])
		chanList[i] <- i
	}
	tmpChan := make(chan int)
	<-tmpChan
}

func connectSocket(c chan int) {
	defer close(c)
	value := <-c
	if DELAY {
		time.Sleep(time.Millisecond * time.Duration(value))
	}
	conn, err := net.Dial("tcp", "127.0.0.1:10191")
	if err != nil {
		println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("start write data\t" + strconv.Itoa(value))
	conn.Write([]byte("hello go socket\t" + strconv.Itoa(value)))
	//var readData []byte = make([]byte, 1024)
	var readData []byte = make([]byte, 1024)
	for {
		readL, err := conn.Read(readData)
		if err != nil {
			println(err.Error())
			break
		}
		if readL == -1 {
			break
		}
		println("from server\t" + string(readData[:readL]) + "\t" + conn.LocalAddr().String())
	}
	//println("from server\t" + string(inData) + "\t" + conn.LocalAddr().String())
}
