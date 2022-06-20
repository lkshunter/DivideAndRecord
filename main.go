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
	PathRecursive := flag.Bool("Recursive", false, "Lists all files in subfolder")
	mp3 := flag.Bool("MP3", true, "Should MP3-files be indexed")
	flac := flag.Bool("FLAC", false, "Should FLAC-files be indexed (default false)")
	ogg := flag.Bool("OGG", false, "Should OGG-files be indexed (default false)")

	// Sorting
	SortByTag := flag.String("SortByTag", "", "Sorting the files by metadata tag. Usable tags are: album, albumartist, artist, date, discnumber, duration, genre, publisher, title, tracknumber")

	verbose := flag.Bool("verbose", false, "Prints more information about audio files")
	export := flag.Bool("export", false, "??? (default false)")

	flag.Parse()

	if *PathToFiles == "" {
		fmt.Println("Sorry no path argument used")
		return
	}

	extensions := make([]string, 0)
	if *mp3 {
		extensions = append(extensions, ".mp3")
	}
	if *flac {
		extensions = append(extensions, ".flac")
	}
	if *ogg {
		extensions = append(extensions, ".ogg")
	}

	UsableTags := [11]string{"album", "albumartist", "artist", "date", "discnumber", "duration", "genre", "publisher", "title", "tracknumber", ""}

	checkTag := false
	for _, tag := range UsableTags {
		if 0 == strings.Compare(strings.ToLower(tag), strings.ToLower(*SortByTag)) {
			checkTag = true
		}
	}

	if !checkTag {
		panic("Wrong Tag")
	}

	// Create a list of all files
	Files := listFiles(*PathToFiles, *PathRecursive, *SortByTag, *verbose, extensions)
	divideFilesToMedium(Files, *PathToFiles, *MediumName, int64(*SideLength), int64(*ReasonableDiff), *SideAmount, *MaxMediums, *verbose, *export)
}

func listFiles(path string, rec bool, tag string, verbose bool, extensions []string) []fs.FileInfo {
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

	if tag != "" {
		sortAudioFilesByTag(AudioFiles, path, tag, verbose)
	}

	return AudioFiles
}

func tagCompare(file1 taggolib.Parser, file2 taggolib.Parser, tag string) bool {
	result := true
	switch tag {
	case "album":
		result = file1.Album() > file2.Album()
		break
	case "albumArtist":
		result = file1.AlbumArtist() > file2.AlbumArtist()
		break
	case "artist":
		result = file1.Artist() > file2.Artist()
		break
	case "date":
		result = file1.Date() > file2.Date()
		break
	case "discNumber":
		result = file1.DiscNumber() > file2.DiscNumber()
		break
	case "duration":
		result = file1.Duration() > file2.Duration()
		break
	case "genre":
		result = file1.Genre() > file2.Genre()
		break
	case "publisher":
		result = file1.Publisher() > file2.Publisher()
		break
	case "title":
		result = file1.Title() > file2.Title()
		break
	case "trackNumber":
		result = file1.TrackNumber() > file2.TrackNumber()
		break
	}
	return result
}

func sortAudioFilesByTag(files []fs.FileInfo, path string, tag string, verbose bool) []fs.FileInfo {
	for i := 0; i < len(files)-1; i++ {
		for j := 0; j < len(files)-i-1; j++ {
			file1, _ := os.Open(path + string(os.PathSeparator) + files[j].Name())
			file2, _ := os.Open(path + string(os.PathSeparator) + files[j+1].Name())
			song1, _ := taggolib.New(file1)
			song2, _ := taggolib.New(file2)

			if tagCompare(song1, song2, strings.ToLower(tag)) {
				files[j], files[j+1] = files[j+1], files[j]
			}

			file1.Close()
			file2.Close()
		}
	}
	if verbose {
		for _, file := range files {
			p, _ := os.Open(path + string(os.PathSeparator) + file.Name())
			song, _ := taggolib.New(p)
			printAudioFiles(file, song, verbose)
		}
	}

	return files
}

func printAudioFiles(file fs.FileInfo, song taggolib.Parser, verbose bool) {
	if verbose {
		fmt.Println(song.TrackNumber(), "-", song.Title(), "-", song.Album(), "-", song.Duration())
	} else {
		fmt.Println(file.Name(), song.Duration())
	}
}

func divideFilesToMedium(files []fs.FileInfo, path string, medium string, duration int64, diff int64, sides uint, maxMed uint, verbose bool, export bool) {
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
				printAudioFiles(file, song, verbose)
			} else {
				FreeSpaceAtMedium += SideMaxPlaytime - SidePlaytime
				fmt.Println("free space at", medium, MediumNumber, "side", sideNr, ":", SideMaxPlaytime-SidePlaytime)
				fmt.Println()
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
					printAudioFiles(file, song, verbose)
				}
			}
		} else {
			return
		}
		_ = p.Close()
	}

	FreeSpaceAtMedium += SideMaxPlaytime - SidePlaytime
	fmt.Println("free space at", medium, MediumNumber, "side", sideNr, ":", SideMaxPlaytime-SidePlaytime)
	fmt.Println()
	fmt.Println("free space of all recorded sides:", FreeSpaceAtMedium)

	if sideNr < int(sides) {
		for sideNr <= int(sides) {
			FreeSpaceAtMedium += SidePlaytime
			sideNr++
		}
	}

	fmt.Println("free space of all recorded medium:", FreeSpaceAtMedium)
}
