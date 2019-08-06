* Json key缺少雙引號並加上雙引號
`/(\s*?{\s*?|\s*?,\s*?)(['"])?([a-zA-Z0-9]+)(['"])?:/`, `$1"$3":`
```go
var re = regexp.MustCompile(`(\s*?{\s*?|\s*?,\s*?)(['"])?([a-zA-Z0-9]+)(['"])?:`)
s := re.ReplaceAllString(jsonData, `$1"$3":`)
```
  
* 雙引號包住的內容
`(").*?(")`

