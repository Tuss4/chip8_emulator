// Referencing:
// http://en.wikipedia.org/wiki/CHIP-8
// http://devernay.free.fr/hacks/chip8/C8TECH10.HTM
// http://www.multigesture.net/articles/how-to-write-an-emulator-chip-8-interpreter/

package main

import (
	"fmt"
	"github.com/tuss4/chip8_emulator/chip_8"
	"io/ioutil"
	"log"
	"os"
)

var (
	system   chip_8.CPU
	rom_path string
	vid      Video
)

func main() {
	vid.width = 640
	vid.height = 320
	system.PC = uint16(0x200)
	if len(os.Args) < 2 {
		fmt.Println("Please specify the path to a rom.")
	} else {
		rom_path = os.Args[len(os.Args)-1]
	}
	// Set up the reading of the bytes
	if rom_path != "" {
		bytes, err := ioutil.ReadFile(rom_path)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Now running ", rom_path)
		vid.title = "Chip-8 Window: " + rom_path
		vid.Initialize()
		system.LoadGame(bytes)
		system.RunCPU()
	} else {
		fmt.Println("No rom specified, dude.")
	}
}
