package tests

import (
	"io/ioutil"
	"os"
	"testing"

	"Notes/cmd"

	"github.com/spf13/cobra"
)

var notesMap = make(map[string]string)

func TestEdit(t *testing.T) {

	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte("test\n")); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	cmd.Edit.Run(&cobra.Command{}, []string{tmpfile.Name()})
	data, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "TEST\n" {
		t.Fatalf("Ожидалось TEST, получено %s", data)
	}
}

func TestRead(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	expected := "test\n"
	if _, err := tmpfile.Write([]byte(expected)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	cmd.Read.Run(&cobra.Command{}, []string{tmpfile.Name()})
	data, err := ioutil.ReadFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != expected {
		t.Fatalf("Ожидалось %s, получено %s", expected, data)
	}
}

// Тест для команды CreateCMD
func TestCreateCMD(t *testing.T) {
	// Создаем имя файла для тестирования
	filename := "testfile"

	// Запускаем команду CreateCMD на нашем файле
	cmd.CreateCMD.Run(&cobra.Command{}, []string{filename})

	// Проверяем, что файл был создан
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("Файл %s не был создан", filename)
	}

	// Удаляем файл после теста
	os.Remove(filename)
}
