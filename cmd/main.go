package main

import (
	"github.com/bandishjilka/virmon/lib/handler"
)

func main() {
	// monitors := monitor.GetMonitors()
	// for i, m := range monitors {
	// 	monitorMsg := fmt.Sprintf("Monitor Index: %d, Width: %d, Height: %d", m.Index, m.Width, m.Height)
	// 	log.Println(monitorMsg)

	// 	monitor.SaveScreenshot(i)
	// }

	handler.StartWebSocketServer()
}
