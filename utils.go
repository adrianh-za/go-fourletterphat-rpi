// MIT License

// Copyright (c) 2019 Adrian Houghton

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE

// NB: Various information was gathered from https://github.com/pimoroni/fourletter-phat

package fourletterphat

import (
	"strconv"
	"time"
	"errors"
	"strings"
	i2c "github.com/d2r2/go-i2c"
	logger "github.com/d2r2/go-logger"
)

const (
	AddressDefault        		= 0x70
	AddressBlinkCommand     	= 0x80
	AddressBlinkDisplayOn   	= 0x01
	AddressSystemSetup      	= 0x20
	AddressOscillator       	= 0x01
	AddressCommandBrightness	= 0xE0
)

// LEDBlink represents a blink setting that can be processed by the LED display
type LEDBlink byte

const (
	// BlinkOff - don't blink the display
	BlinkOff         LEDBlink = 0x00
	// Blink2Hz - blink at 2Hz
	Blink2Hz         LEDBlink = 0x02
	// Blink1Hz - blink at 1Hz
	Blink1Hz         LEDBlink = 0x04
	// BlinkHalfHz - blink at half of 1Hz
	BlinkHalfHz      LEDBlink = 0x06
 )

 // Initialize sets up the LED display
func Initialize(bus *i2c.I2C) (error) {
	// Change logger to not output Debug
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)

	// Instatiate LED device
	_, err := bus.WriteBytes([]byte{ AddressSystemSetup | AddressOscillator })
	if (err != nil) {
		return err
	}

	// Ready the LED device
	ClearChars(bus)
	SetBrightness(bus, 15)
	SetBlink(bus, BlinkOff)

	return nil
}

// SetBrightness sets the brightness of the LED display
// Valid brightness value is 0 to 15
func SetBrightness(bus *i2c.I2C, value int) (error) {
	
	if (value < 0) || (value > 15) {
		errors.New("Brightness value must be of range 0 to 15")
	}
	
	_, err := bus.WriteBytes([]byte{ AddressCommandBrightness | byte(value) })
	return err
}

// SetBlink sets the blink frequency of the LED display
func SetBlink(bus *i2c.I2C, blink LEDBlink) (error) {
	_, err := bus.WriteBytes([]byte{ AddressBlinkCommand | AddressBlinkDisplayOn | byte(blink) })
	return err
}

// ClearChars clears the LED display
func ClearChars(bus *i2c.I2C) (error) {
	var ledChars = []byte{0,0,0,0,0,0,0,0,0}	//Default "empty" array
	bus.WriteBytes(ledChars)
	return nil
}

// WriteCharacters allows you to display characters on the LED display
func WriteCharacters(bus *i2c.I2C, text string) (error) {
	
	// Check that string is not too long (max 4 characters)
	if (len(strings.ReplaceAll(text, ".", "")) > 4)  {
		errors.New("Only 4 characters can be specified")
	}
	
	var ledChars = []byte{0,0,0,0,0,0,0,0,0}	//Default "empty" array

	//Iterate the string
	var count int
	for _, rune := range text {
		var inputChar = CharMap[rune]
		var position = byte(count * 2) + 1

		//Charcter is ".", set set bit 7 on HI for previous character
		if (rune == '.') {
			var periodHi = ledChars[position - 1];
			periodHi |= (1 << 6)
			ledChars[position - 1] = periodHi
		} else {
			var charInt = getIntFromBinaryChar(inputChar)
			var hi, lo = getHighLowByte(charInt)
			ledChars[position] = lo;
			ledChars[position + 1] = hi;

			count++
		}
	}
	
	// Write to LED display
	bus.WriteBytes(ledChars)
	return nil
}

// ScrollCharacters will display the inputted text and scroll characters to the left
// if the length of the text is greater than 4 characters
// "." will be displayed for a character when it comes after a charecter.
// Mulltiple successsive "." will be ignored, only first "." will be displayed
func ScrollCharacters(bus *i2c.I2C, text string, delayMS int, pad bool) error {

	//Remove all duplicates of "."
	for {
		if !strings.Contains(text, "..") {
			break
		}

		text = strings.ReplaceAll(text, "..", ".")
	}

	//Add 4 spaces onto beginning and end of text to make a nice scrolling in/out effect
	if (pad) {
		text = "    " + text + "    "
	}

	//Send each string of 4 characters through to be displayed (exclude "."" though), delaying by specified milliseonds between each display
	var count int
	for {

		var charCount int
		var displayText string
		for {

			//Do check to see if we have reached end of text as we don't want to overflow
			if count+charCount+1 > len(text) {
				break
			}

			//Get the current character
			var currentChar = text[count+charCount : count+charCount+1]

			// First character is "." so skip this iteration ("." must come after letter when displaying)
			if (currentChar == ".") && (charCount == 0) {
				break
			}

			// Add the character to text to display
			displayText = displayText + currentChar

			// Check if there is character after current character, we want to see if it is "."
			var nextChar string
			if len(text) >= count+charCount+2 {
				nextChar = text[count+charCount+1 : count+charCount+2]
				if nextChar == "." {
					displayText = displayText + nextChar
					charCount++
				}
			}

			charCount++

			//Break after adding 4 characters (excluding ".") or if end reached
			if (len(strings.ReplaceAll(displayText, ".", "")) > 3) || (len(displayText) == len(text)) {
				WriteCharacters(bus, displayText)
				time.Sleep(time.Duration(delayMS) * time.Millisecond)
				break
			}
		}

		// Do the looping, break if limit reached
		count++
		if count >= len(text) {
			break
		}
	}

	return nil
}

// getIntFromBinaryChar returns the UINT16 represented by the specified LEDChar.
func getIntFromBinaryChar(binaryChar LEDChar) uint16 {
	result, _ := strconv.ParseUint(string(binaryChar), 2, 16)
	return uint16(result)
}

// getHighLowByte returns the high and low byte of the specified unsigned int16
func getHighLowByte(value uint16) (uint8, uint8) {
	var h, l uint8 = uint8(value >> 8), uint8(value & 0xff)
	return h, l
}