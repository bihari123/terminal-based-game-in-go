package helper

import (
	"os/exec"
	"runtime"
)

// open opens the specified URL in the default browser of the user.
func openUrl(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start"}
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    args = append(args, url)
  return exec.Command(cmd, args...).Start()
}

/*
One last thing to note: http.ListenAndServe() blocks and never returns (if there is no error). So you have to start the server or the browser in another goroutine, for example:

go open("http://localhost:8080/")
panic(http.ListenAndServe(":8080", nil))

*/ 
