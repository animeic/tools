### 获取文档目录

> 目标：获取当前文件目录，并建立md链接

- 忽略文件

index.md index .git main.go go.mod 编译后的 mian

- 数据类型

```go
type FileInfo struct{
    Path string
    Name string
    IsShow bool    // 如果是dir是否可链接 true
}
type Index struct{
    Name string // 文件或者目录名，不含后缀名
    Id string // 与 IndexLink Id对应 name_time
}
type IndexLink struct{
    Id string
    PathFile string // 链接路径 path/name.md
}
// [写作][]
```

- 遇到的问题

递归函数局部存储变量出现问题。append数据不对。需要用全局变量处理

