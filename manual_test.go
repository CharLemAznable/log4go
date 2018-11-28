package log4go

import (
    "bufio"
    "fmt"
    "github.com/CharLemAznable/log4go"
    "io"
    "os"
    "testing"
    "time"
)

func TestConsoleLogWriter_Manual(t *testing.T) {
    log := log4go.NewLogger()
    log.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())
    log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
    time.Sleep(time.Second)
}

func TestFileLogWriter_Manual(t *testing.T) {
    const (
        filename = "flw.log"
    )
    // Get a new logger instance
    log := log4go.NewLogger()

    // Create a default logger that is logging messages of FINE or higher
    log.AddFilter("file", log4go.FINE, log4go.NewFileLogWriter(filename, false))
    log.Close()

    /* Can also specify manually via the following: (these are the defaults) */
    flw := log4go.NewFileLogWriter(filename, false)
    flw.SetFormat("[%D %T] [%L] (%S) %M")
    flw.SetRotate(false)
    flw.SetRotateSize(0)
    flw.SetRotateLines(0)
    flw.SetRotateDaily(false)
    log.AddFilter("file", log4go.FINE, flw)

    // Log some experimental messages
    log.Finest("Everything is created now (notice that I will not be printing to the file)")
    log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
    log.Critical("Time to close out!")

    // Close the log
    log.Close()

    time.Sleep(time.Second * 1)

    // Print what was logged to the file (yes, I know I'm skipping error checking)
    fd, _ := os.Open(filename)
    in := bufio.NewReader(fd)
    fmt.Print("Messages logged to file were: (line numbers not included)\n")
    for lineno := 1; ; lineno++ {
        line, err := in.ReadString('\n')
        if err == io.EOF {
            break
        }
        fmt.Printf("%3d:\t%s", lineno, line)
    }
    fd.Close()

    // Remove the file so it's not lying around
    os.Remove(filename)
}

func TestXMLConfigurationExample(t *testing.T) {
    // Load the configuration (isn't this easy?)
    log4go.LoadConfiguration("manual_test_example.xml")

    // And now we're ready!
    log4go.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
    log4go.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
    log4go.Info("About that time, eh chaps?")
    log4go.Close()

    time.Sleep(time.Second * 10)
}
