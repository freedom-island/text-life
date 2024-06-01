package main

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/uiosun/text-life/handle"
	"github.com/uiosun/text-life/structs/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type Player struct {
	id uint
}

type NowDungeon struct {
	id uint // Player ID
}

type GameObject interface {
	Save() bool
}

func main() {
	h := handle.Handle{}
	err := h.GetConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("handler init is done!")

	app := fiber.New()

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/conn", websocket.New(func(c *websocket.Conn) {
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		var (
			msg []byte
			out []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			req := new(pb.MsgPacket)
			err = proto.Unmarshal(msg, req)
			if err != nil {
				fmt.Println("无法解析 Protocol 数据") // todo 改成日志吧？
				continue
			}

			switch body := req.Body.(type) {
			case *pb.MsgPacket_BodyInfoReq:
				bodyResp := &pb.BodyInfoResp{
					UserId:    body.BodyInfoReq.TargetId,
					Energy:    45,
					Knowledge: 20,
				}
				resp := &pb.MsgPacket{
					RouteId: 1,
					ReqTime: timestamppb.Now(),
					Body:    &pb.MsgPacket_BodyInfoResp{BodyInfoResp: bodyResp},
				}
				out, err = proto.Marshal(resp)
				if err != nil {
					panic(err)
				}
			case *pb.MsgPacket_EatReq:
				// todo eat something use key
				println(body.EatReq.GoodsId)
			case *pb.MsgPacket_FarmReq:
				// todo farm with some land
				println(body.FarmReq.Status)
			default:
				fmt.Println("No matching operations") // todo 改成日志吧？
				continue
			}

			if err = c.WriteMessage(websocket.BinaryMessage, out); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}
