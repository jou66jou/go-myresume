package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/zserge/webview"
)

const (
	batName = "\\fileserver.exe"
)

var (
	version string
	w, h    int
)

func main() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	// handle err...
	if err != nil {
		log.Printf("net.Dial get err : %v", err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")
	log.Println(ip[0])

	file, err := os.Open("webview/config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		tmp := strings.Split(scanner.Text(), "=")
		switch tmp[0] {
		case "hight":
			h, err = strconv.Atoi(tmp[1])
			if err != nil || h == 0 {
				h = 1024
			}
		case "wight":
			w, err = strconv.Atoi(tmp[1])
			if err != nil || w == 0 {
				w = 768
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// just start a webview
	stratWebView(ip[0])

	// // if you want close webview by process
	// var done chan bool
	// done = make(chan bool, 1)
	// browser.Init("https://www.google.com.tw/")
	// go browser.StartBrowser(done)
	// // close browser after 7 sec
	// log.Println("start wait 7 sec to send sign to done chan")
	// time.Sleep(7 * time.Second)
	// done <- true
	// log.Println("done is close, wait stopChan...")
	// // stop main
	// stopChan := make(chan os.Signal, 1)
	// signal.Notify(stopChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)
	// <-stopChan
}

func stratWebView(address string) {

	webView := webview.New(webview.Settings{
		URL:                    "http://" + address + ":1323/",
		Width:                  h,
		Height:                 w,
		Title:                  "機車排氣檢驗系統",
		ExternalInvokeCallback: nil,
		Debug:                  true,
	})

	webView.Run()
}
