package main
//
// This program should produce no output because it runs to fast for the packets to be properly recived
//
import (
    "fmt"
    "net"
)

func main() {
    for i:=1;i<=1024;i++ { //for loop for scanning a range of 1024
        go func(j int) { //greating a go func that initalizes a varibale name "j"
            address := fmt.Sprintf("scanme.nmap.org:%d", j) //creates a formated string to hold the IP-ADDRESS:PORT-NUM
            conn, err := net.Dial("tcp", address) //varibles con and err are equal to the return value of net.Dial
            if err != nil {
                return // returns to the begining of the loop if err is not equal to nil (nil=A VALUE OF 0) 
            }
            conn.Close() //if there is a succesful connection it will close the connection
            fmt.Printf("%d open\n", j) //prints the port number of the open port using the var "j"
        }(i)
    }
}
