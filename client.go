package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"bufio"
	"os"
    "strings"
    "time"
)


func main() {
	port := 5000
	fmt.Printf("#DEBUG DIALING TCP Server on port %d\n", port)
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	conn, err := net.Dial("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not connect\n", err)
		os.Exit(1)
	} else {

        defer conn.Close()
        reader := bufio.NewReader(conn)
		fmt.Printf("#DEBUG MAIN connected\n")
        for i:= 0; i < 10; i++{

            io.WriteString(conn, fmt.Sprintf("Coucou %d\n", i))
            
            resultString, err := reader.ReadString('\n')
            if (err != nil){
                fmt.Printf("DEBUG MAIN could not read from server")
                os.Exit(1)
            }
            resultString = strings.TrimSuffix(resultString, "\n")
            fmt.Printf("#DEBUG server replied : |%s|\n", resultString)
            time.Sleep(1000 * time.Millisecond)
            
    }

	}

}