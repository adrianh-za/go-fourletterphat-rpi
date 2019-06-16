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

package main

import (
	"fmt"
	"time"
	flp "github.com/adrianh-za/go-fourletterphat-rpi"
	"github.com/d2r2/go-i2c"
)

func main() {

	// Initialize the I2C 
	i2c, err := i2c.NewI2C(flp.AddressDefault, 1)
	defer i2c.Close()

	if (err != nil) {
		fmt.Println(err)
	}
	
	// Initialize the LED display
	flp.Initialize(i2c)
	flp.SetBrightness(i2c, 15)
	
	
	// Lets display some static text
	flp.WriteCharacters(i2c, "1")
	time.Sleep(1 * time.Second)
	flp.WriteCharacters(i2c, "22")
	time.Sleep(1 * time.Second)
	flp.WriteCharacters(i2c, "333")
	time.Sleep(1 * time.Second)
	flp.WriteCharacters(i2c, "4444")
	time.Sleep(1 * time.Second)
	
	// Lets display some scolling text
	flp.ScrollCharacters(i2c, "1234567890", 300, true)
	time.Sleep(1 * time.Second)
	flp.ScrollCharacters(i2c, "UPPER CASE - SCROLLING WITH PADDING.", 220, true)
	time.Sleep(1 * time.Second)
	flp.ScrollCharacters(i2c, "UPPER CASE - SCROLLING WITH NO PADDING.", 220, false)
	time.Sleep(1 * time.Second)
	flp.ScrollCharacters(i2c, "LOWER CASE - scrolling with padding.", 220, true)
	time.Sleep(1 * time.Second)
	flp.ScrollCharacters(i2c, "SYMBOLS - {} [] () <> ,. /|\\ @#$~%^@&*+=_-;:'`", 250, true)
	time.Sleep(1 * time.Second)
	
	//Finish up
	flp.ClearChars(i2c)
}
