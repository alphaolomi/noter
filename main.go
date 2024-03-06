package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// Note struct to hold individual notes
type Note struct {
	ID   int
	Text string
}

// In-memory slice to hold notes
var notes []Note
var nextID = 1

// Main function to execute the root command
func main() {
	var rootCmd = &cobra.Command{Use: "notes"}
	// rootCmd.AddCommand(createCmd, readCmd, updateCmd, deleteCmd)
	rootCmd.AddCommand(createCmd, readCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Command to create a note
var createCmd = &cobra.Command{
	Use:   "create [text]",
	Short: "Create a new note",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		noteText := strings.Join(args, " ")
		note := Note{ID: nextID, Text: noteText}
		notes = append(notes, note)
		fmt.Printf("Note created: %d\n", note.ID)
		nextID++
	},
}

// Command to read all notes
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read all notes",
	Run: func(cmd *cobra.Command, args []string) {
		for _, note := range notes {
			fmt.Printf("%d: %s\n", note.ID, note.Text)
		}
	},
}

// // Command to update a note
// var updateCmd = &cobra.Command{
// 	Use:   "update [id] [text]",
// 	Short: "Update an existing note",
// 	Args:  cobra.MinimumNArgs(2),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		id, _ := strconv.Atoi(args[0])
// 		noteText := strings.Join(args[1:], " ")
// 		for i, note := range notes {
// 			if note.ID == id {
// 				notes[i].Text = noteText
// 				fmt.Printf("Note updated: %d\n", id)
// 				return
// 			}
// 		}
// 		fmt.Println("Note not found")
// 	},
// }

// // Command to delete a note
// var deleteCmd = &cobra.Command{
// 	Use:   "delete [id]",
// 	Short: "Delete a note",
// 	Args:  cobra.MinimumNArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		id, _ := strconv.Atoi(args[0])
// 		for i, note := range notes {
// 			if note.ID == id {
// 				notes = append(notes[:i], notes[i+1:]...)
// 				fmt.Printf("Note deleted: %d\n", id)
// 				return
// 			}
// 		}
// 		fmt.Println("Note not found")
// 	},
// }
