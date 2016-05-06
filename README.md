##	WordList Cleaner v0.2.6

Remove non-printable words, trim words length, search duplicates, sorting, words counting.

![Alt text](/screenshot.jpg?raw=true "Usage")

Sequence of the option keys is not critical.

05.09.2015 - Firts commit.

06.09.2015 - Added automatic processing of all files in a directory by extension.

07.09.2015 - Added lines calculator. Fixed some errors.

12.09.2015 - All algorithms has been rewritten to consume less memory. Usage menu and option keys has been changed too.

13.09.2015 - Algorithm for finding duplicates has been optimized.

13.10.2015 - Changed algorithm and view for progress bar. Fixed calculation for strings in files.

### Install:
```
go get github.com/gkiryaziev/go-wordlist-cleaner
```

### Build and Run
```
go build && go-wordlist-cleaner
```

### Usage:
```
go-wordlist-cleaner -min 8 -max 10 -src Source.dic -new New.dic remove trim
go-wordlist-cleaner -src Source.dic -new New.dic trim
go-wordlist-cleaner -a -ext txt duplicate
go-wordlist-cleaner -a sort
go-wordlist-cleaner -a calculate
```