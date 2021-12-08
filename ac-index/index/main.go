package index

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type FileInfo struct {
	Path    string
	Name    string
	ModTime string
	Indent  string
	IsShow  bool
}

// 全局变量 存储递归遍历的数据
// 因为是递归，如果用函数返回数据，会不清楚是哪次调用返回的
var _fis []*FileInfo = nil

// todo 进一步 过滤目录链接 根据isShow
// 接收目录 遍历文件信息
// [][] []: $link 形式目录
func Parse(dir string, indexfile string, show bool) {
	// 递归获取数据，存储在全局变量中
	listFiles(dir, 0, show)
	var strs string = "### 目录 \n"
	for _, _fi := range _fis {
		// 格式化
		ix := &Index{}
		strs += ix.Parse(_fi)

	}
	strs += "\n\n\n\n"
	for _, _fi := range _fis {
		// 格式化
		ixl := &IndexLink{}
		strs += ixl.Parse(_fi)

	}
	_, err := WriteIndex(indexfile, strs)
	if err != nil {
		log.Fatal(err)
	}

}

// 另一种格式的链接目录[]()
func SampleParse(dir string, indexfile string, show bool) {
	listFiles(dir, 0, show)
	var strs string = "### 目录 \n"
	for _, _fi := range _fis {
		// 格式化
		ix := &Index{}
		strs += ix.SampleParse(_fi)
	}
	_, err := WriteIndex(indexfile, strs)
	if err != nil {
		log.Fatal(err)
	}
}

func splitName(name string) string {
	str := strings.Split(name, ".md")
	return str[0]
}

// 递归遍历出文件结构
func listFiles(fp string, level int, isshow bool) {
	indent := "- "
	for i := 0; i < level; i++ {
		indent = "  " + indent
	}
	// 读取文件
	infos, err := ioutil.ReadDir(fp)
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range infos {
		// 过滤特定文件
		if filterFile(fi) {
			continue
		}
		var fileinfo *FileInfo = nil
		// log.Println(fileinfo)
		// 如果当前仅为目录 不显示属性
		fileinfo = &FileInfo{
			Path:    fp,
			Name:    fi.Name(),
			ModTime: fi.ModTime().Local().Format("2006-01-02 15:04:05.03"),
			Indent:  indent,
		}
		if !isshow {
			fileinfo.IsShow = !fi.IsDir()
		} else {
			fileinfo.IsShow = true
		}
		_fis = append(_fis, fileinfo)
		if fi.IsDir() {
			filePath := fp + "/" + fi.Name()
			listFiles(filePath, level+1, isshow)
		}

	}
}

// 过滤文件

func filterFile(fi fs.FileInfo) bool {
	switch fi.Name() {
	case ".git", "index.md", "index1.md", "index", "main.go", 
	".gitignore","go.mod", "main":
		return true // 过滤这些文件
	default:
		return false
	}
}

// 写入文件
func WriteIndex(ixf string, str string) (n int, err error) {
	// 文件存在则删除
	if FileExist(ixf) {
		err := os.Remove(ixf)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.OpenFile(ixf, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	n, err = f.WriteString(str)
	defer f.Close()
	return
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func sha_256(str string, secret string) string {

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))[:16]
	// fmt.Println(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
