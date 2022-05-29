# Divide and Record
This CLI Tool is made for people how want to recorde a huge amount of audio files to older mediums like cassettes or magnetic tape.
The supported data types are:
- .mp3
- .flac
- .ogg

### ToDo
- [x] Sort by filename
- [x] Set maximum medium for Record
- [ ] List all files in subfolders
- [ ] Sort files by different audio tags
- [ ] Export to txt file

## Compile
```
go.exe build -o C:\Path\DivideAndRecord.exe DivideAndRecord
```
___
## Examples
### Default
The only necessary argument ist ``-path``
```
DivideAndRecord.exe -Path="C:\path_to_audio_files"
```

### 90 Min cassette
Cassette 90 minutes with 2 sides with each 45 minutes
```
DivideAndRecord.exe -SideLength=45 -SideAmount=2 -MediumName=Casette -Path="C:\path_to_audio_files"
```

## List of all Arguments
### Audio Medium Settings
```
  -SideLength int
        Number of minutes playing time per page of the medium (default 30)
  -MediumName string
        Name of the medium like cassette or CD (default "Medium")
  -SideAmount int
        The amount of sides of the medium (default 2)
  -ReasonableDiff int
        Reasonable difference between the maximum playing time of the medium and the end of a track in seconds
  -MaxMediums int
        In case the number of media to be recorded is limited
```

### Path & Files Settings
```
  -Path string
        Location of the files to be recorded
  -Recrusiv
        Lists all files in subfolders
  -MP3
        Should MP3-files be indexed (default true)
  -FLAC
        Should FLAC-files be indexed (default false)
  -OGG
        Should OGG-files be indexed (default false)
```

### Different Settings
```
  -h
        help
```
___

## Used librarys
- taggolib
  - This library is used to read out the audio tags
    - ``go get github.com/mdlayher/taggolib``
- bit
  - Not used directly but needed by taggolib
    - ``go get github.com/eaburns/bit``