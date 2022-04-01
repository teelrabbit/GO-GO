package main

import (
    "fmt"
    "net"
)

 func main() {
     //ip := "10.10.177.237" //defineing IP address
     var ip string //making Var name IP that is set to a string type
     fmt.Printf("Enter IP-ADDRESS:: ")
     fmt.Scanln(&ip) //takes user input for the varible ip
     fmt.Printf("Scanning ...\n")
     for i := 1; i < 1024; i++ { //for-loop for iterat8ing though 1024 ports
        address := fmt.Sprintf("%s:%d",ip, i) //the var address is equal to a formated string that is the IP_ADDRESS:PORT_NUM
        conn, err:= net.Dial("tcp", address) //conn and err are equal to the return value from net.Dial
        if err != nil { 
            // if nil, the port is either closed of filtered
            continue
        }
        conn.Close() //if there is a succesful connection it closes the connection
        fmt.Printf("%d/tcp open\n", i) //prints the open port number
    }
 }
