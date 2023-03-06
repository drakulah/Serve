package log

import (
	"fmt"
	"strings"
)

func I(message ...string) {
	logTxt := fmt.Sprintf("%s%s %s%s I %s %s", fgMagenta, LogDate(), fgWhite, bgCyan, reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}

func E(message ...string) {
	logTxt := fmt.Sprintf("%s%s %s%s E %s %s", fgMagenta, LogDate(), fgWhite, bgRed, reset, strings.Join(message, " "))
	fmt.Println(logTxt)
}
