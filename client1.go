package main

import "net"
import "fmt"
import "bufio"
import "os"
import "io/ioutil"
import "strings"

func main() {

  fmt.Println("Waiting for server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":5008")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
    //read file name in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Name of the file to analyse: ")
    text, err := reader.ReadString('\n')

    //Delete the \n character
    text = strings.TrimSuffix(text, "\n")
    
    if err != nil {
      fmt.Println("can't find file")
    }

    //read content of file
    contBytes, _ := ioutil.ReadFile(text)

    //convert content from bytes to string
    contText := string(contBytes)

    // send to socket
    fmt.Fprintf(conn, contText)
    
    // listen for reply
    // message, _ := bufio.NewReader(conn).ReadString('\n')
    // fmt.Print("Message from server: "+message)
  }
}