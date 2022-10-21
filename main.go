package main

import (
	"fmt"
	"os"
)

func main()  {
	host := os.Getenv("DB_HOST")
	fmt.Println(host)

}