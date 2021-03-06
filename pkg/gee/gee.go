package gee

import (
	"bufio"
	"os"
	"strconv"

	// "sync"

	model "github.com/hahwul/gee/pkg/model"
	printing "github.com/hahwul/gee/pkg/printing"
)

// Gee is running gee
func Gee(options model.Options) {
	printing.DebugMsg("MSG", "Starting gee", options.Debug)
	sc := bufio.NewScanner(os.Stdin)
	mode := os.O_CREATE | os.O_WRONLY
	var files = []*os.File{}
	stdLine := 1
	stdPointer := 1
	distributePointer := 0

	if options.Append {
		mode = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	}

	for _, filename := range options.Files {
		f, err := os.OpenFile(filename, mode, 0644)
		if err != nil && options.Debug {
			printing.DebugMsg("ERROR", err, options.Debug)
		} else {
			files = append(files, f)
		}
	}
	//	var wg sync.WaitGroup
	//	wg.Add(1 + len(options.Files))
	var uniqTable []string
	for sc.Scan() {
		l := sc.Text()
		runChecker := false
		if options.Uniq {
			runChecker, uniqTable = Uniq(l, uniqTable)
			printing.DebugMsg("runChecker", runChecker, options.Debug)
			printing.DebugMsg("uniqTable", uniqTable, options.Debug)
		}
		if !runChecker {
			regexBool := true
			if options.Regex != "" {
				regexBool = Regex(options.Regex, l)
			}
			if options.RegexV != "" {
				regexBool = RegexV(options.RegexV, l)
			}
			if regexBool {
				linec, line := StringProc(l, stdLine, options)

				// Print to Stdout
				StdPrint(linec, options)

				// Write to files
				if (stdLine > options.ChunkedLine) && (options.ChunkedLine != 0) {
					ClosedFiles(files)
					for _, filename := range options.Files {
						f, err := os.OpenFile(filename+"_"+strconv.Itoa(stdPointer), mode, 0644)
						if err != nil && options.Debug {
							printing.DebugMsg("ERROR", err, options.Debug)
						} else {
							files = append(files, f)
						}
					}
					stdLine = 1
					stdPointer = stdPointer + 1
				}
				if options.Distribute && (len(files) > 0) {
					if distributePointer < len(files) {
						WriteFile(files[distributePointer], line, options)
						distributePointer = distributePointer + 1
					} else {
						distributePointer = 0
						WriteFile(files[distributePointer], line, options)
						distributePointer = distributePointer + 1
					}
				} else {
					for _, k := range files {
						WriteFile(k, line, options)
					}
				}
				stdLine = stdLine + 1
			}
		}
	}
	//	wg.Wait()

	// Graceful shutdown
	ClosedFiles(files)
	printing.DebugMsg("MSG", "Finish gee, graceful shutdown...", options.Debug)
}
