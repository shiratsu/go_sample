package main

import (
    "log"
    "fmt"
    "os/exec"
)

func main() {
    log.Print("これはコマンドを実行するプログラムです!\n")
    command := exec.Command("sleep", "2")
    error := command.Start()

    if error != nil {
        panic(fmt.Sprintf("エラー1: %v", error)) // コマンドがなかった場合とか
    }

    log.Print("コマンドを開始しました")

    error = command.Wait()

    if error != nil {
        panic(fmt.Sprintf("エラー2: %v", error)) // コマンドがエラーステータスで終了した場合
    }

    log.Print("コマンドが終了しました")
}