package fdbmeter

import (
	"embed"
	"encoding/json"
	"io/fs"
	"testing"
)

//go:embed testdata/*.json
var statusFixtures embed.FS

func TestStatusDecode(t *testing.T) {
	files, err := fs.Glob(statusFixtures, "testdata/*.json")
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			f, err := statusFixtures.Open(file)
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if err := f.Close(); err != nil {
					t.Fatal(err)
				}
			}()
			var status Status
			if err := json.NewDecoder(f).Decode(&status); err != nil {
				t.Fatal(err)
			}
		})
	}
}
