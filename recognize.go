// Sample speech-quickstart uses the Google Cloud Speech API to transcribe
// audio.
package main

import (
	"context"
	"fmt"

	speech "cloud.google.com/go/speech/apiv1p1beta1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"
)

func Recognize(fileData []byte) (string, error) {
	ctx := context.Background()

	// Creates a client.
	client, err := speech.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to create client: %v", err)
	}

	// Detects speech in the audio file.
	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_MP3,
			LanguageCode:    "en-US",
			SampleRateHertz: 48000,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: fileData},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to recognize: %v", err)
	}

	// Prints the results.
	transcript := ""
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			transcript = fmt.Sprintf("\"%v\" (confidence=%3f)\n", alt.Transcript, alt.Confidence)
		}
	}

	return transcript, nil
}
