package monitor

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/kbinani/screenshot"
)

type MonitorInfo struct {
	Width  int
	Height int
}

// function to get list of monitors
func GetMonitors() (monitors []MonitorInfo) {
	monitorCount := screenshot.NumActiveDisplays()
	for i := 0; i < monitorCount; i++ {
		rect := screenshot.GetDisplayBounds(i)
		monitors = append(monitors, MonitorInfo{
			Width:  rect.Max.X,
			Height: rect.Max.Y,
		})
	}

	return
}

func CaptureMonitor(monitorIndex int) (image *image.RGBA, err error) {
	return screenshot.CaptureDisplay(monitorIndex)
}

func SaveScreenshot(monitorIndex int) {
	img, err := CaptureMonitor(monitorIndex)
	if err != nil {
		msg := fmt.Sprintf("error occurred in capturing display, error: %s", err.Error())
		log.Println(msg)
		return
	}

	f, err := os.Create("image.png")
	if err != nil {
		msg := fmt.Sprintf("error while creating file, error: %s", err.Error())
		log.Println(msg)
		return
	}

	err = png.Encode(f, img)
	if err != nil {
		msg := fmt.Sprintf("error while encoding into png, error: %s", err.Error())
		log.Println(msg)
		return
	}
}
