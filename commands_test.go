package main

import (
	"bytes"
	// "fmt"
	// "strings"
	"testing"

	"github.com/spf13/cobra"
)

// TestCreateNote tests the creation of a note.
func TestCreateNote(t *testing.T) {
	// Setup - reset the environment
	notes = []Note{} // Reset notes slice to empty
	nextID = 1       // Reset next ID to 1

	// Create a new Cobra command to test
	var testCmd = &cobra.Command{
		Use:   "test",
		Short: "A test command",
	}
	testCmd.AddCommand(createCmd) // Add the create command to this test command

	// Execute the command with a test note
	args := "This is a test note"
	b := bytes.NewBufferString("")
	testCmd.SetOut(b)
	testCmd.SetArgs([]string{"create", args})
	if err := testCmd.Execute(); err != nil {
		t.Fatalf("Unexpected error executing command: %v", err)
	}

	// Verify the note was added correctly
	expectedNote := Note{ID: 1, Text: args}
	if len(notes) != 1 || notes[0] != expectedNote {
		t.Errorf("Expected note was not added correctly. Got %v, want %v", notes, expectedNote)
	}

	// Optionally, verify output
	// expectedOutput := fmt.Sprintf("Note created: %d\n", expectedNote.ID)
	// if gotOutput := b.String(); !strings.Contains(gotOutput, expectedOutput) {
	// 	t.Errorf("Unexpected command output. Got %s, want to contain %s", gotOutput, expectedOutput)
	// }
}
