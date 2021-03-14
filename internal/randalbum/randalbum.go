package randalbum

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

// Album holds title and artist data
type Album struct {
	Title  string
	Artist string
}

type albums struct {
	length int
	data   []Album
}

var albms albums

// GetAlbum returns a random album from a known CSV
func GetAlbum() Album {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(albms.length)

	return albms.data[n]
}

func init() {
	log.Println("randalbum: init(): Starting album data initialization.")
	file, err := os.Open("../../assets/docs/albums.csv")
	if err != nil {
		log.Println("randalbum: init(): Could not open csv file.")
	}

	r := csv.NewReader(bufio.NewReader(file))

	firstRead := true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("randalbum: init(): Could not read record from csv file.")
		}
		if firstRead {
			firstRead = false
			continue
		}

		albms.data = append(albms.data, Album{
			Title:  record[0],
			Artist: record[1],
		})
	}

	albms.length = len(albms.data)
	log.Println("randalbum: init(): Finished.")
}
