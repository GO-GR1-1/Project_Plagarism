package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	port := 5050
	fmt.Printf("#DEBUG DIALING TCP Server on port %d\n", port)
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	conn, err := net.Dial("tcp", portString)
	defer conn.Close()
	if err != nil {

		fmt.Printf("#DEBUG MAIN could not connect\n")
		os.Exit(1)
	} else {

		reader := bufio.NewReader(conn)
		fmt.Printf("#DEBUG MAIN connected\n")

		readConsole := bufio.NewReader(os.Stdin)
		textToSend, _ := readConsole.ReadString('\n')

		io.WriteString(conn, fmt.Sprintf("%s\n", textToSend))
		resultString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from server")
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		fmt.Printf("#DEBUG server replied : |%s|\n", resultString)
		time.Sleep(100 * time.Millisecond)

		//}

	}
}
