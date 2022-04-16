package main

import (
	"math"

	"github.com/fogleman/gg"
)

func main() {
	// dc := gg.NewContext(720, 720)
	// dc.DrawCircle(500, 500, 100)
	// dc.SetRGB(0, 0, 0)
	// dc.Fill()
	// dc.SavePNG("raiden_logo.png")
	var w, h int = 720, 720
	dc := gg.NewContext(w, h)
	dc.SetRGBA(1, 1, 0, 1)
	dc.SetLineWidth(1)

	var (
		topx, topy       float64 = 360, 0
		bottomx, bottomy float64 = 360, 720
		leftx, lefty     float64 = 360 - 180*math.Sqrt(2), 360
		rightx, righty   float64 = 360 + 180*math.Sqrt(2), 360

		innerleftx, innerlefty   float64 = 240, 360
		innerrightx, innerrighty float64 = 480, 360
		// topleftx, toplefty         float64 = 360 - 60*math.Sqrt(2), 360 - 60*math.Sqrt(2)
		// bottomrightx, bottomrighty float64 = 360 + 60*math.Sqrt(2), 360 + 60*math.Sqrt(2)
		topleftx, toplefty         float64 = 360, 240
		bottomrightx, bottomrighty float64 = 360, 480
	)

	// 外框
	dc.DrawLine(topx, topy, leftx, lefty)
	dc.DrawLine(topx, topy, rightx, righty)
	dc.DrawLine(bottomx, bottomy, leftx, lefty)
	dc.DrawLine(bottomx, bottomy, rightx, righty)
	// 内线
	dc.DrawLine(topx, topy, topleftx, toplefty)
	dc.DrawLine(innerrightx, innerrighty, topleftx, toplefty)
	dc.DrawLine(bottomx, bottomy, bottomrightx, bottomrighty)
	dc.DrawLine(innerleftx, innerlefty, bottomrightx, bottomrighty)
	dc.DrawLine(innerleftx, innerlefty, topx, topy)
	dc.DrawLine(innerrightx, innerrighty, bottomx, bottomy)
	dc.DrawCircle(360, 360, 60)
	dc.Stroke()
	dc.SavePNG("raiden_logo.png")
}
