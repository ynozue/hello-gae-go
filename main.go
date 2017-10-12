package main

import (
	"os"
	"log"
	"net/http"
	"github.com/ynozue/hellow-gae-go/appengine"
	"fmt"
)

func main() {
	log.Printf("start pid %d\n", os.Getpid())
	// appengine package をロード
	fmt.Println(appengine.New)
	http.ListenAndServe(":8080", nil)
}
