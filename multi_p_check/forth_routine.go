package main

import (
    "log"
    "time"
)

func main() {
    log.Print("started.")

    // チャネル
    finished := make(chan bool)

    go func() {
        // 1秒かかるコマンド
        log.Print("sleep1 started.")
        time.Sleep(1 * time.Second)
        log.Print("sleep1 finished.")
        finished <- true
    }()

    go func() {
        // 2秒かかるコマンド
        log.Print("sleep2 started.")
        time.Sleep(2 * time.Second)
        log.Print("sleep2 finished.")
        finished <- true
    }()

    go func() {
        // 3秒かかるコマンド
        log.Print("sleep3 started.")
        time.Sleep(3 * time.Second)
        log.Print("sleep3 finished.")
        finished <- true
    }()

    // 終わるまで待つ
    for i := 1; i <= 3; i++ {
        <-finished
    }

    log.Print("all finished.")
}
