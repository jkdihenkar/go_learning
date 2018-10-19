package serve_utils

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func HandleConnection(c net.Conn) {
	defer c.Close()

	// Get details of client -
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Recieved connection close - %s\n", err)
				return
			}
			fmt.Printf("Error - %s\n", err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "SHUTDOWN" {
			break
		}
		rand.Seed(time.Now().Unix())
		result := strconv.Itoa(rand.Intn(999)) + "\n"
		c.Write([]byte(string(result)))
	}
}
