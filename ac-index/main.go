package main

import (
	"docs/index"
	"log"
	"os"
)

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 无参数默认parse 和 有目录链接
	// sp 使用sampleparse & d 无目录链接
	fg := os.Args[1:]
	var is_show bool = true
	var is_sample bool = false
	for _, v := range fg {
		if v == "d" {
			is_show = false
		}
		if v == "sp" {
			is_sample = true

		}
	}
	if is_sample {
		indexfile2 := rootDir + "/" + "index1.md"
		index.SampleParse(rootDir, indexfile2, is_show)
	} else {
		indexfile := rootDir + "/" + "index.md"
		index.Parse(rootDir, indexfile, is_show)
	}

}
