package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestUploadChatFlagIsOptional(t *testing.T) {
	cmd := NewUpload()

	chatFlag := cmd.Flags().Lookup("chat")
	if chatFlag == nil {
		t.Fatal("chat flag is missing")
	}
	if chatFlag.Annotations[cobra.BashCompOneRequiredFlag] != nil {
		t.Fatal("chat flag should be optional so empty chat can use Saved Messages")
	}

	pathFlag := cmd.Flags().Lookup("path")
	if pathFlag == nil {
		t.Fatal("path flag is missing")
	}
	if pathFlag.Annotations[cobra.BashCompOneRequiredFlag] == nil {
		t.Fatal("path flag should remain required")
	}
}
