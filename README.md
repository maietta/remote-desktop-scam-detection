# Remote Desktop Scam Detection

A Windows-based security application designed to detect and prevent potential remote desktop scam attempts by monitoring browser activity and common remote desktop applications.

## Overview

This application actively monitors your system for:
- Common remote desktop applications (AnyDesk, RustDesk, TeamViewer)
- Browser activity containing suspicious keywords
- Developer tools access (F12 or Ctrl+Shift+I) during suspicious activity

When suspicious activity is detected, the application:
- Automatically terminates detected remote desktop sessions
- Displays a notification warning about potential scam activity
- Provides guidance on next steps

## Features

- Real-time monitoring of browser windows and remote desktop applications
- Automatic process termination for known remote desktop applications
- Keyword-based detection system using customizable wordlists
- Windows notification system integration
- Developer tools access detection
- Lightweight and efficient background operation

## Prerequisites

- Windows operating system
- Go 1.16 or later
- Administrator privileges (for process termination)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/maietta/remote-desktop-scam-detection.git
cd remote-desktop-scam-detection
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o scam-detection.exe
```

## Configuration

The application uses a `wordlist.txt` file to define suspicious keywords. You can customize this file to add or remove keywords based on your needs.

## Usage

1. Run the application with administrator privileges:
```bash
.\scam-detection.exe
```

2. The application will run in the background, monitoring for suspicious activity.

3. When suspicious activity is detected, you'll receive a Windows notification with guidance.

## Wordlist Customization

Edit the `wordlist.txt` file to add or remove keywords. Each keyword should be on a new line and in lowercase.

## Security Considerations

- The application requires administrator privileges to terminate processes
- It only monitors for specific remote desktop applications
- The wordlist can be customized to match your specific needs
- The application runs locally and doesn't send any data externally

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This application is provided as-is and should be used as part of a comprehensive security strategy. It is not a replacement for proper security practices and user awareness. 