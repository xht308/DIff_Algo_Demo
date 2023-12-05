package main

import (
	"flag"
	"strings"
	"time"
	"fmt"
)

var (
	// The input files
	srcFile        string
	destFile       string
	// Character mode
	charMode       bool
	// Verbose level
	verboseMode    uint
	// Algorithm
	algorithm      string
	// Processing timeEnable
	timeEnable     bool
)

/////////////////////////////
/** Command Line Interface */
/////////////////////////////

func main() {
	// Parse the flags
	flag.StringVar(&srcFile, "s", "", "The source file (string)")
	flag.StringVar(&destFile, "d", "", "The destination file (string)")
	flag.BoolVar(&charMode, "c", false, "Enable Character mode")
	flag.UintVar(&verboseMode, "v", 1, "Verbose level (0, 1, 2, 3)")
	flag.StringVar(&algorithm, "a", "mayers", "The algorithm to use (basic, mayers, linearspace)")
	flag.BoolVar(&timeEnable, "t", false, "Print the processing time")
	flag.Parse()

	// If the source or destination file is not specified, print the usage
	if srcFile == "" || destFile == "" {
		flag.Usage()
		return
	}

	// Read the content to diff
	var src, dest []string
	// Diff two strings
	if charMode {
		src = strings.Split(srcFile, "")
		dest = strings.Split(destFile, "")
	// Diff two files
	} else {
		src = readFileLines(srcFile)
		dest = readFileLines(destFile)
	}

	// Get the operations
	var operations []operation
	startTime := time.Now()
	var endTime time.Time
	switch algorithm {
	case "basic":
		temp := getBasicDiffOperations(src, dest)
		endTime = time.Now()
		operations = temp.Slice()
	case "mayers":
		temp := getMayersDiffOperations(src, dest)
		endTime = time.Now()
		operations = temp.Slice()
	case "linearspace":
		operations = getLinearMayersDiffOperations(src, dest)
		endTime = time.Now()
	default:
		panic("Invalid algorithm")
	}

	// Print the operations
	switch verboseMode {
	case 0:
	case 1:
		printOperations(operations)
	case 2:
		printOperationsVerbose(operations, src, dest)
	case 3:
		printOperationsFancy(operations, src, dest)
	default:
		panic("Invalid verbose level")
	}

	// Print the processing time
	if timeEnable {
		fmt.Printf("Processing time: %vms\n", endTime.Sub(startTime).Milliseconds())
	}
}