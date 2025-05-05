package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/mitchellh/go-ps"
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
	ret, _, _ := procGetWindowTextLength.Call(uintptr(hwnd))
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

// func PS() {
// 	ps, _ := ps.Processes()
// 	fmt.Println(ps[0].Executable())

// 	for pp := range ps {
// 		fmt.Printf("%d %s\n", ps[pp].Pid(), ps[pp].Executable())
// 	}
// }

// FindProcess( key string ) ( int, string, error )
func FindProcess(key string) (int, string, error) {
	pname := ""
	pid := 0
	err := errors.New("not found")
	ps, _ := ps.Processes()

	for i, _ := range ps {
		if ps[i].Executable() == key {
			pid = ps[i].Pid()
			pname = ps[i].Executable()
			err = nil
			break
		}
	}
	return pid, pname, err
}

func killProcessByName(procname string) int {
	kill := exec.Command("taskkill", "/im", procname, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}
	return 0
}

func trigger() {

	usualSuspects := []string{
		"AnyDesk.exe",
		"RustDesk.exe",
		"TeamViewer.exe",
	}

	for _, suspect := range usualSuspects {
		if pid, s, err := FindProcess(suspect); err == nil {
			fmt.Printf("Pid:%d, Pname:%s\n", pid, s)

			// Kill the process
			killProcessByName(suspect)
		}
	}
	// ok := dialog.Message("%s", "Do you want to continue?").Title("Are you sure?").YesNo()
	// if ok {
	// 	// Send notification

	notification := toast.Notification{
		AppID:   "Scam Detection Monitor",
		Title:   "Possible Scam in Progress!",
		Message: "If you are seeing this after being receiving a call, it appears you may be a victim of a scam. The remote settion has been terminated out of an abundance of caution. Please hang up and call the company directly to verify the legitimacy of the call.",
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}

}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func buildListFromFile() []string {
	// Open the wordlist.txt file
	file, err := os.Open("wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines

}

func main() {

	// Build the list of words to look for
	words := buildListFromFile()

	for {
		if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
			text := GetWindowText(HWND(hwnd))

			//browsers := [4]string{"Brave", "Gogle Chrome", "Microsoft Edge", "Mozilla Firefox"}

			// The tab name is anything after the last dash
			tabName := strings.Split(text, " - ")[len(strings.Split(text, " - "))-1]

			// Trim the whitespace
			tabName = strings.TrimSpace(tabName)

			// Does browsers contain the tab name?
			if contains([]string{"Brave", "Gogle Chrome", "Microsoft​ Edge", "Mozilla Firefox"}, tabName) {

				// Remove the browser name from the tab name
				// Remove tabName from text
				data := strings.Replace(text, " - "+tabName, "", -1)
				data = strings.ToLower(data)

				// split the data into words
				dataWords := strings.Split(data, " ")

				// Check if any of the words in the dataWords slice are in the words slice
				for _, word := range dataWords {
					if contains(words, word) {

						// If CTRL + SHIFT + i is pressed.
						if GetAsyncKeyState(0x10) != 0 && GetAsyncKeyState(0x11) != 0 && GetAsyncKeyState(0x49) != 0 {
							trigger()
							time.Sleep(1 * time.Second)

							// If or F12 is pressed.
						} else if GetAsyncKeyState(123) != 0 {
							trigger()
							time.Sleep(1 * time.Second)

						}
						break
					}
				}

			}

		}
	}
}
