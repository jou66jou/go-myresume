## gotest檔案說明

```
├── debug                       * 個人debug使用
│   └── http.go                 # 輸出request.body
├── docker
│   ├── nsq
│   │   └── docker-compose.yml  # nsq docker
│   └── sql
│       └── docker-compose.yml  # sql db docker
├── main.go
├── models                      * 可直接引用模組
│   ├── auth
│   │   └── middleware.go       # jwt驗證
│   ├── helper
│   │   ├── respjson.go         # http response json
│   │   └── router_register.go  # 註冊route
│   ├── inisetting 
│   │   ├── conf
│   │   │   └── app.ini         # config文件
│   │   └── setting.go          # config初始設定
│   ├── spl_init.go             # sql init
│   └── user                    * 一般使用者相關功能
│       └── user.go             # 登入、註冊、查詢方法
└── practice                    * 特定練習
    └── syncmap.go              # map併發讀寫
```


## golang效能筆記

### Pool

1. 大連接或資料庫大量寫入時採用Pool避免漏失訊息，也能節省記憶體的創建又銷毀成本
2. 用channel建立工作池，並建立排隊佇列

### 記憶體

1. 注意記憶體逃逸，函數除了大資料外不以傳址返回
2. 堆上分配會在GC時影響效能

### String

1. string為`struct{str *unsafe.Point, len int}`，因此string頻繁更改會重新指向造成GC回收負擔且效能降低。
2. 大量合併或拼接時使用`join`或`bytes.Buffer`。
3. 如果大量字串處理可考慮直接使用`[]byte`。

### Map

1. 並發讀寫會出錯。
2. 加入讀寫鎖`sync.RWMutex`。
3. 超過四核心處理器的競爭鎖，建議使用`sync.Map`實現。

### Slice

1. 考慮以下目的，將stu設置為字典中的值：
```go
m := make(map[string]*student)
stus := []student{
    {Name: "zhou", Age: 24},
    {Name: "li", Age: 23},
    {Name: "wang", Age: 22},
}
for _, stu := range stus {
    m[stu.Name] = &stu // 將stu設置為字典中的值
}
```
以上程式碼會造成m字典中的值皆為`&student[lastIndex]`，欲達成原本目的應將原碼改為：
```go
for i, v := range stus {
    m[v.Name] = &stus[i] // 將stus[i]設置為字典中的值
}
```
2. slice delete