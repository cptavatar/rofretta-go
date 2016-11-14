package main

import (
	"fmt"
	"os"
	"github.com/cptavatar/rofretta-go/detection"
)

func main() {
	detector := detection.CreateVoicingDetector()
	possibleChords := os.Args[1:]
	for _, chord := range possibleChords {
		voicings, valid := detector.ParseAndIdentifyVoicing(chord)
		if valid && len(voicings) > 0 {
			if len(voicings) == 1 {
				fmt.Printf("%s : %s \n", chord, voicings[0].Name())
			} else {
				fmt.Printf("%s : %s", chord, voicings[0].Name())
				for _, voicing := range voicings[1:] {
					fmt.Printf(", %s", voicing.Name())
				}
				fmt.Print("\n")

			}
		} else {
			fmt.Println("Could not identify ", chord)
		}
	}
}