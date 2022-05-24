package main

import (
	"flag"
	"fmt"
	"gitlab.com/anost/audiotags"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	SideLength := flag.Int("SideLength", 30, "Number of minutes playing time per page of the medium")
	SideAmount := flag.Int("SideAmount", 2, "The amount of sides of the medium")
	MaxMediums := flag.Int("MaxMediums", -1, "In case the number of media to be recorded is limited")

	MediumName := flag.String("MediumName", "Cassette", "")
	PathToFiles := flag.String("Path", "", "Location of the files to be recorded")
	//FileExtensions := flag.String("FileExtensions", ".mp3", "List of all file extensions")

	SortByName := flag.Bool("SortByName", false, "Sorting the files by name")
	SortByTrackNumber := flag.Bool("SortByTrackNumber", true, "Sorting the files by the track number")

	flag.Parse()

	fmt.Println("Name des Medium:", *MediumName)
	fmt.Println("Spielzeit in Minuten pro Seite:", *SideLength)
	fmt.Println("Anzahl der Seite:", *SideAmount)
	fmt.Println("Anzahl der Medien:", *MaxMediums)
	fmt.Println("Speicherort der Dateien:", *PathToFiles)
	fmt.Println("Sortieren nach Name:", *SortByName)
	fmt.Println("Sortieren nach Track:", *SortByTrackNumber)
	//fmt.Println("Liste aller Dateiendungen:", *FileExtensions)

	files := listFiles(*PathToFiles, ".mp3")

	readLength(files, *PathToFiles)
}

func listFiles(path string, extensions ...string) []fs.FileInfo {

	AudioFiles := make([]fs.FileInfo, 0)
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, ext := range extensions {
		for _, file := range files {
			if strings.Contains(file.Name(), ext) {
				AudioFiles = append(AudioFiles, file)
			}
		}
	}

	return AudioFiles
}

func readLength(files []fs.FileInfo, path string) {

	for _, file := range files {
		if file.IsDir() == false {

			pathToFile := path + "\\" + file.Name()

			test, _ := audiotags.ReadAudioProperties(pathToFile)

			fmt.Println(test.LengthMs)

			test = nil

		}
	}
}
