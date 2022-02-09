package soundboard

import (
	"io/ioutil"
	"log"
	"strings"
)

type Audio struct {
	Title    string
	FileName string
}

var clips []Audio

func GetClips() []Audio {
	return clips
}

func GetAudioFiles(folderName string) []Audio {
	path := "../../assets/audio/" + folderName + "/"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("soundboard: init(): Failed to read audio directory.")
	}

	var audioFiles []Audio

	for _, f := range files {
		fileName := strings.TrimSuffix(f.Name(), ".mp3")
		title := strings.Title(strings.ReplaceAll(fileName, "-", " "))
		audioFiles = append(audioFiles, Audio{title, folderName + "/" + fileName})
	}

	return audioFiles
}

func init() {
	log.Println("soundboard: init(): Starting soundboard data initialization.")

	clips = GetAudioFiles("clips")

	log.Println("soundboard: init(): Finished.")
}
