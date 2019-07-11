// 僅實作websocket方法的p2p簡易版本，完整實作請參考

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Peer struct {
	socket *websocket.Conn
	send   chan []byte
	taget  string
}

var (
	Peers []Peer
	port  string
)

func main() {
	flag.StringVar(&port, "p", "", "listen port")
	flag.Parse()
	if port != "8080" { //8080為種子
		go p2p()
	}
	log.Fatal(RunHTTP(port))
}

func p2p() {

	res, err := http.Get("http://127.0.0.1:8080/peers")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	addrs := []interface{}{}
	json.Unmarshal(b, &addrs)
	addrs = append(addrs, "127.0.0.1:8080")
	fmt.Printf("%+v\n", addrs)
	for _, v := range addrs {
		u := url.URL{Scheme: "ws", Host: v.(string), Path: "/new", RawQuery: "port=" + port}
		var dialer *websocket.Dialer

		conn, _, err := dialer.Dial(u.String(), nil)
		if err != nil {
			panic("p2p err: " + err.Error())
		}
		fmt.Println("new addr :" + conn.RemoteAddr().String())
		go func(conn *websocket.Conn) {
			conn.SetReadDeadline(time.Now().Add(100 * time.Minute))
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					fmt.Println("read err :", err)
					return
				}
				fmt.Printf("received: %s\n", string(message))
			}
		}(conn)
	}

}

func RunHTTP(port string) error {
	mux := makeMuxRouter()
	httpAddr := port
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/peers", GetPeers).Methods("GET")
	muxRouter.HandleFunc("/new", NewWS)
	return muxRouter
}

func GetPeers(res http.ResponseWriter, req *http.Request) {

	var addrs []string
	for _, p := range Peers {
		addrs = append(addrs, p.taget)
	}
	fmt.Println(port + " server GetPeers : ")
	fmt.Printf("%+v\n", addrs)
	b, e := json.Marshal(addrs)
	if e != nil {
		fmt.Println(e)
	}
	res.Write(b)
}

func NewWS(res http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	rPort, ok := q["port"]
	if !ok {
		fmt.Println("url value port is nil")
		http.NotFound(res, req)
		return
	}

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("new client error: " + err.Error())
		http.NotFound(res, req)
		return
	}
	ip := strings.Split(conn.RemoteAddr().String(), ":")
	fmt.Println("new Peer :" + ip[0] + ":" + rPort[0])
	p := Peer{conn, make(chan []byte), ip[0] + ":" + rPort[0]}
	Peers = append(Peers, p)

	// 新節點監聽
	go p.Write()
	go p.Read()
}

func (c *Peer) Read() {
	defer func() {
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			c.socket.Close()
			break
		}
		fmt.Println(string(message))
		// Manager.broadcast <- jsonMessage
	}
}

func (c *Peer) Write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
