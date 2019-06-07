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

// Grabbed the binary values for letters/numbers/symbols from Pimoroni Python source code
// https://github.com/pimoroni/fourletter-phat/blob/master/library/fourletterphat/alphanum4.py

// LEDChar represents a character that can be displayed by the LED display
type LEDChar string

const (
	// CharEmpty is a character where no LEDs are lit
	CharEmpty 		LEDChar = "0000000000000000"
	// CharAll is a character where all LEDs are lit
	CharAll 		LEDChar = "1111111111111111"
)

// CharMap is a list of all characters that can be displayed with there associated binary bit mask
// that tells the display which parts of the LED to light up.
var CharMap = map[rune]LEDChar{
	' ': "0000000000000000",
	'!': "0000000000000110",
	'"': "0000001000100000",
	'#': "0001001011001110",
	'$': "0001001011101101",
	'%': "0000110000100100",
	'&': "0010001101011101",
	'\'':"0010000100000000",
	'(': "0010010000000000",
	')': "0000100100000000",
	'*': "0011111111000000",
	'+': "0001001011000000",
	',': "0000100000000000",
	'-': "0000000011000000",
	'.': "0000000000000000",
	'/': "0000110000000000",
	'0': "0000110000111111",
	'1': "0000000000000110",
	'2': "0000000011011011",
	'3': "0000000010001111",
	'4': "0000000011100110",
	'5': "0010000001101001",
	'6': "0000000011111101",
	'7': "0000000000000111",
	'8': "0000000011111111",
	'9': "0000000011101111",
	':': "0001001000000000",
	';': "0000101000000000",
	'<': "0010010000000000",
	'=': "0000000011001000",
	'>': "0000100100000000",
	'?': "0001000010000011",
	'@': "0000001010111011",
	'A': "0000000011110111",
	'B': "0001001010001111",
	'C': "0000000000111001",
	'D': "0001001000001111",
	'E': "0000000011111001",
	'F': "0000000001110001",
	'G': "0000000010111101",
	'H': "0000000011110110",
	'I': "0001001000000000",
	'J': "0000000000011110",
	'K': "0010010001110000",
	'L': "0000000000111000",
	'M': "0000010100110110",
	'N': "0010000100110110",
	'O': "0000000000111111",
	'P': "0000000011110011",
	'Q': "0010000000111111",
	'R': "0010000011110011",
	'S': "0000000011101101",
	'T': "0001001000000001",
	'U': "0000000000111110",
	'V': "0000110000110000",
	'W': "0010100000110110",
	'X': "0010110100000000",
	'Y': "0001010100000000",
	'Z': "0000110000001001",
	'[': "0000000000111001",
	'\\':"0010000100000000",
	']': "0000000000001111",
	'^': "0000110000000011",
	'_': "0000000000001000",
	'`': "0000000100000000",
	'a': "0001000001011000",
	'b': "0010000001111000",
	'c': "0000000011011000",
	'd': "0000100010001110",
	'e': "0000100001011000",
	'f': "0000000001110001",
	'g': "0000010010001110",
	'h': "0001000001110000",
	'i': "0001000000000000",
	'j': "0000000000001110",
	'k': "0011011000000000",
	'l': "0000000000110000",
	'm': "0001000011010100",
	'n': "0001000001010000",
	'o': "0000000011011100",
	'p': "0000000101110000",
	'q': "0000010010000110",
	'r': "0000000001010000",
	's': "0010000010001000",
	't': "0000000001111000",
	'u': "0000000000011100",
	'v': "0010000000000100",
	'w': "0010100000010100",
	'x': "0010100011000000",
	'y': "0010000000001100",
	'z': "0000100001001000",
	'{': "0000100101001001",
	'|': "0001001000000000",
	'}': "0010010010001001",
	'~': "0000010100100000",
}