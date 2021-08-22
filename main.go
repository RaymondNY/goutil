package main

import (
	"zqx.com/goutil/uncompress"
)

func main() {
	// err := a.Unarchive("/home/files/虚拟机.zip", "/home/files")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	uncompress.Unzip("/home/files/虚拟机.zip", "/home/files")
}
