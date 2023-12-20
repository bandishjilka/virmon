package handler

import (
	"bufio"
	"bytes"
	"image/png"
	"log"
	"net/http"

	// "golang.org/x/net/websocket"
	"github.com/bandishjilka/virmon/lib/monitor"
	"github.com/gorilla/websocket"
)

const PORT = ":8080"

var upgrader = websocket.Upgrader{}

func StartWebSocketServer() {
	http.HandleFunc("/", webSocketHandler)
	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// s := websocket.Server{
	// 	Handler: webSocketHanlder,
	// 	Config: websocket.Config{
	// 		Header: make(http.Header),
	// 	},
	// }
	// s.ServeHTTP(w, r)
	// })
	log.Printf("web socket listening on port: %s", PORT)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Panicf("unable to start websocket server, error: %s", err.Error())
	}

}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error while upgrading, error: %s", err.Error())
		return
	}
	defer c.Close()

	monitors := monitor.GetMonitors()
	for i := range monitors {
		img, err := monitor.CaptureMonitor(i)
		if err != nil {
			log.Printf("error occurred in capturing display, error: %s", err.Error())
			return
		}

		var b bytes.Buffer
		empty := bufio.NewWriter(&b)
		err = png.Encode(empty, img)
		empty.Flush()
		if err != nil {
			log.Printf("error while encoding into png, error: %s", err.Error())
			return
		}

		c.WriteMessage(websocket.BinaryMessage, b.Bytes())
	}
}

// func webSocketHanlder(ws *websocket.Conn) {
// 	// websocket.Message.Send(ws, "hello from server")
// 	monitors := monitor.GetMonitors()
// 	for i, m := range monitors {
// 		monitorMsg := fmt.Sprintf("Monitor Index: %d, Width: %d, Height: %d", m.Index, m.Width, m.Height)
// 		log.Println(monitorMsg)

// 		img, err := monitor.CaptureMonitor(i)
// 		if err != nil {
// 			log.Printf("error occurred in capturing display, error: %s", err.Error())
// 			return
// 		}

// 		// b := new(bytes.Buffer)
// 		var b bytes.Buffer
// 		empty := bufio.NewWriter(&b)
// 		err = png.Encode(empty, img)
// 		empty.Flush()
// 		if err != nil {
// 			log.Printf("error while encoding into png, error: %s", err.Error())
// 			return
// 		}

// 		// err = websocket.Message.Send(ws, b)
// 		// if err != nil {
// 		// 	log.Printf("error while sending image to websocket, error: %s", err.Error())
// 		// 	return
// 		// }

// 		_, err = ws.Write(b.Bytes())
// 		if err != nil {
// 			log.Printf("error while sending image to websocket, error: %s", err.Error())
// 			return
// 		}
// 	}
// }
