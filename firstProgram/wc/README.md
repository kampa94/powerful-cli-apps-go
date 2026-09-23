# gowc

A minimal [wc](https://man7.org/linux/man-pages/man1/wc.1.html)-style word count tool written in Go. Reads text from standard input and counts lines, words, characters, or bytes — just like the classic Unix tool.

## Usage

Pipe text into `gowc` and pick a mode with a flag:

```sh
cat file.txt | gowc            # count words (default)
cat file.txt | gowc -l         # count lines
cat file.txt | gowc -b         # count bytes
cat file.txt | gowc -c         # count characters (runes)
echo "Hello World" | gowc      # 2
```

### Flags

| Flag | Description                  |
|------|------------------------------|
| `-l` | Count lines                 |
| `-b` | Count bytes                 |
| `-c` | Count characters (runes)    |
| *(none)* | Count words (default)   |

## Build

```sh
go build -o gowc .
```

## Test

```sh
go test ./...
```

## How it works

`gowc` uses a `bufio.Scanner` and picks a split function based on the chosen flag:

- `-l` → `bufio.ScanLines`
- `-c` → `bufio.ScanRunes`
- `-b` → `bufio.ScanBytes`
- default → `bufio.ScanWords`

It then counts the number of tokens produced by the scanner. The counting logic lives in the testable pure function `count()` in `main.go`.