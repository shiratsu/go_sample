package main

import (
    "log"
    "os/exec"
)

func main() {
    log.Print("これはコマンドを実行するプログラムです!\n")
    out, error := exec.Command("date").Output()

    if error != nil {
        log.Printf("エラー: %s", error)
        return
    }

    log.Printf("現在時刻: %s", out)
}