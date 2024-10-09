# cpz

A small, fast, and modern utility command line tool for `'cp/mv'` alternative or replacement that written in **Go**. It's also allows you to filter files based on filenames, extensions, or regular expressions during the `copy or move process`.

## Features

- Copy or move files or directories while showing a progress bar.
- Exclude files based on their names, extensions, or regular expressions.
- Supports concurrency for faster copying or moving.
- Verify files to guarantee that source and destination file are same after copying or moving.

## Installation

To install **cpz**, clone the repository and run the following commands:

```bash
git clone https://github.com/Kei-K23/go-cpz.git
cd go-cpz
go build -o cpz
```

## Usage

```bash
cpz cp [source] [destination] [flags]
```

### Arguments

- **source**: The path of the file or directory to copy.
- **destination**: The path where the file or directory will be copied.

### Flags

- `-p, --progress`: Show a progress indicator during the copy operation.
- `-f, --filter`: Exclude files by name. Provide multiple values as a comma-separated list.
- `-e, --extensions`: Exclude files by extension. Provide multiple extensions as a comma-separated list (e.g., .log, .tmp).
- `-r, --regex`: Exclude files by regular expression patterns.

### Examples

#### Basic File Copy

```bash
cpz cp /path/to/source/file.txt /path/to/destination/file.txt
```

#### Directory Copy with Progress Bar

```bash
cpz cp /path/to/source/directory /path/to/destination/directory -p
```

#### Exclude Files by Name

```bash
cpz cp /path/to/source /path/to/destination -f "README.md,.env"
```

This will exclude files named README.md and .env from being copied.

#### Exclude Files by Extension

```bash
cpz cp /path/to/source /path/to/destination -e ".log,.tmp"
```

This will exclude files with the .log and .tmp extensions from being copied.

#### Exclude Files by Regular Expression

```bash
cpz cp /path/to/source /path/to/destination -r "^._*backup\.zip$,^old*._"
```

This will exclude files ending with _backup.zip and files starting with old_.
