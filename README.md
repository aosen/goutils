# utils
Golang工具箱

* 将Jsonp转化为Json
JsonpToJson modify jsonp string to json string
Example: forbar({a:"1",b:2}) to {"a":"1","b":2}
```Golang
JsonpToJson(json string) string
```

* 获取工作路径
```Golang
GetWDPath() string {
```

* 判断目录是否存在
```Golang
IsDirExists(path string) bool
```

* 判断文件是否存在
```Golang
IsFileExists(path string) bool
```

* 判断字符串是否为数字字符串
```Golang
IsStringNum(a string) bool
```

* 将xml转化为map[string]string
```Golang
XML2mapstr(xmldoc string) map[string]string {
```

* 将字符串转化成hash
```Golang
MakeHash(s string) string
```
