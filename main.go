package main

import (
	"context"
	"fmt"
	"os"

	"github.com/deepgram/deepgram-go-sdk"
)

func main() {
	dgClient, err := deepgram.New("<YOUR_DEEPGRAM_API_KEY>")
	if err != nil {
		fmt.Println("Failed to create Deepgram client:", err)
		return
	}

	audioFile, err := os.Open("path_to_your_audio_file.wav")
	if err != nil {
		fmt.Println("Failed to open audio file:", err)
		return
	}
	defer audioFile.Close()

	opts := deepgram.TranscriptionOptions{
		Model: "general", // Change as needed
	}

	transcript, err := dgClient.Transcription().FromStream(context.Background(), audioFile, opts)
	if err != nil {
		fmt.Println("Failed to transcribe audio:", err)
		return
	}

	fmt.Println("Transcription:", transcript.Results.Channels[0].Alternatives[0].Transcript)
}
