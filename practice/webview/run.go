package main

import (
	"log"
	"net"
	"strings"

	"github.com/zserge/webview"
)

const (
	batName = "\\fileserver.exe"
)

var (
	version string
)

func main() {
	// address, err := command.GetIP()
	// if err != nil {
	// 	log.Println("get ip fail : ", err)
	// }
	// if len(address) == 0 {
	// 	log.Println("not found local ip address")
	// }
	// for i := range address {
	// 	log.Printf("address %v : %s", i, address[i])
	// }
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
		Width:                  500,
		Height:                 800,
		Resizable:              true,
		ExternalInvokeCallback: nil,
	})
	webView.Run()
}
