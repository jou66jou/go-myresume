## gotest檔案說明

```
├── debug                       * 個人debug使用
│   └── http.go                 # 輸出request.body
├── docker
│   ├── nsq
│   │   └── docker-compose.yml  # nsq docker
│   └── sql
│       └── docker-compose.yml  # sql db docker
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
│   ├── initdb                  
│   │   └── p2p.ws.go           # sql init
│   └── user                    * 一般使用者相關功能
│       └── user.go             # 登入、註冊、查詢方法
│── practice                    * 特定練習
│   ├── syncmap.go              # map併發讀寫
│   ├── p2p-websocket           # websocket模擬p2p
│   │   └── p2p.ws.go 
│   └── testinterface           # interface練習      
│── regular                    
│   └─── README.md              # 正則相關
└── studygolang.md              # golang中文社群回覆紀錄      
```


## golang筆記

### Pool

1. 大連接或資料庫大量寫入時採用Pool避免漏失訊息，也能節省記憶體的創建又銷毀成本
2. 用channel建立工作池，並建立排隊佇列
3. 高頻率短websocket連線可使用epoll

### 記憶體

1. 注意記憶體逃逸，函數除了大資料外不以傳址返回
2. 堆上分配會在GC時影響效能

### String

1. string為`struct{str *unsafe.Point, len int}`，因此string頻繁更改會重新指向造成GC回收負擔且效能降低。
2. 大量合併或拼接時使用`join`或`bytes.Buffer`。
3. 如果大量字串處理可考慮直接使用`[]byte`。
4. 字串與 []byte 之間的轉換是複製（有記憶體損耗），可以用map[string] []byte建立字串與[]byte之間對映，也可range來避免記憶體分配來提高效能
```go
//[]byte: 
for i,v := range []byte(str) {
    // do something
}
```
5. 使用for range迭代String，是以rune來迭代的。一個字元，也可以有多個rune組成。需要處理字元，儘量使用 golang.org/x/text/unicode/norm 包。

### Map

1. 並發讀寫會出錯。
2. 加入讀寫鎖`sync.RWMutex`。
3. 超過四核心處理器的競爭鎖，建議使用`sync.Map`實現。
4. 使用for range迭代map時每次迭代的順序可能不一樣，因為map的迭代是隨機的。


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
### Channel

1. 非緩衝channel執行時，若沒有進入goroution會報錯`fatal error: all goroutines are asleep - deadlock!`。



### RESTful API 範例

**/projects**
* `GET` : 取得全部project
* `POST` : 建立新project

**/projects/:title**
* `GET` : 取得此project
* `PUT` : 更新此project
* `DELETE` : 刪除此project

**/projects/:title/archive**  
* `PUT` : 封存此project
* `DELETE` : 回復此project

**/projects/:title/tasks**  
* `GET` : 取得此project所有任務
* `POST` : 新增一個任務到此project

**/projects/:title/tasks/:id**
* `GET` : 取得此project的該任務
* `PUT` : 更新此project的該任務
* `DELETE` : 刪除此project的該任務

**/projects/:title/tasks/:id/complete**
* `PUT` : 將此project的該任務設為complete
* `DELETE` : 將此project的該任務設為undo