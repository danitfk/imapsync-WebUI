package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "strings"

)


func hello(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    switch r.Method {
    case "GET":
         http.ServeFile(w, r, "form.html")
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
	fmt.Fprintf(w, "Process should began, Please wait for 10 minutes")
        name := r.FormValue("email")
        oldpass := r.FormValue("oldpass")
	newpass := r.FormValue("newpass")
	host1	:= r.FormValue("host1")
	host2	:= r.FormValue("host2")
	cmdName := ("sudo /usr/bin/imapsync --host1 " + host1 + "--user1 " + name + " --password1 " + oldpass + " --host2 " + host2 + " --user2 " + name + " --password2 " + newpass)
	cmdArgs := strings.Fields(cmdName)
	cmd	:= exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	cmd.Start()

    default:
        fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func main() {
    http.HandleFunc("/", hello)
    fmt.Printf("Starting server for testing HTTP POST...\n")

    if err := http.ListenAndServe(":666", nil); err != nil {
        log.Fatal(err)
    }
}


