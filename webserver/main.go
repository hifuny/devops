/*
* @Author: Hifun
* @Date: 2020/1/3 11:48
 */
package main

import (
    "io"
    "net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w,"<h1>Hello ,this is my first web page!</h1>")
}

func main() {
    http.HandleFunc("/", firstPage)
    http.ListenAndServe(":8000", nil)
}