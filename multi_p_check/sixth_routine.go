package main

import (
    "log"
    "sync"
    "time"
)

func main() {
    log.Print("started.")

    // 配列
    funcs := []func(){
        func() {
            // 1秒かかるコマンド
            log.Print("sleep1 started.")
            time.Sleep(1 * time.Second)
            log.Print("sleep1 finished.")
        },
        func() {
            // 2秒かかるコマンド
            log.Print("sleep2 started.")
            time.Sleep(2 * time.Second)
            log.Print("sleep2 finished.")
        },
        func() {
            // 3秒かかるコマンド
            log.Print("sleep3 started.")
            time.Sleep(3 * time.Second)
            log.Print("sleep3 finished.")
        },
    }

    var waitGroup sync.WaitGroup

    // 関数の数だけ並列化する
    for _, sleep := range funcs {
        waitGroup.Add(1) // 待つ数をインクリメント

        // Goルーチンに入る
        go func(function func()) {
            defer waitGroup.Done() // 待つ数をデクリメント
            function()
        }(sleep)

    }

    waitGroup.Wait() // 待つ数がゼロになるまで処理をブロックする

    log.Print("all finished.")
}