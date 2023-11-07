package helper_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"sua.quadra/services/whatsapp/helper"
)

func TestTranscriptAudio(t *testing.T) {

	file := "audio.ogg"
	fileBytes, err := os.ReadFile(file)
	require.NoError(t, err)

	transcript, err := helper.OggToTranscript(fileBytes)
	require.NoError(t, err)
	require.NotEmpty(t, transcript)

}
