// Command go-simulator-motion is a go-simulator plugin with the motions of a living scene: the
// passing of the day, wind gusts, drifting clouds, tides and the motions entities ask for in
// their properties (bob, spin, path). The host starts it; it talks over standard input and output.
package main

import (
	"log"

	"github.com/akrck02/go-simulator-motion/motion"
)

func main() {
	log.SetFlags(0)
	if err := motion.Plugin().Serve(); err != nil {
		log.Fatal(err)
	}
}
