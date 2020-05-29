package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ( // MARK: var

	// timers
	framecount int
	// player
	px, py    int
	pmoveon   bool
	direction string
	pblock    = 15
	pblocknew = pblock
	// images
	//trees
	tree1 = rl.NewRectangle(9, 598, 56, 202)
	//terrain
	ter1 = rl.NewRectangle(0, 1, 256, 148)
	ter2 = rl.NewRectangle(0, 150, 256, 139)
	ter3 = rl.NewRectangle(0, 291, 256, 159)
	imgs rl.Texture2D
	// map
	drawblocknext, drawblock, blockactive, blockx, blocky, nextx, nexty int
	mapa                                                                = 280000
	blocks                                                              = make([]bool, mapa)
	levelmap                                                            = make([]string, mapa)
	lineswitch                                                          bool
	// input
	mousepos rl.Vector2
	// debug
	debugon bool
	// camera screen
	borderson, gridon, fullscreenon bool
	screenW                         = int32(1600)
	screenH                         = int32(900)
	camera                          rl.Camera2D
)

// MARK: notes
/*



	7 X 8 OUTER = 56
	7 X 8 INNER = 56
	TOTAL 1 SCREEN = 112 BLOCKS

	(7 X 50) X (8 X 50) = 350 X 400 = 140 000
	(7 X 50) X (8 X 50) = 350 X 400 = 140 000
	TOTAL 280 000 BLOCKS



*/
func getactiveblock() { // MARK: getactiveblock()
	ychange := float32(0)
	if mousepos.X > 64 && mousepos.X < 192 {
		blockactive = drawblocknext + 50
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 192 && mousepos.X < 320 {
		blockactive = drawblocknext + 101
		for a := 0; a < 7; a++ {
			if mousepos.Y > 32+ychange && mousepos.Y < 192+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 320 && mousepos.X < 448 {
		blockactive = drawblocknext + 51
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 448 && mousepos.X < 576 {
		blockactive = drawblocknext + 102
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 576 && mousepos.X < 704 {
		blockactive = drawblocknext + 52
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 704 && mousepos.X < 832 {
		blockactive = drawblocknext + 103
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 832 && mousepos.X < 960 {
		blockactive = drawblocknext + 53
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 960 && mousepos.X < 1088 {
		blockactive = drawblocknext + 104
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 1088 && mousepos.X < 1216 {
		blockactive = drawblocknext + 54
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 1216 && mousepos.X < 1344 {
		blockactive = drawblocknext + 105
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 1344 && mousepos.X < 1472 {
		blockactive = drawblocknext + 55
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	} else if mousepos.X > 1472 && mousepos.X < 1600 {
		blockactive = drawblocknext + 106
		for a := 0; a < 7; a++ {
			if mousepos.Y > 0+ychange && mousepos.Y < 128+ychange {
				blockactive += (a * 100)
			}
			ychange += 128
		}
	}

}

func input() { // MARK: keys input

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		pblocknew = blockactive
		if pblocknew != pblock {
			getblocknumbers()
			choosedirection()
		}
	}
	if rl.IsKeyPressed(rl.KeyRightControl) {
		camera.Zoom = 1.0
		camera.Target.X = 128.0
		camera.Target.Y = 64.0
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		camera.Target.X -= 16
	}
	if rl.IsKeyDown(rl.KeyRight) {
		camera.Target.X += 16
	}
	if rl.IsKeyDown(rl.KeyUp) {
		camera.Target.Y -= 16
	}
	if rl.IsKeyDown(rl.KeyDown) {
		camera.Target.Y += 16
	}
	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDivide) {
		if borderson {
			borderson = false
		} else {
			borderson = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom < 2.1 {
			camera.Zoom += 0.1
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom > 0.5 {
			camera.Zoom -= 0.1
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}
}
func grid() {
	xchange := int32(0)
	for b := 0; b < 7; b++ {
		ychange := int32(0)
		for a := 0; a < 7; a++ {
			rl.DrawLine(0+xchange, 64+ychange, 128+xchange, 0+ychange, rl.Magenta)
			rl.DrawLine(0+xchange, 64+ychange, 128+xchange, 128+ychange, rl.Magenta)
			rl.DrawLine(128+xchange, 0+ychange, 256+xchange, 64+ychange, rl.Magenta)
			rl.DrawLine(128+xchange, 128+ychange, 256+xchange, 64+ychange, rl.Magenta)
			ychange += 128
		}
		xchange += 256
	}
}
func choosedirection() { // MARK: choosedirection()
	if pblocknew > pblock {
		if pblocknew-pblock < 8 {
			direction = "r"
			pmoveon = true
		} else if (pblocknew-pblock)%100 == 0 {
			direction = "d"
			pmoveon = true
		} else if px > blockx {
			direction = "dl"
			pmoveon = true
		} else if px < blockx {
			direction = "dr"
			pmoveon = true
		}

	} else {
		if pblock-pblocknew < 8 && pblock-pblocknew > 0 {
			direction = "l"
			pmoveon = true
		} else if (pblock-pblocknew)%100 == 0 {
			direction = "u"
			pmoveon = true
		} else if px > blockx {
			direction = "ul"
			pmoveon = true
		} else if px < blockx {
			direction = "ur"
			pmoveon = true
		}
	}
}
func checkpblockud(dir string) { // MARK: checkpblockud()
	switch dir {
	case "dr":
		// check down left
		check2 := false
		for a := 1; a < 10; a++ {
			nextx = px - 128*a
			nexty = py + 64*a
			check := nextblock()
			if check == pblocknew {
				direction = "dl"
				check2 = true
				moveplayer()
				break
			}
		}
		// check down right
		if !check2 {
			for a := 1; a < 10; a++ {
				nextx = px + 128*a
				nexty = py - 64*a
				check := nextblock()
				if check == pblocknew {
					direction = "ur"
					moveplayer()
					break
				}
			}
		}
	case "dl":
		// check up left
		check2 := false
		for a := 1; a < 10; a++ {
			nextx = px - 128*a
			nexty = py - 64*a
			check := nextblock()
			if check == pblocknew {
				direction = "ul"
				check2 = true
				moveplayer()
				break
			}
		}
		// check down right
		if !check2 {
			for a := 1; a < 10; a++ {
				nextx = px + 128*a
				nexty = py + 64*a
				check := nextblock()
				if check == pblocknew {
					direction = "dr"
					moveplayer()
					break
				}
			}
		}
	case "ul":
		// check down left
		check2 := false
		for a := 1; a < 10; a++ {
			nextx = px - 128*a
			nexty = py + 64*a
			check := nextblock()
			if check == pblocknew {
				direction = "dl"
				check2 = true
				moveplayer()
				break
			}
		}
		// check up right
		if !check2 {
			for a := 1; a < 10; a++ {
				nextx = px + 128*a
				nexty = py - 64*a
				check := nextblock()
				if check == pblocknew {
					direction = "ur"
					moveplayer()
					break
				}
			}
		}
	case "ur":
		// check down right
		check2 := false
		for a := 1; a < 10; a++ {
			nextx = px + 128*a
			nexty = py + 64*a
			check := nextblock()
			if check == pblocknew {
				direction = "dr"
				check2 = true
				moveplayer()
				break
			}
		}
		// check up left
		if !check2 {
			for a := 1; a < 10; a++ {
				nextx = px - 128*a
				nexty = py - 64*a
				check := nextblock()
				if check == pblocknew {
					direction = "ul"
					moveplayer()
					break
				}
			}
		}
	}
}
func moveplayer() { // MARK: moveplayer()

	switch direction {
	case "dr":
		if pblock != pblocknew {
			if framecount%15 == 0 {
				nextx = px + 128
				nexty = py + 64
				pblock = nextblock()
				checkpblockud("dr")
			}
		} else {
			direction = ""
			pmoveon = false
		}
	case "dl":
		if pblock != pblocknew {
			if framecount%15 == 0 {
				nextx = px - 128
				nexty = py + 64
				pblock = nextblock()
				checkpblockud("dl")
			}
		} else {
			direction = ""
			pmoveon = false
		}
	case "ur":
		if pblock != pblocknew {
			if framecount%15 == 0 {
				nextx = px + 128
				nexty = py - 64
				pblock = nextblock()
				checkpblockud("ur")
			}
		} else {
			direction = ""
			pmoveon = false
		}
	case "ul":
		if pblock != pblocknew {
			if framecount%15 == 0 {
				nextx = px - 128
				nexty = py - 64
				pblock = nextblock()
				checkpblockud("ul")
			}
		} else {
			direction = ""
			pmoveon = false
		}
	case "d":
		if pblock != pblocknew {
			if framecount%6 == 0 {
				pblock += 100
			}
		} else {
			direction = ""
			pmoveon = false
		}
	case "u":
		if pblock != pblocknew {
			if framecount%6 == 0 {
				pblock -= 100
			}
		} else {
			direction = ""
			pmoveon = false
		}
	case "r":
		if pblock != pblocknew {
			if framecount%6 == 0 {
				pblock++
			}
		} else {
			direction = ""
			pmoveon = false
		}
	case "l":
		if pblock != pblocknew {
			if framecount%6 == 0 {
				pblock--
			}
		} else {
			direction = ""
			pmoveon = false
		}
	}
}
func getblocknumbers() { // getblocknumbers()
	drawx := 0
	drawy := 0
	drawblock = drawblocknext
	linecount := 0
	for a := 0; a < 112; a++ {
		//  get pblocknew xy
		if drawblock == pblocknew {
			blockx = drawx
			blocky = drawy
		}
		// draw player
		if drawblock == pblock {
			px = drawx
			py = drawy
		}
		linecount++
		drawblock++
		drawx += 256
		if lineswitch {
			if linecount == 7 {
				drawx = 0
				drawy += 64
				linecount = 0
				drawblock += 43
				lineswitch = false
			}
		} else {
			if linecount == 7 {
				drawx = 128
				drawy += 64
				linecount = 0
				drawblock += 43
				lineswitch = true
			}
		}
	}
}
func nextblock() int { // MARK: nextblock

	block := drawblocknext + 50

	if nextx == 128 {
		if nexty == 64 {
			block = drawblocknext + 50
		} else if nexty == 192 {
			block = drawblocknext + 150
		} else if nexty == 320 {
			block = drawblocknext + 250
		} else if nexty == 448 {
			block = drawblocknext + 350
		} else if nexty == 576 {
			block = drawblocknext + 450
		} else if nexty == 704 {
			block = drawblocknext + 550
		} else if nexty == 832 {
			block = drawblocknext + 650
		}
	} else if nextx == 256 {
		if nexty == 128 {
			block = drawblocknext + 101
		} else if nexty == 256 {
			block = drawblocknext + 201
		} else if nexty == 384 {
			block = drawblocknext + 301
		} else if nexty == 512 {
			block = drawblocknext + 401
		} else if nexty == 640 {
			block = drawblocknext + 501
		} else if nexty == 768 {
			block = drawblocknext + 601
		} else if nexty == 896 {
			block = drawblocknext + 701
		}
	} else if nextx == 384 {
		if nexty == 64 {
			block = drawblocknext + 51
		} else if nexty == 192 {
			block = drawblocknext + 151
		} else if nexty == 320 {
			block = drawblocknext + 251
		} else if nexty == 448 {
			block = drawblocknext + 351
		} else if nexty == 576 {
			block = drawblocknext + 451
		} else if nexty == 704 {
			block = drawblocknext + 551
		} else if nexty == 832 {
			block = drawblocknext + 651
		}
	} else if nextx == 512 {
		if nexty == 128 {
			block = drawblocknext + 102
		} else if nexty == 256 {
			block = drawblocknext + 202
		} else if nexty == 384 {
			block = drawblocknext + 302
		} else if nexty == 512 {
			block = drawblocknext + 402
		} else if nexty == 640 {
			block = drawblocknext + 502
		} else if nexty == 768 {
			block = drawblocknext + 602
		} else if nexty == 896 {
			block = drawblocknext + 702
		}
	} else if nextx == 640 {
		if nexty == 64 {
			block = drawblocknext + 52
		} else if nexty == 192 {
			block = drawblocknext + 152
		} else if nexty == 320 {
			block = drawblocknext + 252
		} else if nexty == 448 {
			block = drawblocknext + 352
		} else if nexty == 576 {
			block = drawblocknext + 452
		} else if nexty == 704 {
			block = drawblocknext + 552
		} else if nexty == 832 {
			block = drawblocknext + 652
		}
	} else if nextx == 768 {
		if nexty == 128 {
			block = drawblocknext + 103
		} else if nexty == 256 {
			block = drawblocknext + 203
		} else if nexty == 384 {
			block = drawblocknext + 303
		} else if nexty == 512 {
			block = drawblocknext + 403
		} else if nexty == 640 {
			block = drawblocknext + 503
		} else if nexty == 768 {
			block = drawblocknext + 603
		} else if nexty == 896 {
			block = drawblocknext + 703
		}
	} else if nextx == 896 {
		if nexty == 64 {
			block = drawblocknext + 53
		} else if nexty == 192 {
			block = drawblocknext + 153
		} else if nexty == 320 {
			block = drawblocknext + 253
		} else if nexty == 448 {
			block = drawblocknext + 353
		} else if nexty == 576 {
			block = drawblocknext + 453
		} else if nexty == 704 {
			block = drawblocknext + 553
		} else if nexty == 832 {
			block = drawblocknext + 653
		}
	} else if nextx == 1024 {
		if nexty == 128 {
			block = drawblocknext + 104
		} else if nexty == 256 {
			block = drawblocknext + 204
		} else if nexty == 384 {
			block = drawblocknext + 304
		} else if nexty == 512 {
			block = drawblocknext + 404
		} else if nexty == 640 {
			block = drawblocknext + 504
		} else if nexty == 768 {
			block = drawblocknext + 604
		} else if nexty == 896 {
			block = drawblocknext + 704
		}
	} else if nextx == 1152 {
		if nexty == 64 {
			block = drawblocknext + 54
		} else if nexty == 192 {
			block = drawblocknext + 154
		} else if nexty == 320 {
			block = drawblocknext + 254
		} else if nexty == 448 {
			block = drawblocknext + 354
		} else if nexty == 576 {
			block = drawblocknext + 454
		} else if nexty == 704 {
			block = drawblocknext + 554
		} else if nexty == 832 {
			block = drawblocknext + 654
		}
	} else if nextx == 1280 {
		if nexty == 128 {
			block = drawblocknext + 105
		} else if nexty == 256 {
			block = drawblocknext + 205
		} else if nexty == 384 {
			block = drawblocknext + 305
		} else if nexty == 512 {
			block = drawblocknext + 405
		} else if nexty == 640 {
			block = drawblocknext + 505
		} else if nexty == 768 {
			block = drawblocknext + 605
		} else if nexty == 896 {
			block = drawblocknext + 705
		}
	} else if nextx == 1408 {
		if nexty == 64 {
			block = drawblocknext + 55
		} else if nexty == 192 {
			block = drawblocknext + 155
		} else if nexty == 320 {
			block = drawblocknext + 255
		} else if nexty == 448 {
			block = drawblocknext + 355
		} else if nexty == 576 {
			block = drawblocknext + 455
		} else if nexty == 704 {
			block = drawblocknext + 555
		} else if nexty == 832 {
			block = drawblocknext + 655
		}
	} else if nextx == 1536 {
		if nexty == 128 {
			block = drawblocknext + 106
		} else if nexty == 256 {
			block = drawblocknext + 206
		} else if nexty == 384 {
			block = drawblocknext + 306
		} else if nexty == 512 {
			block = drawblocknext + 406
		} else if nexty == 640 {
			block = drawblocknext + 506
		} else if nexty == 768 {
			block = drawblocknext + 606
		} else if nexty == 896 {
			block = drawblocknext + 706
		}
	}

	return block

}
func createmap() { // MARK: createmap()
	for a := 0; a < mapa; a++ {
		levelmap[a] = "."
	}
	for a := 0; a < mapa; a++ {
		if rolldice()+rolldice() == 12 {
			levelmap[a] = "#"
		}
	}
	for a := 0; a < mapa; a++ {
		if rolldice()+rolldice() == 12 {
			levelmap[a] = "$"
		}
	}
}
func drawconsole() { // MARK:drawconsole()
	count := 0
	for a := 0; a < 98; a++ {
		b := levelmap[a]
		print(b)
		count++
		if lineswitch {
			if count == 6 {
				count = 0
				println()
				lineswitch = false
			}
		} else {
			if count == 7 {
				count = 0
				println()
				lineswitch = true
			}
		}
	}
}
func updrawblock() { // MARK:updrawblock()
	if framecount%60 == 0 {
		drawblocknext = pblock - 353
	}
}
func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window
	start()
	createmap()
	//	drawconsole()
	raylib()
}
func animate() { // MARK: animate()
	if framecount%12 == 0 {
		tree1.X += 64
		if tree1.X > 570 {
			tree1.X = 9
		}
	}
}
func debug() { // MARK: debug
	rl.DrawRectangle(screenW-300, 0, 500, screenW, rl.Fade(rl.Black, 0.9))

	mouseposXTEXT := fmt.Sprintf("%.0f", mousepos.X)
	mouseposYTEXT := fmt.Sprintf("%.0f", mousepos.Y)
	zoomTEXT := fmt.Sprintf("%.1f", camera.Zoom)
	lineswitchTEXT := strconv.FormatBool(lineswitch)
	blockactiveTEXT := strconv.Itoa(blockactive)
	pblocknewTEXT := strconv.Itoa(pblocknew)
	pblockTEXT := strconv.Itoa(pblock)
	pxTEXT := strconv.Itoa(px)
	pyTEXT := strconv.Itoa(py)
	blockxTEXT := strconv.Itoa(blockx)
	blockyTEXT := strconv.Itoa(blocky)
	drawblocknextTEXT := strconv.Itoa(drawblocknext)
	pmoveonTEXT := strconv.FormatBool(pmoveon)

	rl.DrawText(mouseposXTEXT, screenW-290, 10, 10, rl.White)
	rl.DrawText("mouseposX", screenW-200, 10, 10, rl.White)
	rl.DrawText(mouseposYTEXT, screenW-290, 20, 10, rl.White)
	rl.DrawText("mouseposY", screenW-200, 20, 10, rl.White)
	rl.DrawText(zoomTEXT, screenW-290, 30, 10, rl.White)
	rl.DrawText("zoom", screenW-200, 30, 10, rl.White)
	rl.DrawText(lineswitchTEXT, screenW-290, 40, 10, rl.White)
	rl.DrawText("lineswitch", screenW-200, 40, 10, rl.White)
	rl.DrawText(blockactiveTEXT, screenW-290, 50, 10, rl.White)
	rl.DrawText("blockactive", screenW-200, 50, 10, rl.White)
	rl.DrawText(pblocknewTEXT, screenW-290, 60, 10, rl.White)
	rl.DrawText("pblocknew", screenW-200, 60, 10, rl.White)
	rl.DrawText(pblockTEXT, screenW-290, 70, 10, rl.White)
	rl.DrawText("pblock", screenW-200, 70, 10, rl.White)
	rl.DrawText(pxTEXT, screenW-290, 80, 10, rl.White)
	rl.DrawText("px", screenW-200, 80, 10, rl.White)
	rl.DrawText(pyTEXT, screenW-290, 90, 10, rl.White)
	rl.DrawText("py", screenW-200, 90, 10, rl.White)
	rl.DrawText(blockxTEXT, screenW-290, 100, 10, rl.White)
	rl.DrawText("blockx", screenW-200, 100, 10, rl.White)
	rl.DrawText(blockyTEXT, screenW-290, 110, 10, rl.White)
	rl.DrawText("blocky", screenW-200, 110, 10, rl.White)
	rl.DrawText(drawblocknextTEXT, screenW-290, 120, 10, rl.White)
	rl.DrawText("drawblocknext", screenW-200, 120, 10, rl.White)
	rl.DrawText(pmoveonTEXT, screenW-290, 130, 10, rl.White)
	rl.DrawText("pmoveon", screenW-200, 130, 10, rl.White)

	/*
		blockactiveTEXT := strconv.Itoa(blockactive)
		mouseposXTEXT := fmt.Sprintf("%.0f", mousepos.X)
		mouseposYTEXT := fmt.Sprintf("%.0f", mousepos.Y)
		pblockTEXT := strconv.Itoa(pblock)
		pblocknewTEXT := strconv.Itoa(pblocknew)

		rl.DrawText(mouseposXTEXT, screenW-290, 10, 10, rl.White)
		rl.DrawText("mouseposX", screenW-200, 10, 10, rl.White)
		rl.DrawText(mouseposYTEXT, screenW-290, 20, 10, rl.White)
		rl.DrawText("mouseposY", screenW-200, 20, 10, rl.White)
		rl.DrawText(blockactiveTEXT, screenW-290, 30, 10, rl.White)
		rl.DrawText("blockactive", screenW-200, 30, 10, rl.White)
		rl.DrawText(pblockTEXT, screenW-290, 40, 10, rl.White)
		rl.DrawText("pblock", screenW-200, 40, 10, rl.White)
		rl.DrawText(pblocknewTEXT, screenW-290, 50, 10, rl.White)
		rl.DrawText("pblocknew", screenW-200, 50, 10, rl.White)
		rl.DrawText(direction, screenW-290, 60, 10, rl.White)
		rl.DrawText("direction", screenW-200, 60, 10, rl.White)
	*/

}
func start() { // MARK: start
	camera.Zoom = 1.0
	camera.Target.X = 128.0
	camera.Target.Y = 64.0
	//	debugon = true
	// borderson = true
	//	gridon = true
	drawblocknext = 139672
	pblock = 140025
}
func updateall() { // MARK: updateall
	input()
	getactiveblock()
	updrawblock()
	animate()
	if pmoveon {
		moveplayer()
	}
	if debugon {
		debug()
	}
	if gridon {
		grid()
	}
}
func raylib() { // MARK: raylib()
	rl.InitWindow(screenW, screenH, "bloadi")
	rl.SetExitKey(rl.KeyEnd)                 // key to end the game and close window
	imgs = rl.LoadTexture("imgs_bloadi.png") // load images
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() { // MARK: WindowShouldClose

		if fullscreenon {
			rl.ToggleFullscreen()
		}

		framecount++
		mousepos = rl.GetMousePosition()
		updateall()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode2D(camera)
		// MARK: draw map layer 1
		drawx := 0
		drawy := 0
		drawblock = drawblocknext
		linecount := 0
		for a := 0; a < 112; a++ {

			checkmap := levelmap[drawblock]
			switch checkmap {
			case ".":
				tilev := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, ter1, tilev, rl.White)
			//	rl.DrawCircle(int32(drawx+128), int32(drawy+64), 10, rl.Red)
			case "#":
				tilev := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, ter2, tilev, rl.White)
				//	rl.DrawCircle(int32(drawx+128), int32(drawy+64), 10, rl.Blue)
				treev := rl.NewVector2(float32(drawx+128), float32(drawy-200))
				rl.DrawTextureRec(imgs, tree1, treev, rl.White)
			case "$":
				tilev := rl.NewVector2(float32(drawx), float32(drawy))
				rl.DrawTextureRec(imgs, ter3, tilev, rl.White)
				//	rl.DrawCircle(int32(drawx+128), int32(drawy+64), 10, rl.Green)
			}

			//  get pblocknew xy
			if drawblock == pblocknew {
				blockx = drawx
				blocky = drawy
			}
			// draw player
			if drawblock == pblock {
				rl.DrawCircle(int32(drawx+128), int32(drawy+64), 10, rl.Red)
				px = drawx
				py = drawy
			}
			// draw active block
			if drawblock == blockactive {
				//	rl.DrawCircle(int32(drawx+128), int32(drawy+64), 10, rl.Blue)
				triv1 := rl.NewVector2(float32(drawx), float32(drawy+64))
				triv2 := rl.NewVector2(float32(drawx+128), float32(drawy))
				triv3 := rl.NewVector2(float32(drawx+128), float32(drawy+128))
				triv4 := rl.NewVector2(float32(drawx+256), float32(drawy+64))
				rl.DrawTriangle(triv3, triv2, triv1, rl.Fade(rl.Black, 0.3))
				rl.DrawTriangle(triv4, triv2, triv3, rl.Fade(rl.Black, 0.3))

			}

			linecount++
			drawblock++
			drawx += 256

			if lineswitch {
				if linecount == 7 {
					drawx = 0
					drawy += 64
					linecount = 0
					drawblock += 43
					lineswitch = false
				}
			} else {
				if linecount == 7 {
					drawx = 128
					drawy += 64
					linecount = 0
					drawblock += 43
					lineswitch = true
				}
			}

		} // draw map layer 1
		// MARK: draw map layer 2

		// draw map layer 2

		rl.EndMode2D() // MARK: draw no camera

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
