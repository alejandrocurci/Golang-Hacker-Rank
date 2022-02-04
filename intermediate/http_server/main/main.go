package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hacker-rank/intermediate/http_server"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	actionsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	http.HandleFunc("/get", http_server.GetHandler)
	http.HandleFunc("/post", http_server.PostHandler)
	http.HandleFunc("/delete", http_server.DeleteHandler)
	go http.ListenAndServe(portSuffix, nil)
	time.Sleep(100 * time.Millisecond)

	var actions []string

	for i := 0; i < int(actionsCount); i++ {
		actionsItem := readLine(reader)
		actions = append(actions, actionsItem)
	}

	for _, actionStr := range actions {
		var action http_server.Action
		err := json.Unmarshal([]byte(actionStr), &action)
		checkError(err)
		switch action.Type {
		case "post":
			_, err := http.Post(address+"/post", "application/json", strings.NewReader(action.Payload))
			checkError(err)
		case "delete":
			client := &http.Client{}
			req, err := http.NewRequest("DELETE", address+"/delete?id="+action.Payload, nil)
			checkError(err)
			resp, err := client.Do(req)
			checkError(err)
			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}
		case "get":
			resp, err := http.Get(address + "/get?id=" + action.Payload)
			checkError(err)
			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}
			var lake http_server.Lake
			err = json.NewDecoder(resp.Body).Decode(&lake)
			checkError(err)
			fmt.Fprintf(writer, "%s\n", lake.Name)
			fmt.Fprintf(writer, "%d\n", lake.Area)
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

const portSuffix = ":3333"

var address = "http://127.0.0.1" + portSuffix

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
