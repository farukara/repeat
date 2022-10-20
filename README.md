# Spaced Repetition for for Taskwarrior

## Overview
## Installation
### Build from source you need [Go](https://go.dev/).
- download this repo by:

    $ git clone https://github.com/farukara/repeat
    cd show note
    make

make runs go build and copies the executable binary "re" to usr/local/bin

### Directly download executable (not available yet)
- download the file from release page 
- save file somewhere in the $PATH
- you can see folders that are on the $PATH by:

    printenv $PATH

on the console.

## Commands
- re 12 1w
schedules task 12 1w later at 6am.

## Usage
-Available spaces times: 1d,2d,3d,5d,1w,2w,3w,1m..11m,1y..5y

## Tips
## Configuration

## Advantages

*It's decoupled from Taskwarrior*. Notes are kept another folder and do not interfere with Taskwarrior at all.
*Extensible*. You can extend the functionality to suit your needs, such as spaced repetition.

## Improvements Needed

- Only testted and used in macOS. It should run alright in Linux. For other platforms it's not been used nor tested.

## Concepts

## TODOs
