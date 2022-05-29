# Divide and Record
This CLI Tool is made for people how want to recorde a huge amount of audio files to older mediums like cassettes or magnetic tape.
The supported data types are:
- .mp3
- .flac
- .ogg

### ToDo
- [x] Sort by filename
- [x] Set maximum medium for Record
- [x] implement verbose argument
- [ ] List all files in subfolders (work in progress)
  - [ ] Create one report for all folders
  - [ ] Create for every folder his own report
- [ ] Sort files by different audio tags (work in progress)
- [ ] Export to txt file (work in progress)

## Compile
```
go.exe build -o C:\Path\DivideAndRecord.exe DivideAndRecord
```
## Examples
### Default
The only necessary argument ist ``-path``
```
.\DivideAndRecord.exe -Path="C:\path_to_audio_files"
```

### 90 Min cassette
Cassette 90 minutes with 2 sides with each 45 minutes
```
.\DivideAndRecord.exe -SideLength=45 -SideAmount=2 -MediumName=cassette -ReasonableDiff=30 -verbose=false -Path="C:\path_to_audio_files"
```

Output:
```
cassette : 1 Side : 1
001 - Kapitel 1.mp3 3m2s
002 - Kapitel 2.mp3 3m22s
003 - Kapitel 3.mp3 3m5s
004 - Kapitel 4.mp3 3m14s
005 - Kapitel 5.mp3 3m25s
006 - Kapitel 6.mp3 3m21s
007 - Kapitel 7.mp3 3m16s
008 - Kapitel 8.mp3 3m35s
009 - Kapitel 9.mp3 3m40s
010 - Kapitel 10.mp3 3m44s
011 - Kapitel 11.mp3 3m22s
012 - Kapitel 12.mp3 2m30s
013 - Kapitel 13.mp3 3m6s
free space at cassette 1 side 1 : 2m18s

015 - Kapitel 15.mp3 3m4s
016 - Kapitel 16.mp3 3m22s
017 - Kapitel 17.mp3 3m15s
018 - Kapitel 18.mp3 3m21s
019 - Kapitel 19.mp3 3m21s
020 - Kapitel 20.mp3 3m25s
021 - Kapitel 21.mp3 3m36s
022 - Kapitel 22.mp3 3m9s
023 - Kapitel 23.mp3 3m52s
024 - Kapitel 24.mp3 3m27s
025 - Kapitel 25.mp3 3m27s
026 - Kapitel 26.mp3 3m29s
027 - Kapitel 27.mp3 3m8s
free space at cassette 1 side 2 : 1m4s

029 - Kapitel 29.mp3 3m23s
030 - Kapitel 30.mp3 3m30s
031 - Kapitel 31.mp3 3m9s
032 - Kapitel 32.mp3 3m36s
033 - Kapitel 33.mp3 3m27s
034 - Kapitel 34.mp3 3m13s
035 - Kapitel 35.mp3 3m23s
036 - Kapitel 36.mp3 3m31s
037 - Kapitel 37.mp3 3m45s
038 - Kapitel 38.mp3 3m36s
039 - Kapitel 39.mp3 3m36s
040 - Kapitel 40.mp3 3m37s
041 - Kapitel 41.mp3 3m30s
free space at cassette 2 side 1 : -16s

043 - Kapitel 43.mp3 3m42s
044 - Kapitel 44.mp3 3m32s
045 - Kapitel 45.mp3 3m21s
046 - Kapitel 46.mp3 3m26s
047 - Kapitel 47.mp3 3m18s
048 - Kapitel 48.mp3 3m38s
049 - Kapitel 49.mp3 3m34s
050 - Kapitel 50.mp3 3m20s
051 - Kapitel 51.mp3 3m15s
052 - Kapitel 52.mp3 3m14s
053 - Kapitel 53.mp3 2m59s
054 - Kapitel 54.mp3 3m18s
055 - Kapitel 55.mp3 3m53s
free space at cassette 2 side 2 : 30s

057 - Kapitel 57.mp3 3m33s
058 - Kapitel 58.mp3 3m52s
059 - Kapitel 59.mp3 3m17s
060 - Kapitel 60.mp3 3m25s
061 - Kapitel 61.mp3 3m32s
062 - Kapitel 62.mp3 3m28s
063 - Kapitel 63.mp3 3m39s
064 - Kapitel 64.mp3 3m25s
065 - Kapitel 65.mp3 3m39s
066 - Kapitel 66.mp3 3m16s
067 - Kapitel 67.mp3 3m32s
068 - Kapitel 68.mp3 2m52s
069 - Kapitel 69.mp3 3m15s
free space at cassette 3 side 1 : 15s

071 - Kapitel 71.mp3 3m23s
072 - Kapitel 72.mp3 3m34s
073 - Kapitel 73.mp3 3m6s
074 - Kapitel 74.mp3 3m27s
075 - Kapitel 75.mp3 3m17s
076 - Kapitel 76.mp3 3m27s
077 - Kapitel 77.mp3 3m25s
078 - Kapitel 78.mp3 3m15s
079 - Kapitel 79.mp3 3m11s
080 - Kapitel 80.mp3 2m49s
081 - Kapitel 81.mp3 2m38s
082 - Kapitel 82.mp3 3m22s
083 - Kapitel 83.mp3 3m30s
free space at cassette 3 side 2 : 2m36s

085 - Kapitel 85.mp3 3m16s
086 - Kapitel 86.mp3 3m34s
087 - Kapitel 87.mp3 3m27s
088 - Kapitel 88.mp3 3m21s
089 - Kapitel 89.mp3 3m11s
090 - Kapitel 90.mp3 3m37s
091 - Kapitel 91.mp3 3m41s
092 - Kapitel 92.mp3 3m48s
093 - Kapitel 93.mp3 3m10s
094 - Kapitel 94.mp3 3m18s
095 - Kapitel 95.mp3 3m30s
096 - Kapitel 96.mp3 3m11s
097 - Kapitel 97.mp3 3m17s
free space at cassette 4 side 1 : 39s

099 - Kapitel 99.mp3 3m15s
100 - Kapitel 100.mp3 3m13s
101 - Kapitel 101.mp3 3m26s
102 - Kapitel 102.mp3 3m19s
103 - Kapitel 103.mp3 3m15s
104 - Kapitel 104.mp3 3m24s
105 - Kapitel 105.mp3 3m11s
106 - Kapitel 106.mp3 3m23s
107 - Kapitel 107.mp3 3m29s
108 - Kapitel 108.mp3 3m10s
109 - Kapitel 109.mp3 3m21s
110 - Kapitel 110.mp3 3m26s
111 - Kapitel 111.mp3 3m38s
free space at cassette 4 side 2 : 1m30s

113 - Kapitel 113.mp3 3m17s
114 - Kapitel 114.mp3 3m32s
115 - Kapitel 115.mp3 3m22s
116 - Kapitel 116.mp3 3m15s
117 - Kapitel 117.mp3 3m26s
118 - Kapitel 118.mp3 3m39s
119 - Kapitel 119.mp3 3m47s
120 - Kapitel 120.mp3 3m44s
121 - Kapitel 121.mp3 3m18s
122 - Kapitel 122.mp3 3m21s
123 - Kapitel 123.mp3 3m17s
124 - Kapitel 124.mp3 3m7s
125 - Kapitel 125.mp3 2m58s
free space at cassette 5 side 1 : 57s

127 - Kapitel 127.mp3 3m20s
128 - Kapitel 128.mp3 3m42s
129 - Kapitel 129.mp3 3m27s
130 - Kapitel 130.mp3 3m7s
131 - Kapitel 131.mp3 3m18s
132 - Kapitel 132.mp3 3m15s
133 - Kapitel 133.mp3 3m39s
134 - Kapitel 134.mp3 3m15s
135 - Kapitel 135.mp3 3m28s
136 - Kapitel 136.mp3 3m21s
137 - Kapitel 137.mp3 3m29s
138 - Kapitel 138.mp3 3m33s
139 - Kapitel 139.mp3 3m54s
free space at cassette 5 side 2 : 12s

141 - Kapitel 141.mp3 3m29s
142 - Kapitel 142.mp3 3m11s
143 - Kapitel 143.mp3 3m6s
144 - Kapitel 144.mp3 3m33s
145 - Kapitel 145.mp3 3m13s
146 - Kapitel 146.mp3 3m25s
147 - Kapitel 147.mp3 3m49s
148 - Kapitel 148.mp3 3m30s
149 - Kapitel 149.mp3 3m18s
150 - Kapitel 150.mp3 3m20s
151 - Kapitel 151.mp3 3m6s
152 - Kapitel 152.mp3 3m3s
153 - Kapitel 153.mp3 3m22s
free space at cassette 6 side 1 : 1m35s

155 - Kapitel 155.mp3 3m23s
156 - Kapitel 156.mp3 3m24s
157 - Kapitel 157.mp3 3m39s
158 - Kapitel 158.mp3 3m31s
159 - Kapitel 159.mp3 3m31s
160 - Kapitel 160.mp3 3m38s
161 - Kapitel 161.mp3 3m18s
162 - Kapitel 162.mp3 3m22s
163 - Kapitel 163.mp3 3m43s
164 - Kapitel 164.mp3 3m25s
165 - Kapitel 165.mp3 3m29s
166 - Kapitel 166.mp3 3m22s
167 - Kapitel 167.mp3 3m31s
free space at cassette 6 side 2 : -16s

169 - Kapitel 169.mp3 2m59s
170 - Kapitel 170.mp3 3m11s
171 - Kapitel 171.mp3 3m7s
172 - Kapitel 172.mp3 3m25s
173 - Kapitel 173.mp3 3m16s
174 - Kapitel 174.mp3 3m38s
175 - Kapitel 175.mp3 3m43s
176 - Kapitel 176.mp3 3m29s
177 - Kapitel 177.mp3 3m27s
178 - Kapitel 178.mp3 3m29s
179 - Kapitel 179.mp3 3m30s
180 - Kapitel 180.mp3 3m30s
181 - Kapitel 181.mp3 3m12s
free space at cassette 7 side 1 : 1m4s

183 - Kapitel 183.mp3 3m30s
184 - Kapitel 184.mp3 3m21s
185 - Kapitel 185.mp3 3m20s
186 - Kapitel 186.mp3 3m51s
187 - Kapitel 187.mp3 3m52s
188 - Kapitel 188.mp3 3m23s
189 - Kapitel 189.mp3 3m36s
190 - Kapitel 190.mp3 3m23s
191 - Kapitel 191.mp3 3m20s
192 - Kapitel 192.mp3 3m51s
193 - Kapitel 193.mp3 3m34s
194 - Kapitel 194.mp3 3m33s
free space at cassette 7 side 2 : 2m26s

196 - Kapitel 196.mp3 3m12s
197 - Kapitel 197.mp3 3m8s
198 - Kapitel 198.mp3 2m30s
199 - Kapitel 199.mp3 3m5s
200 - Kapitel 200.mp3 3m52s
201 - Kapitel 201.mp3 3m39s
free space at cassette 8 side 1 : 25m34s

free space of all recorded sides: 40m8s
free space of all recorded medium: 1h19m0s
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
  -verbose
        Prints more information about audio files
  -h
        help
```

## Used librarys
- taggolib
  - This library is used to read out the audio tags
    - ``go get github.com/mdlayher/taggolib``
- bit
  - Not used directly but needed by taggolib
    - ``go get github.com/eaburns/bit``