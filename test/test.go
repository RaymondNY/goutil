package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"zqx.com/goutil/ttime"
	"zqx.com/goutil/zipcompress"
	"zqx.com/goutil/zippwdcompress"
)

var (
	logdir = "/data/logs"
)

func main() {
	t := ttime.GetPeriod(1631515159, 1631601559)
	var waitc []*os.File
	var names []string
	defer func() {
		for _, b := range waitc {
			b.Close()
			fmt.Println("=====")
		}
	}()
	var files []string
	var filesd []string
	root := "/home/zqx/project/logtest"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
		st := t[i]
		for _, file := range files {
			if zippwdcompress.IsFile(file) {
				if strings.Contains(file, st) {
					f, err := os.Open(file)
					if err != nil {
						fmt.Println(err.Error())
					}
					names = append(names, f.Name())
					waitc = append(waitc, f)
				}
			}
		}
	}
	docker := "/var/lib/docker/containers"
	errd := filepath.Walk(docker, func(path string, info os.FileInfo, err error) error {
		filesd = append(filesd, path)
		return nil
	})
	if errd != nil {
		panic(errd)
	}
	for _, file := range filesd {
		if strings.Contains(file, "json.log") {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err.Error())
			}
			names = append(names, f.Name())
			waitc = append(waitc, f)
		}
	}
	for _, a := range waitc {
		fmt.Println(a.Name())
	}
	e := zipcompress.Compress(waitc, "/home/a.zip")
	if e != nil {
		fmt.Println(e.Error())
	}
	e1 := zippwdcompress.Zip("/home/b.zip", "Xad@8031009", names)
	if e1 != nil {
		fmt.Println(e1.Error())
	}
	//tree("/home/zqx/project/logtest")
}

func tree(dstPath string) error {
	dstF, err := os.Open(dstPath)
	if err != nil {
		return err
	}
	defer dstF.Close()
	fileInfo, err := dstF.Stat()
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() { //如果dstF是文件

		fmt.Println(dstPath)
		return nil
	} else { //如果dstF是文件夹
		fmt.Println(dstF.Name())
		dir, err := dstF.Readdir(0) //获取文件夹下各个文件或文件夹的fileInfo
		if err != nil {
			return err
		}
		for _, fileInfo = range dir {
			err = tree(dstPath + "/" + fileInfo.Name())
			if err != nil {
				return err
			}
		}
		return nil
	}

}
