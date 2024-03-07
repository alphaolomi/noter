# Noter

A simple note taking app built with Go.

## Installation

```bash
curl -sL $GH_REPO/releases/latest/download/noter_linux_amd64.tar.gz | tar xz
```

## Usage

```bash
$ ./noter.exe 
Usage:
  notes [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create a new note
  remove      Remove a note
  update      Update a note
  help        Help about any command
  read        Read all notes

Flags:
  -h, --help   help for notes

Use "notes [command] --help" for more information about a command.  
```


<small>Built with <a href="https://cobra.dev/">Cobra</a></small>