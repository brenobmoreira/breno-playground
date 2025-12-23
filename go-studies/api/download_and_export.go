package main

import (
	"github.com/brenobmoreira/breno-playground/api/ftp"
)

func main() {
	err := ftp.DownloadDBC()
	if err != nil {
		panic(err)
	}
}
