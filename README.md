# barrel up

This Go program is designed for to automate the creation of index files that export from subdirectories within the `src` folder.

## Features

- Walks through the `src` directory and its subdirectories
- Creates or updates `index.ts` files in each subdirectory
- Only exports files that actually contain export statements
- Avoids creating duplicate exports
- Supports `.tsx`, `.jsx`, `.ts`, and `.js` files

## Install/Usage

- Download the binary from releases and run it `barrel-up` or get the source and compile it `go build`/run it `go run main.go`

1. Navigate to the project's root directory in the terminal
2. Run the program:
   ```
   go run path/to/indexer.go
   ```
   or
   ```
   barrel-up
   ```
## Notes

- It assumes it's run from the project root directory
- It only processes subdirectories within `src`, not the top-level `src` directory
- Existing `index.ts` files will be overwritten
- The program doesn't handle commented-out exports

## Customization

Because I use TS for most projects, some adjustments could be needed if the specific project structure differs. Key areas to consider:
- File extensions (in the `createIndexFile` function)
- Export statement format (in the `hasExports` function) 
- Index file naming (currently set to `index.ts`)
