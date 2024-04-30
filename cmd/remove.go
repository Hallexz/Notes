package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func deleteNote(noteID string) error {
	notesMap := make(map[string]string)

	if _, exists := notesMap[noteID]; !exists {
		return fmt.Errorf("заметка с идентификатором %s не найдена", noteID)
	}

	delete(notesMap, noteID)
	return nil
}

var Remove = &cobra.Command{
	Use:   "deleteNote",
	Short: "Удалить заметку",
	Long:  `Эта команда удаляет заметку.`,
	Run: func(cmd *cobra.Command, args []string) {
		noteID := args[0]
		err := deleteNote(noteID)
		if err != nil {
			fmt.Println("Ошибка при удалении заметки:", err)
			return
		}

		fmt.Println("Заметка удалена")
	},
}
