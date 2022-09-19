package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7735"

	"tinygo.org/x/drivers/shifter"
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
		Frequency: 8000000,
	})
	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})

	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7735.Config{
		Rotation: st7735.ROTATION_90,
	})

	var buttons shifter.Device
	buttons = shifter.NewButtons()
	buttons.Configure()

	display.EnableBacklight(true)
	display.FillScreen(color.RGBA{R: 0, G: 0, B: 0})
	CurrentMode := 0
	OldMode := -1
	funMode := false
	released := true

	objectX := int16(10)
	objectY := int16(10)
	objectVelocityX := int16(2)
	objectVelocityY := int16(4)

	for {
		pressed, _ := buttons.ReadInput()

		if released && buttons.Pins[shifter.BUTTON_START].Get() {
			funMode = !funMode
			if !funMode {
				OldMode = -1
			}
			// Debounce
			time.Sleep(100 * time.Millisecond)
		}
		if released && buttons.Pins[shifter.BUTTON_SELECT].Get() {
			if CurrentMode < 2 {
				CurrentMode += 1
			} else {
				CurrentMode = 0
			}
			// Debounce
			time.Sleep(100 * time.Millisecond)
		}

		if pressed == 0 {
			released = true
		} else {
			released = false
		}
		//|| funMode
		if OldMode != CurrentMode {
			display.FillRectangleWithBuffer(0, 0, 160, 36, logoRGBAheader)
			if CurrentMode == 0 {
				display.FillRectangleWithBuffer(0, 36, 160, 71, logoRGBA2)
			} else if CurrentMode == 1 {
				display.FillRectangleWithBuffer(0, 36, 160, 71, logoRGBA3)
			} else if CurrentMode == 2 {
				display.FillRectangleWithBuffer(0, 36, 160, 71, logoRGBA4)
			} else {
				display.FillScreen(color.RGBA{R: 0, G: 0, B: 0})
			}
			display.FillRectangleWithBuffer(0, 107, 160, 21, logoRGBAfooter)
			OldMode = CurrentMode
		}
		if funMode {
			objectX += objectVelocityX
			objectY += objectVelocityY
			if objectX < 0 || objectX+28 > 160 {
				if objectX < 0 {
					objectX = 1
				} else if objectX+28 > 160 {
					objectX = 160 - 29
				}
				objectVelocityX = -objectVelocityX
			}
			if objectY < 0 || objectY+28 > 128 {
				if objectY < 0 {
					objectY = 1
				} else if objectY+28 > 128 {
					objectY = 128 - 29
				}
				objectVelocityY = -objectVelocityY
			}
			for x := int16(0); x < 28; x++ {
				display.DrawFastVLine(objectX+x, objectY, objectY+28, color.RGBA{R: 255, G: 255, B: 255})
			}
			display.DrawFastVLine(objectX+3, objectY+3, objectY+19, color.RGBA{R: 255, G: 102, B: 255})
			display.DrawFastVLine(objectX+20, objectY+3, objectY+7, color.RGBA{R: 255, G: 102, B: 255})
			display.DrawFastVLine(objectX+6, objectY+7, objectY+17, color.RGBA{R: 255, G: 102, B: 255})
			display.DrawFastVLine(objectX+20, objectY+17, objectY+20, color.RGBA{R: 255, G: 102, B: 255})
			display.DrawFastHLine(objectX+3, objectX+20, objectY+3, color.RGBA{R: 255, G: 102, B: 255})
			display.DrawFastHLine(objectX+6, objectX+20, objectY+7, color.RGBA{R: 255, G: 102, B: 255})
			display.DrawFastHLine(objectX+6, objectX+20, objectY+17, color.RGBA{R: 255, G: 102, B: 255})
			display.DrawFastHLine(objectX+3, objectX+20, objectY+20, color.RGBA{R: 255, G: 102, B: 255})

			display.DrawFastVLine(objectX+9, objectY+9, objectY+20, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastVLine(objectX+26, objectY+10, objectY+13, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastVLine(objectX+12, objectY+13, objectY+20, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastVLine(objectX+9, objectY+23, objectY+27, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastVLine(objectX+23, objectY+17, objectY+23, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastVLine(objectX+26, objectY+17, objectY+27, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastHLine(objectX+9, objectX+26, objectY+10, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastHLine(objectX+12, objectX+26, objectY+13, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastHLine(objectX+9, objectX+12, objectY+20, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastHLine(objectX+23, objectX+26, objectY+17, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastHLine(objectX+9, objectX+23, objectY+23, color.RGBA{R: 255, G: 76, B: 38})
			display.DrawFastHLine(objectX+9, objectX+26, objectY+27, color.RGBA{R: 255, G: 76, B: 38})

			time.Sleep(30 * time.Millisecond)
		}

	}
}
