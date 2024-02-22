package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	flag "github.com/spf13/pflag"
)

/*
	Реализовать простейший telnet-клиент.

	Примеры вызовов:
	go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123


	Требования:
	Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
	Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
	При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout

*/

var timeout string

func init() {
	flag.StringVarP(&timeout, "timeout", "t", "10s", "timeout for connecting to the server")
}

func scanStdIn(stdinChan chan<- string, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = fmt.Sprintf("%s\n", text)
		stdinChan <- text
	}
	cancel()
}

func scanConnIn(conn net.Conn, connChan chan<- string, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		connChan <- text
	}
	cancel()
}

func writeData(ctx context.Context, conn net.Conn, connChan <-chan string, stdinChan <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case text := <-connChan:
			fmt.Println(text)
		case text := <-stdinChan:
			conn.Write([]byte(text))
		}
	}
}

func main() {
	flag.Parse()
	timeoutDuration, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("wrong timeout value: %s", err)
	}
	addressArg := flag.Args()
	if len(addressArg) != 2 {
		log.Fatal("введите хост и порт для подключения в формате HOST PORT")
	}

	ctx, cancel := context.WithCancel(context.Background())
	connChan := make(chan string)
	stdinChan := make(chan string)

	var wg sync.WaitGroup

	address := fmt.Sprintf("%s:%s", addressArg[0], addressArg[1])

	conn, err := net.DialTimeout("tcp", address, timeoutDuration)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go scanStdIn(stdinChan, cancel)
	go scanConnIn(conn, connChan, cancel)

	wg.Add(1)
	go func() {
		defer wg.Done()
		writeData(ctx, conn, connChan, stdinChan)
	}()
	wg.Wait()
}
