package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

type oplog struct {
	Bucket        string `json:"bucket"`
	User          string `json:"user"`
	Operation     string `json:"operation"`
	RemoteAddr    string `json:"remote_addr"`
	URI           string `json:"uri"`
	HTTPStatus    string `json:"http_status"`
	ErrorCode     string `json:"error_code"`
	BytesSent     uint32 `json:"bytes_sent"`
	BytesReceived uint32 `json:"bytes_received"`
	ObjectSize    uint32 `json:"object_size"`
	TotalTime     uint32 `json:"total_time"`
	UserAgent     string `json:"user_agent"`
	TransactionID string `json:"trans_id"`
}

func main() {

	opsSock := os.Getenv("RGW_OPS_SOCK")
	if opsSock == "" {
		panic("environment variable RGW_OPS_SOCK not set")
	}

	// wait till socket exists
	for {
		if _, err := os.Stat(opsSock); os.IsNotExist(err) {
			time.Sleep(30 * time.Second)
		} else {
			break
		}
	}

	c, err := net.Dial("unix", opsSock)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := c.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "error closing connection: %v\n", err)
		}
	}()

	// fetch json-entry from logline
	re := regexp.MustCompile(`{.*}`)
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf[:])
		if err != nil {
			panic(err)
		}

		l := re.Find(buf[0:n])
		ol := oplog{}
		err = json.Unmarshal(l, &ol)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// do not log read ops (poor mans "audit policy")
		if !strings.HasPrefix(ol.Operation, "get_") && !strings.HasPrefix(ol.Operation, "list_") && !strings.HasPrefix(ol.Operation, "stat_") {
			fmt.Println(string(l))
		}
	}
}
