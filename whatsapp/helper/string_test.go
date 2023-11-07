package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseJid(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"5511989070670@whasapp.com", true},
		{"5511989070670", true},
		{"11989070670", false},
		{"11989070670@whasapp.com", false},
		{"+5511989070670", true},
	}

	for _, tc := range testCases {
		_, actual := ParseJID(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, actual)
		}
	}
}

func TestExtractPhoneFromVcard(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"BEGIN:VCARD\nVERSION:3.0\nN:;Test;;;\nFN:Test\nTEL;type=CELL;type=VOICE;waid=5511989070670:+55 11 98907-0670\nEND:VCARD", "5511989070670"},
		{"BEGIN:VCARD\nVERSION:3.0\nN:;Test;;;\nFN:Test\nTEL;type=CELL;type=VOICE;waid=5511989070670:+55 11 98907-0670\nTEL;type=CELL;type=VOICE;waid=5511989070670:+55 11 98907-0670\nEND:VCARD", "5511989070670,5511989070670"},
	}

	for _, tc := range testCases {
		actual := ExtractPhoneFromVcard(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, actual)
		}
	}
}

func TestExtractLinkFromFile(t *testing.T) {
	test := "<https://sq-atendimento.slack.com/files/U04VBSBNAGM/F052KR7EB54/luiz_paulo_cafeteria____im__vel_rua_ministro_jesu__no_cardoso__120.pdf|https://sq-atendimento.slack.com/files/U04VBSBNAGM/F052KR7EB54/luiz_paulo_cafeteria____im__vel_rua_ministro_jesu__no_cardoso__120.pdf>"
	expected := "https://sq-atendimento.slack.com/files/U04VBSBNAGM/F052KR7EB54/luiz_paulo_cafeteria____im__vel_rua_ministro_jesu__no_cardoso__120.pdf"
	actual := ExtractLinkFromFile(test)

	require.Equal(t, expected, actual)

}
