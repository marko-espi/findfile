package findfile_test

import (
	"fmt"
	"testing"
	"testing/fstest"

	"github.com/marko-espi/findfile"
)

func TestFile(t *testing.T) {
	fs := fstest.MapFS{
		"file01.txt":                 {},
		"folder01/file01.txt":        {},
		"folder01/file02.txt":        {},
		"folder02/file01.txt":        {},
		"folder03/folder01/file.txt": {},
	}
	t.Run("good path", func(t *testing.T) {
		want := 5
		got, err := findfile.File(fs, ".txt")
		if err != nil {
			fmt.Println(err)
		}

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("ext error", func(t *testing.T) {
		_, err := findfile.File(fs, "")
		if err == nil {
			t.Error("expected error but got nothing")
		}
	})
}
