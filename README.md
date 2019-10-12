# Description
eol is using for changing text's the end of line.

usage
```Shell
$ eol
A line separators changer for text, it supports crlf(\r\n), lf(\n) and cr(\r).

Usage:
  eol [command]

Available Commands:
  cr          Convert all end of line to cr(\r)
  crlf        Convert all end of line to crlf(\r\n)
  encoding    Get all supported encoding
  help        Help about any command
  lf          Convert all end of line to lf(\n)

Flags:
  -h, --help      help for eol
      --version   version for eol

Use "eol [command] --help" for more information about a command.
```

## Supported encoding
Unicode: utf-8, utf-16b(big-endian), utf-16l(little-endian)  
CKJ: gbk, gb18030, big-5, shift-j, euckr

## Auto skip binrary file
`eol` will auto detect binrary files and skip them via finding unicode control char in file.

## Sample

01. Change all files in current work directory's eol to lf(\n), eol will try all encoding for those files.
```
$ eol lf .
```

02. Change all files in 'dir' and file 'text01.txt' 'text02.txt' to crlf(\r\n), read those file as utf-8 encoding.
```Shell
$ eol crlf -e utf8 ./dir text01.txt text02.txt
```
