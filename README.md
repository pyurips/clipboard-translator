A simple CLI application that monitors your clipboard and automatically translates any new copied text using the DeepL API.

## How it works
- The program runs in the terminal and checks every second for changes in the clipboard text.
- If new text is detected, it is sent for translation (from EN to PT-BR by default) using the DeepL API.
- The translation result is displayed in the terminal.

## Requirements
- Go 1.24 or higher
- DeepL API account and key (https://www.deepl.com/pro-api)

## Installation and Usage
1. Clone the repository and enter the project folder.
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Create a `.env` file in the project root with the following content:
   ```env
   DEEPL_API_KEY=your_deepl_api_key_here
   ```
   (You can use the `.env.example` file as a template.)
4. Build the project:
   ```sh
   go build -o app.exe
   ```
5. Run the program:
   ```sh
   ./app.exe
   ```

## Notes
- The translated text is limited to 100 characters per request (see `MAX_CHAR_LIMIT` in `schemas.go`).
- Source and target languages can be changed in `schemas.go`.
- The program uses the [atotto/clipboard](https://github.com/atotto/clipboard) library to access the clipboard.

## License
MIT License
