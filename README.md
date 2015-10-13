##	WordList Cleaner v.0.2.6

Remove non-printable words, trim words length, search duplicates, sorting, words counting.

![Alt text](/screenshot.jpg?raw=true "Usage")

Sequence of the option keys is not critical.

05.09.2015 - Firts commit.

06.09.2015 - Added automatic processing of all files in a directory by extension.

07.09.2015 - Added lines calculator. Fixed some errors.

12.09.2015 - All algorithms has been rewritten to consume less memory. Usage menu and option keys has been changed too.

13.09.2015 - Algorithm for finding duplicates has been optimized.

13.10.2015 - Changed algorithm and view for progress bar. Fixed calculation for strings in files.

Examples:
```
wordlistcleaner.exe -min 8 -max 10 -src Source.dic -new New.dic remove trim
wordlistcleaner.exe -src Source.dic -new New.dic trim
wordlistcleaner.exe -a -ext txt duplicate
wordlistcleaner.exe -a sort
wordlistcleaner.exe -a calculate
```