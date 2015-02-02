package main

import "fmt"
import "time"

func myRoutine(mess string) {
    fmt.Printf(
        "ルーチンが%sを受け取ったお（＾ω＾ ≡ ＾ω＾）\n", mess,
    )
}

func main() {
    // goステートメントに関数を渡すだけで
    // goroutineがつくられる
    go myRoutine("hoge")
    go myRoutine("fuga")
 
    // goroutineはmainが死ぬと死ぬので
    // とりあえずmainを延命するよ
    time.Sleep(1 * time.Second)
    fmt.Println("Main ＼(^o^)／")
}