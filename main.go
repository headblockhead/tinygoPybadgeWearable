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
	released := true
	for {
		pressed, _ := buttons.ReadInput()

		if released && buttons.Pins[shifter.BUTTON_SELECT].Get() {
			if CurrentMode < 2 {
				CurrentMode += 1
			} else {
				CurrentMode = 0
			}
		}
		if released && buttons.Pins[shifter.BUTTON_START].Get() {
			break
		}
		if pressed == 0 {
			released = true
		} else {
			released = false
		}

		if CurrentMode != OldMode {
			if CurrentMode == 0 {
				display.FillRectangleWithBuffer(0, 0, 160, 128, logoRGBA1)
			} else if CurrentMode == 1 {
				display.FillRectangleWithBuffer(0, 0, 160, 128, logoRGBA2)
			} else if CurrentMode == 2 {
				display.FillRectangleWithBuffer(0, 0, 160, 128, logoRGBA3)
			} else {
				display.FillRectangleWithBuffer(0, 0, 160, 128, logoRGBA1)
			}
			OldMode = CurrentMode
		}
		time.Sleep(3 * time.Millisecond)
	}
}
