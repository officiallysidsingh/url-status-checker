# URL Status Checker

A simple, elegant command-line tool built with [Bubbletea](https://github.com/charmbracelet/bubbletea) that checks the HTTP status of any URL you enter. This tool provides real-time feedback with a clean terminal user interface.

## Features

- Clean and intuitive terminal user interface
- Real-time URL status checking
- Timeout handling for slow responses
- Error handling for invalid URLs
- Beautiful text input with placeholder
- Loading state feedback

## Installation

To install this application, you'll need Go installed on your system. Then run:

```bash
# Clone the repository
git clone [your-repo-url]
cd [your-repo-name]

# Install dependencies
go mod download

# Build the application
go build -o url-checker
```

## Dependencies

- github.com/charmbracelet/bubbles/textinput
- github.com/charmbracelet/bubbletea

## Usage

1. Run the application:
```bash
./url-checker
```

2. Enter a URL in the text input field
3. Press Enter to check the URL status
4. The application will display either:
   - The HTTP status code and its text representation
   - An error message if something went wrong
   - A loading message while checking

### Controls

- `Enter` - Submit URL for checking
- `Ctrl+C` - Quit the application
- `q` - Quit the application
- `ESC` - Quit the application

## Example Output

```
Enter a URL to check its status:

> https://example.com

Status: 200 OK
```

## Error Handling

The application handles various error cases:
- Empty URL inputs
- Invalid URLs
- Connection timeouts (10-second timeout)
- Network errors
- Server errors
