package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"unicode"

	"github.com/dsnet/try"
	"github.com/liblaf/color-make/pkg/logging"
	"github.com/liblaf/color-make/pkg/shutil"
	"github.com/pterm/pterm"
)

var fatalPattern *regexp.Regexp
var errorPattern *regexp.Regexp
var warningPattern *regexp.Regexp
var infoPattern *regexp.Regexp

var errorStyle *pterm.Style
var fatalStyle *pterm.Style
var warningStyle *pterm.Style
var infoStyle *pterm.Style
var execStyle *pterm.Style

func main() {
	defer try.F(logging.Fatalln)
	pipe := bufio.NewReader(os.Stdin)
	for {
		bytes, _, err := pipe.ReadLine()
		if err != nil {
			break
		}
		line := string(bytes)
		writer := os.Stdout
		if fatalPattern.MatchString(line) {
			writer = os.Stderr
			line = fatalStyle.Sprint(line)
		} else if errorPattern.MatchString(line) {
			writer = os.Stderr
			line = errorStyle.Sprint(line)
		} else if warningPattern.MatchString(line) {
			writer = os.Stderr
			line = warningStyle.Sprint(line)
		} else if infoPattern.MatchString(line) {
			writer = os.Stdout
			line = infoStyle.Sprint(line)
		} else if shutil.Has(firstField(line)) {
			writer = os.Stdout
			line = execStyle.Sprint("+ " + line)
		}
		try.E1(fmt.Fprintln(writer, line))
	}
}

func init() {
	defer try.F(logging.Fatalln)
	fatalPattern = try.E1(regexp.Compile(`^make(\[\d+\])?: \*\*\* .*$`))
	errorPattern = try.E1(regexp.Compile(`^.+:\d+: \*\*\* .*\.  Stop\.$`))
	warningPattern = try.E1(regexp.Compile(`^.+:\d+: .*$`))
	infoPattern = try.E1(regexp.Compile(`^make(\[\d+\])?: .*$`))

	fatalStyle = pterm.NewStyle(pterm.Bold, pterm.FgLightRed)
	errorStyle = pterm.NewStyle(pterm.Bold, pterm.FgLightRed)
	warningStyle = pterm.NewStyle(pterm.Bold, pterm.FgLightYellow)
	infoStyle = pterm.NewStyle(pterm.Bold, pterm.FgLightCyan)
	execStyle = pterm.NewStyle(pterm.Bold, pterm.FgLightGreen)
}

func firstField(line string) string {
	for i := 0; i < len(line); i++ {
		if unicode.IsSpace(rune(line[i])) {
			return line[:i]
		}
	}
	return ""
}
