# Remote Desktop Scam Detection

A Windows-based security application designed to detect and prevent potential remote desktop scam attempts by monitoring browser activity and common remote desktop applications.

## Project Status and Future Potential

This project serves as a proof of concept demonstrating how remote desktop security can be enhanced to protect users from potential scams. The core functionality could be integrated directly into popular remote desktop solutions such as:
- RustDesk
- TeamViewer
- AnyDesk
- Other remote desktop applications

While currently Windows-focused, the underlying concepts and detection mechanisms could be adapted for:
- macOS (using similar system API calls)
- Linux (using X11/Wayland window management APIs)
- Other operating systems

Such integration would provide native protection against common remote access scam techniques across all major platforms.

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

- It only monitors for specific remote desktop applications
- The wordlist can be customized to match your specific needs
- The application runs locally and doesn't send any data externally

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This application is provided as-is and should be used as part of a comprehensive security strategy. It is not a replacement for proper security practices and user awareness. 