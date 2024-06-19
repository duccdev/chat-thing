# Building

## Prerequisites

- Go
- Bash

## Process

Run `./build.sh <os> <arch>`  
`os` argument has to be a supported `GOOS` argument  
`arch` argument has to be a supported `GOARCH` argument  
If all goes well, you should have a binary in the build folder

# Development

## Formatting

- Use 4-width hard tabs in your code. You're not a p###y.
- Put your first curly brace on the same function declaration line. You're not Microsoft.
- Make your variables clear and concise but don't make them insanely long or short.
- Use lowerCamelCase for variables and functions (unless public, then use Upper)
- Use snake_case for API response values
- Use early returns, try not to use `else` unless absolutely necessary

## Running it

Use the `dev.sh` script (it accepts arguments)
