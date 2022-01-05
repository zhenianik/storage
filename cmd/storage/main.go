package main

import (
	"fmt"
	"github.com/zhenianik/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}
