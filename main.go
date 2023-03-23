package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a .wav file path as a command line argument")
	}

	audioFile := os.Args[1]

	// Open audio file
	f, err := os.Open(audioFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Decode the WAV file
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize speaker with the audio format
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	// Play the audio
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
