package main

import (
	"os"
	"os/exec"
	"strings"
	"bufio"
	"log"
)

func sanitizeInput(in string) []string {
	var out []string
	for o := 0; o < len(in); o++ {
		x := string(in[o])
		switch x {
		case " ": out = append(out, "space")
		case "!": out = append(out, "exclam")
		case "#": out = append(out, "numbersign")
		case "$": out = append(out, "dollar")
		case "%": out = append(out, "percent")
		case "&": out = append(out, "ampersand")
		case "'": out = append(out, "apostrophe")
		case "(": out = append(out, "parenleft")
		case ")": out = append(out, "parenright")
		case "*": out = append(out, "asterisk")
		case "+": out = append(out, "plus")
		case ",": out = append(out, "comma")
		case "-": out = append(out, "minus")
		case ".": out = append(out, "period")
		case "/": out = append(out, "slash")
		case ":": out = append(out, "colon")
		case ";": out = append(out, "semicolon")
		case "<": out = append(out, "less")
		case "=": out = append(out, "equal")
		case ">": out = append(out, "greater")
		case "?": out = append(out, "question")
		case "@": out = append(out, "at")
		case "[": out = append(out, "bracketleft")
		case "\"": out = append(out, "quotedbl")
		case "\\": out = append(out, "backslash")
		case "\n": out = append(out, "Return")
		case "\r": out = append(out, "Return")
		case "\t": out = append(out, "Tab")
		case "]": out = append(out, "bracketright")
		case "^": out = append(out, "asciicircum")
		case "_": out = append(out, "underscore")
		case "`": out = append(out, "grave")
		case "{": out = append(out, "braceleft")
		case "|": out = append(out, "bar")
		case "}": out = append(out, "braceright")
		case "~": out = append(out, "asciitilde")
		default:
			out = append(out, x)
		}
	}
	return out
}

// Command wrapper for xdotool
func xdoCommand(c string, additional string) string {
	var out *exec.Cmd
	if (additional != "") {
		out = exec.Command("xdotool", c, additional)
	} else {
		out = exec.Command("xdotool", c)
	}
	stdout, err := out.Output()
	if err != nil {
		log.Fatal(c + ": " + err.Error())
	}
	return string(stdout)
}

func parseStdin() (string, string) {
	// https://stackoverflow.com/questions/49704456/how-to-read-from-device-when-stdin-is-pipe
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return string(stdin), "selectwindow"
	} else if (len(os.Args) > 1) {
		return strings.Join(os.Args[1:], " "), "selectwindow"
	} else {
		clipboard := exec.Command("xclip", "-o")
		clipout, err := clipboard.Output()
		if err != nil {
			log.Fatal("Clipboard error: " + err.Error())
		}
		return string(clipout), "getactivewindow"
	}
}

func main() {
	input, target := parseStdin()

	targetWindow := xdoCommand(target, "")
	xdoCommand("windowactivate", targetWindow)
	xdoCommand("windowfocus", targetWindow)

	in := sanitizeInput(input)

	for i := 0; i < len(in); i++ {
		xdoCommand("key", in[i])
	}
}
