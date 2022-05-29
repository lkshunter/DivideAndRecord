package main

import (
	"flag"
	"fmt"
	"github.com/mdlayher/taggolib"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	// Information about the Medium
	SideLength := flag.Uint("SideLength", 30, "Number of minutes playing time per page of the medium")
	SideAmount := flag.Uint("SideAmount", 2, "The amount of sides of the medium")
	MediumName := flag.String("MediumName", "Cassette", "Name of the medium like cassette or CD")
	ReasonableDiff := flag.Uint("ReasonableDiff", 0, "Reasonable difference between the maximum playing time of the medium and the end of a track in seconds")
	MaxMediums := flag.Uint("MaxMediums", 0, "In case the number of media to be recorded is limited")

	// Information about the Files
	PathToFiles := flag.String("Path", "", "Location of the files to be recorded")
	PathRecrusiv := flag.Bool("Recrusiv", false, "Lists all files in subfolders")
	mp3 := flag.Bool("MP3", true, "Should MP3-files be indexed")
	flac := flag.Bool("FLAC", false, "Should FLAC-files be indexed (default false)")
	ogg := flag.Bool("OGG", false, "Should OGG-files be indexed (default false)")

	// Sorting
	// SortByName := flag.Bool("SortByName", false, "Sorting the files by name")
	// SortByTrackNumber := flag.Bool("SortByTrackNumber", true, "Sorting the files by the track number")

	export := flag.Bool("export", false, "??? (default false)")

	flag.Parse()

	if *PathToFiles == "" {
		fmt.Println("Sorry no path argument used")
		return
	}

	extensions := make([]string, 0)
	if *mp3 == true {
		extensions = append(extensions, ".mp3")
	}
	if *flac == true {
		extensions = append(extensions, ".flac")
	}
	if *ogg == true {
		extensions = append(extensions, ".ogg")
	}

	// Create a list of all files
	Files := listFiles(*PathToFiles, *PathRecrusiv, extensions)
	divideFilesToMedium(Files, *PathToFiles, *MediumName, int64(*SideLength), int64(*ReasonableDiff), *SideAmount, *MaxMediums, *export)
}

func listFiles(path string, rec bool, extensions []string) []fs.FileInfo {
	AudioFiles := make([]fs.FileInfo, 0)
	files, err := ioutil.ReadDir(path)

	//walk, _ := filepath.Walk(path, ".")
	//fmt.Println(walk)

	if err != nil {
		log.Fatal(err)
	}

	for _, ext := range extensions {
		for _, file := range files {
			if strings.Contains(strings.ToLower(file.Name()), ext) {
				AudioFiles = append(AudioFiles, file)
			}
		}
	}

	return AudioFiles
}

func divideFilesToMedium(files []fs.FileInfo, path string, medium string, duration int64, diff int64, sides uint, maxMed uint, export bool) {
	var SidePlaytime time.Duration = 0 // Time
	var SideMaxPlaytime = time.Duration(duration * 60000000000)
	var MediumNumber = 1
	var sideNr = 1
	var FreeSpaceAtMedium time.Duration = 0

	fmt.Println(medium, ":", MediumNumber, "Side :", sideNr)
	for _, file := range files {
		p, _ := os.Open(path + string(os.PathSeparator) + file.Name())
		song, _ := taggolib.New(p)

		if int(maxMed) >= MediumNumber || maxMed == 0 {
			if (SidePlaytime + song.Duration()) <= (SideMaxPlaytime + time.Duration(diff*1000000000)) {
				SidePlaytime += song.Duration()
				fmt.Println(file.Name(), song.Duration())
			} else {
				FreeSpaceAtMedium += SideMaxPlaytime - SidePlaytime
				fmt.Println("free space at", medium, MediumNumber, "side", sideNr, ":", SideMaxPlaytime-SidePlaytime, "\n")
				SidePlaytime = 0

				if sideNr == int(sides) {
					sideNr = 1
					MediumNumber += 1
				} else {
					sideNr += 1
				}

				if int(maxMed) >= MediumNumber {
					fmt.Println(medium, ":", MediumNumber, "Side :", sideNr)

					SidePlaytime += song.Duration()
					fmt.Println(file.Name(), song.Duration())
				}
			}
		} else {
			return
		}
		_ = p.Close()
	}
	FreeSpaceAtMedium += SideMaxPlaytime - SidePlaytime
	fmt.Println("free space at", medium, MediumNumber, "side", sideNr, ":", SideMaxPlaytime-SidePlaytime, "\n")
	fmt.Println("free space of all medium:", FreeSpaceAtMedium)
}
