package main

import "net"
import "fmt"
import "bufio"
//import "strings" // only needed below for sample processing

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "127.0.0.1:5008")
  for { 
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
    fmt.Print("Text To analyse:", string(message))
    
    // // sample process for string received
    // newmessage := strings.ToUpper(message)
    // // send new string back to client
    // conn.Write([]byte(newmessage + "\n"))

    
  }
}