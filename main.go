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
	video    Video
)

func main() {
	system.PC = uint16(0x200)
	video.SetWidthHeight(64, 32)
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
		c := make(chan chip_8.Signal)
		defer close(c)
		fmt.Println("Now running ", rom_path)
		video.SetTitle("Chip-8 Window: " + rom_path)
		system.LoadGame(bytes)
		go system.RunCPU(c)
		video.Initialize(c)
	}
}
