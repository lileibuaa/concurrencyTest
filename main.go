package main

import (
	"net"
	"fmt"
	"strconv"
	"time"
	"io/ioutil"
	"reflect"
)

func main() {
	chanList := make([]chan int, 10000)
	fmt.Println(len(chanList))
	for i := 0; i < len(chanList); i++ {
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
	time.Sleep(time.Millisecond * time.Duration(value/10))
	conn, err := net.Dial("tcp", "127.0.0.1:9191")
	if err != nil {
		println(err.Error(), reflect.TypeOf(err))
		return
	}
	defer conn.Close()
	fmt.Println("start write data\t" + strconv.Itoa(value))
	conn.Write([]byte("hello go socket\t" + strconv.Itoa(value)))
	inData, err := ioutil.ReadAll(conn)
	if err != nil {
		println(err.Error())
		println(reflect.TypeOf(err))
		return
	}
	println("from server\t" + string(inData))
}
