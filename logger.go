package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

var WHITE = color.New(color.FgWhite).SprintFunc()
var RED = color.New(color.FgRed).SprintFunc()
var GREEN = color.New(color.FgGreen).SprintFunc()
var YELLOW = color.New(color.FgYellow).SprintFunc()
var BLUE = color.New(color.FgBlue).SprintFunc()

func Log(v ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	str := fmt.Sprintf("%s [LOG] ", currentTime)
	fmt.Printf(WHITE(str) + " ")
	fmt.Println(v...)
}

func LogError(v ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	str := fmt.Sprintf("%s [ERROR] ", currentTime)
	fmt.Printf(RED(str) + " ")
	fmt.Println(v...)
	os.Exit(1)
}

func LogSuccess(v ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	str := fmt.Sprintf("%s [SUCCESS] ", currentTime)
	fmt.Printf(GREEN(str) + " ")
	fmt.Println(v...)
}

func LogWarning(v ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	str := fmt.Sprintf("%s [WARNING] ", currentTime)
	fmt.Printf(YELLOW(str) + " ")
	fmt.Println(v...)
}

func LogInfo(v ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	str := fmt.Sprintf("%s [INFO] ", currentTime)
	fmt.Printf(BLUE(str) + " ")
	fmt.Println(v...)
}
