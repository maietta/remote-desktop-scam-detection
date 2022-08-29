package main

import (
	"log"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
	"gopkg.in/toast.v1"
)

var (
	mod                     = windows.NewLazyDLL("user32.dll")
	procGetWindowText       = mod.NewProc("GetWindowTextW")
	procGetWindowTextLength = mod.NewProc("GetWindowTextLengthW")
	procGetAsyncKeyState    = mod.NewProc("GetAsyncKeyState")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func getWindow(funcName string) uintptr {
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}

func GetAsyncKeyState(vKey int) int16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return int16(ret)
}

func main() {
	for {
		if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
			text := GetWindowText(HWND(hwnd))

			// fmt.Println("window :", text, "# hwnd:", hwnd)

			if strings.Contains(text, "Brave") || strings.Contains(text, "Google Chrome") || strings.Contains(text, "Microsoft Edge") || strings.Contains(text, "Firefox") {

				// Detect if F12 is pressed
				if GetAsyncKeyState(123) != 0 {

					notification := toast.Notification{
						AppID:   "Scam Detection Monitor",
						Title:   "Possible Scam in Progress!",
						Message: "If you are seeing this after being receiving a call, it appears you may be a victim of a scam. Please hang up and call the company directly to verify the legitimacy of the call.",
						// Icon:    "go.png", // This file must exist (remove this line if it doesn't)
						Actions: []toast.Action{
							{"protocol", "Disconnect Scammer", ""},
							{"protocol", "Learn More", ""},
						},
					}
					err := notification.Push()
					if err != nil {
						log.Fatalln(err)
					}

					time.Sleep(2 * time.Second)

				}
			}

		}
	}
}
