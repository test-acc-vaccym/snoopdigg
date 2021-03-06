package main

import (
	"os"
	"fmt"
	"path/filepath"
	"encoding/json"
	"github.com/botherder/go-files"
	"github.com/botherder/go-autoruns"
	log "github.com/Sirupsen/logrus"
)

func generateAutoruns() {
	log.Info("Identifying files marked for persistence...")

	// Fetch autoruns.
	autoruns := autoruns.Autoruns()

	// Make backup of autoruns executables.
	for _, autorun := range(autoruns) {
		if _, err := os.Stat(autorun.ImagePath); err == nil {
			copyName := fmt.Sprintf("%s_%s.bin", autorun.MD5, autorun.ImageName)
			copyPath := filepath.Join(acq.Autoruns, copyName)
			files.Copy(autorun.ImagePath, copyPath)
		}
	}

	// Store the json list to file.
	autorunsJSONPath := filepath.Join(acq.Storage, "autoruns.json")
	autorunsJSON, err := os.Create(autorunsJSONPath)
	if err != nil {
		log.Error("Unable to save autoruns to file: ", err.Error())
		return
	}
	defer autorunsJSON.Close()

	// Encoding into json.
	buf, _ := json.MarshalIndent(autoruns, "", "    ")

	autorunsJSON.WriteString(string(buf[:]))
	autorunsJSON.Sync()

	log.Info("Autoruns collected successfully!")
}
