package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Используем глобальную переменную notesMap
// Используем глобальную переменную notesMap
var notesMap = make(map[string]string)

func deleteNote(noteID string) error {
	// Проверяем, есть ли заметки для удаления
	if len(notesMap) == 0 {
		return fmt.Errorf("нет заметок для удаления")
	}

	// Проверяем, существует ли заметка с таким идентификатором
	if _, exists := notesMap[noteID]; !exists {
		return fmt.Errorf("заметка с идентификатором %s не найдена", noteID)
	}

	// Удаляем заметку
	delete(notesMap, noteID)
	return nil
}

var Remove = &cobra.Command{
	Use:   "deleteNote",
	Short: "Удалить заметку",
	Long:  `Эта команда удаляет заметку.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Проверяем, есть ли аргументы команды
		if len(args) == 0 {
			fmt.Println("Ошибка: не указан идентификатор заметки для удаления")
			return
		}

		noteID := args[0]
		err := deleteNote(noteID)
		if err != nil {
			fmt.Println("Ошибка при удалении заметки:", err)
			return
		}

		fmt.Println("Заметка удалена")
	},
}
