package main

import (
	"fmt"
	"net"
	"sync" // the waitgroup functions is used on go rutine waits for runs, calls, and done
)

func main() {
	var wg sync.WaitGroup // declaring a variable named wg to hold the function call sync.WaitGroup
	fmt.Printf("Scanning ...\n")
	for i := 1; i <= 80; i++ { // creating a for loop to use on the go func created bellow
		wg.Add(1)        //calls add function to add 1 go-rutine (the go func) to the waitGroup
		go func(j int) { //creating a go-rutine that declares a varible named j that is an integer
			address := fmt.Sprintf("10.0.0.203:%d", j) //creating an varible named address, then uses fmt.Sprintf to create a formated function
			conn, err := net.Dial("tcp", address)      //creates two varibles named conn, and error that hold the function call net.Dial() which takes to string arguments "tcp" (the protocal) and the formated string named address
			if err != nil {                            //if statment that says if err is does not equal 0 then retrun to the begining of the function
				return
			}
			conn.Close()                   //if conn is equal to 0 then it will use the Close function to close the tcp connection
			fmt.Printf("%d/tcp open\n", j) //print the open port
		}(i) //is used to to end the forloop count
	}
	wg.Wait() //calls the function Wait() that waits untill the counter is 0
}
