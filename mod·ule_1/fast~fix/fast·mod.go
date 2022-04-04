package main

import (
	"fmt"
	"net"
	"sort"
	//"sync" // the waitgroup functions is used on go rutine waits for runs, calls, and done
)

func workers(ports, results chan int, ip string) { //chan creates a channel that is a data type of an int named ports, and wg points to the function WorkGroup

	for p := range ports { // used to use items from channels
		address := fmt.Sprintf("%s:%d", ip, p) //creates an var named address, then uses fmt to create a formated string that holds the varible ip and a loop counter
		conn, err := net.Dial("tcp", address) // creates 2 varibles that hold return value of the net.Dial function which has 2 arguments (the protocal and the address/port)
		if err != nil { //if the var err does not equal to 0
			results <- 0 // asssigns 0 through a channel to results
			continue // the loop continues
		}
		conn.Close() //if 0 then conn will close the tcp connection
		results <- p
	}
}

func main() {
	var ip string //var made  named ip that holds a string value
	fmt.Print("Enter IP·Address:: ") //uses print func to ask the user for an IP
	fmt.Scanf("%s", &ip) //scans for the value for ip

	ports := make(chan int, 100) // make is used for channels 
	results := make(chan int)
	var openports []int // initializer

	for i := 0; i < cap(ports); i++ { //cap is used to give a capasity for the link
		go workers(ports, results, ip) //calls the go func "workers" 
	}

	go func() { 
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports { // prints if the port is opened 
		fmt.Printf("%d open\n", port)
	}

}
