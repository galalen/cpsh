# cpsh

Copy without leaving your terminal 💻

## Features
- Copy text directly: `cpsh "hello world"`
- Copy file contents: `cpsh -f file.txt`
- Copy file paths for non-text files: `cpsh -f document.pdf`
- Copy from pipe: `echo hello | cpsh`

## Installation
```bash
go install github.com/galalen/cpsh@latest
```

## Building from source
```bash
git clone https://github.com/galalen/cpsh
cd cpsh
go build
```
