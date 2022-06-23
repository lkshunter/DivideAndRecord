package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/mdlayher/taggolib"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	PathRecursive := flag.Bool("Recursive", false, "Generates for every subfolder his own report (default false)")
	mp3 := flag.Bool("MP3", true, "Should MP3-files be indexed")
	flac := flag.Bool("FLAC", false, "Should FLAC-files be indexed (default false)")
	ogg := flag.Bool("OGG", false, "Should OGG-files be indexed (default false)")

	// Sorting
	SortByTag := flag.String("SortByTag", "", "Sorting the files by metadata tag. Usable tags are: album, albumartist, artist, date, discnumber, duration, genre, publisher, title, tracknumber")

	// other Tags
	verbose := flag.Bool("verbose", false, "Prints more information about audio files")
	export := flag.Bool("export", false, "??? (default false)")

	flag.Parse()

	// checks if the path exists
	if _, err := os.Stat(*PathToFiles); os.IsNotExist(err) {
		panic(err)
	}

	// checks if at least one file extension is selected
	extensions, ErrExtensions := SelectExtensions(*mp3, *flac, *ogg)
	if ErrExtensions != nil {
		panic(ErrExtensions)
	}

	// checks if the tag is usable
	checkTag, ErrTag := CheckSelectedTag(*SortByTag)
	if !checkTag {
		panic(ErrTag)
	}

	// checks if PathRecursive is selected
	if *PathRecursive {
		Files := ListFiles(*PathToFiles, *SortByTag, *verbose, extensions)
		DivideFilesToMedium(Files, *PathToFiles, *MediumName, int64(*SideLength), int64(*ReasonableDiff), *SideAmount, *MaxMediums, *verbose, *export)
	} else {
		Files := ListFilesRec(*PathToFiles, *SortByTag, *verbose, extensions)
		for path, files := range Files {
			fmt.Print(path, "\n")
			DivideFilesToMedium(files, path, *MediumName, int64(*SideLength), int64(*ReasonableDiff), *SideAmount, *MaxMediums, *verbose, *export)
			fmt.Println()
		}
	}
}

func SelectExtensions(mp3 bool, flac bool, ogg bool) ([]string, error) {
	extensions := make([]string, 0)
	if mp3 {
		extensions = append(extensions, ".mp3")
	}
	if flac {
		extensions = append(extensions, ".flac")
	}
	if ogg {
		extensions = append(extensions, ".ogg")
	}

	if len(extensions) == 0 {
		return extensions, errors.New("logic error - no extensions are selected")
	}

	return extensions, nil
}

func CheckSelectedTag(tagInput string) (bool, error) {
	UsableTags := [11]string{"album", "albumartist", "artist", "date", "discnumber", "duration", "genre", "publisher", "title", "tracknumber", ""}

	checkTag := false
	for _, tag := range UsableTags {
		if strings.EqualFold(strings.ToLower(tag), strings.ToLower(tagInput)) {
			checkTag = true
		}
	}

	if !checkTag {
		return checkTag, errors.New("spell error - the selected tag are note supported")
	}

	return checkTag, nil
}

func ListFiles(path string, tag string, verbose bool, extensions []string) []fs.FileInfo {
	AudioFiles := make([]fs.FileInfo, 0)
	files, err := ioutil.ReadDir(path)

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
		SortAudioFilesByTag(AudioFiles, path, tag, verbose)
	}

	return AudioFiles
}

func ListFilesRec(path string, tag string, verbose bool, extensions []string) map[string][]fs.FileInfo {
	AudioFolder := make(map[string][]fs.FileInfo)

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && CheckFolderContentOnExtensions(path, extensions) {
			fmt.Println(path)
			AudioFolder[path] = ListFiles(path, tag, verbose, extensions)
		}
		return err
	})
	fmt.Println()

	if err != nil {
		log.Fatal(err)
	}

	return AudioFolder
}

func CheckFolderContentOnExtensions(path string, extensions []string) bool {
	files, err := ioutil.ReadDir(path)
	check := false
	if err != nil {
		log.Fatal(err)
	}

	for _, ext := range extensions {
		for _, file := range files {
			if strings.Contains(strings.ToLower(file.Name()), ext) {
				check = true
			}
		}
	}
	return check
}

func TagCompare(file1 taggolib.Parser, file2 taggolib.Parser, tag string) bool {
	result := true
	switch tag {
	case "album":
		result = file1.Album() > file2.Album()
	case "albumArtist":
		result = file1.AlbumArtist() > file2.AlbumArtist()
	case "artist":
		result = file1.Artist() > file2.Artist()
	case "date":
		result = file1.Date() > file2.Date()
	case "discNumber":
		result = file1.DiscNumber() > file2.DiscNumber()
	case "duration":
		result = file1.Duration() > file2.Duration()
	case "genre":
		result = file1.Genre() > file2.Genre()
	case "publisher":
		result = file1.Publisher() > file2.Publisher()
	case "title":
		result = file1.Title() > file2.Title()
	case "trackNumber":
		result = file1.TrackNumber() > file2.TrackNumber()
	}
	return result
}

func SortAudioFilesByTag(files []fs.FileInfo, path string, tag string, verbose bool) []fs.FileInfo {
	for i := 0; i < len(files)-1; i++ {
		for j := 0; j < len(files)-i-1; j++ {
			file1, _ := os.Open(path + string(os.PathSeparator) + files[j].Name())
			file2, _ := os.Open(path + string(os.PathSeparator) + files[j+1].Name())
			song1, _ := taggolib.New(file1)
			song2, _ := taggolib.New(file2)

			if TagCompare(song1, song2, strings.ToLower(tag)) {
				files[j], files[j+1] = files[j+1], files[j]
			}

			errFile1 := file1.Close()
			errFile2 := file2.Close()

			if errFile1 != nil {
				panic(errFile1)
			}
			if errFile2 != nil {
				panic(errFile2)
			}
		}
	}
	if verbose {
		for _, file := range files {
			p, _ := os.Open(path + string(os.PathSeparator) + file.Name())
			song, _ := taggolib.New(p)
			PrintAudioFiles(file, song, verbose)
		}
	}

	return files
}

func PrintAudioFiles(file fs.FileInfo, song taggolib.Parser, verbose bool) {
	if verbose {
		fmt.Println(song.TrackNumber(), "-", song.Title(), "-", song.Album(), "-", song.Duration())
	} else {
		fmt.Println(file.Name(), song.Duration())
	}
}

func DivideFilesToMedium(files []fs.FileInfo, path string, medium string, duration int64, diff int64, sides uint, maxMed uint, verbose bool, export bool) {
	var SidePlaytime time.Duration // Time
	var SideMaxPlaytime = time.Duration(duration * 60000000000)
	var MediumNumber = 1
	var sideNr = 1
	var FreeSpaceAtMedium time.Duration

	fmt.Println(medium, MediumNumber, "side", sideNr)
	for _, file := range files {
		p, _ := os.Open(path + string(os.PathSeparator) + file.Name())
		song, _ := taggolib.New(p)

		if int(maxMed) >= MediumNumber || maxMed == 0 {
			if (SidePlaytime + song.Duration()) <= (SideMaxPlaytime + time.Duration(diff*1000000000)) {
				SidePlaytime += song.Duration()
				PrintAudioFiles(file, song, verbose)
			} else {
				FreeSpaceAtMedium += SideMaxPlaytime - SidePlaytime
				fmt.Println("free space at", medium, MediumNumber, "side", sideNr, ":", SideMaxPlaytime-SidePlaytime)
				fmt.Println()
				SidePlaytime = 0

				if sideNr == int(sides) {
					sideNr = 1
					MediumNumber++
				} else {
					sideNr++
				}

				if int(maxMed) <= MediumNumber {
					fmt.Println(medium, MediumNumber, "side", sideNr)
					SidePlaytime += song.Duration()
					PrintAudioFiles(file, song, verbose)
				}
			}
		} else {
			return
		}
		_ = p.Close()
	}

	FreeSpaceAtMedium += SideMaxPlaytime - SidePlaytime
	fmt.Print("free space at ", medium, " ", MediumNumber, " side ", sideNr, ": ", SideMaxPlaytime-SidePlaytime, "\n\n")
	fmt.Println("free space of all recorded sides:", FreeSpaceAtMedium)

	if sideNr < int(sides) {
		for sideNr <= int(sides) {
			FreeSpaceAtMedium += SidePlaytime
			sideNr++
		}
	}

	fmt.Println("free space of all recorded medium:", FreeSpaceAtMedium)
}
