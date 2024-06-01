// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go:build ignore
// +bui1ld ignore

package main

import (
	"flag"
	"github.com/pterm/pterm"
	"github.com/uiosun/text-life/client/world"
	"github.com/uiosun/text-life/structs/pb"
	"google.golang.org/protobuf/proto"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/fasthttp/websocket"
)

var addr = flag.String("addr", "localhost:3000", "http service address")

/*
PTerm: https://github.com/pterm/pterm?tab=readme-ov-file
*/
func main() {
	w := world.World{}
	w.Init()

	for {
		// input with single line
		result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("fill your select, and use 'turn' to next turn").Show()
		pterm.Println()

		switch result {
		case "get-map":
			w.Turn()
		case "exit":
			fallthrough
		case "E":
			pterm.Println("Bey~")
			return
		}

		w.RefreshUI()
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/conn"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s, type: %s", message, websocket.FormatMessageType(mt))

			var resp = &pb.BodyInfoResp{}
			err = proto.Unmarshal(message, resp)
			if err != nil {
				panic(err)
			}
			println(resp.Knowledge)
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
