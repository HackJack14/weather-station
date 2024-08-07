package db

import (
    "log"
	"os/exec"
	"runtime"
)

func Execute(statement string) {
    terminal := "";
    if runtime.GOOS == "windows" {
        terminal = "powershell"
    } else {
        terminal = "bash"
    }
	cmd := exec.Command(terminal, "./sqlite3", "weatherstation.db", "\"" + statement + "\"", "> result.txt") 
	log.Println(cmd.String())
	cmd.Start()
}