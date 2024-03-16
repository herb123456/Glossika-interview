# Glossika Interview

## 題目

```text
請建立一個新的專案
環境 MySQL / Redis / Go 1.22 以上
完成底下的需求，並且將專案放到 Github 上，並且提供一份 README.md 說明如何使用你的專案

專案需要有使用者的註冊、登入、驗證 email 的功能
有一份推薦商品資料清單 API 讓使用者可以取得這份資料
這份推薦商品清單需要 Database query 3 秒才能取得，請放在 Redis 中，並且設定 10 分鐘過期
推薦商品 API 預計每分鐘會有 300 次的取用 
*寄送驗證信件的行為可以不用真的寄出信件，只需要呼叫一個空實作的寄送 Email function
*這是後端面試題目，無需實作前端部分，只需實作 API 功能即可

1. 註冊
  - 帳號為 email
  - 密碼 不少於六個字 不多於 16 個字 需要有一個大寫 有一個小寫 跟一個特殊符號 ()[]{}<>+-*/?,.:;"'_\|~`!@#$%^&=

2. 驗證 email (兩者擇其一實作，也可都實作)
  - 點擊 email 中的連結後，驗證 email
  - 發送驗證碼給使用者的 email，使用者輸入驗證碼後，驗證 email

3. 登入
  - 使用者可以使用 email 跟密碼登入

4. GET /recommendation  
  - 需要登入後才能使用
  - 回應推薦商品的資料
```

## 程式架構

```
.
├── config
│   ├── config.go
│   ├── config.yaml
│   └── config.yaml.example
├── controller
│   ├── login.go
│   ├── recommendation.go
│   ├── register.go
│   └── verifyemail.go
├── database
│   ├── connection.go
│   ├── migrations
│   └── models
│       └── user.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── db_connection.go
│   ├── error_handler.go
│   └── jwt_authentication.go
├── mycache
│   ├── interface.go
│   ├── keys.go
│   └── redis.go
├── myerror
│   ├── apperror.go
│   ├── general_error.go
│   ├── token_error.go
│   ├── user_error.go
│   └── validate_error.go
├── myredis
│   └── new.go
├── myvalidators
│   ├── init.go
│   └── password.go
├── README.md
├── response
│   └── response.go
├── router
│   └── route.go
├── services
│   ├── email.go
│   ├── user.go
│   └── user_service_test.go
└── utils
    └── random.go
```

使用Viper套件來讀取設定檔，設定檔的範例在`config/config.yaml.example`，請複製一份到`config/config.yaml`。

ORM使用GORM，model放在`database/models`中。

自定義了一個myerror套件，用來處理錯誤訊息，可以將特定錯誤對應到特定訊息與錯誤碼

統一處理錯誤的好處是封裝的response可以直接回傳錯誤訊息，而不用在每個地方都寫一次。

使用Gin來建立API。

使用JWT來做驗證。

## 環境

我使用docker-compose建構這次面試所需的環境，包含了MySQL、Redis、Go。

直接執行以下指令即可啟動環境：

```bash
docker-compose up
```

在Dockerfile中，我使用了multi-stage build，將編譯後的執行檔放到一個小的容器中，以減少容器的大小。

## APIs

在啟動後會有四個API endpoint：

```
[GIN-debug] POST   /register                 --> Glossika_interview/controller.UsersController.Register-fm (5 handlers)
[GIN-debug] POST   /verify-email             --> Glossika_interview/controller.UsersController.VerifyEmail-fm (5 handlers)
[GIN-debug] POST   /login                    --> Glossika_interview/controller.UsersController.Login-fm (5 handlers)
[GIN-debug] GET    /recommendation           --> Glossika_interview/controller.RecommendationController.GetRecommendation-fm (6 handlers)
```

分別為註冊、驗證email、登入、取得推薦商品資料。

### 註冊

註冊的API endpoint為`/register`，使用POST method。

Request body：

```json
{
    "email": "test@test.com",
    "password": "Password123!"
}
```

Response：

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "email": "test@test.com",
    "verified": false,
    "verification_code_expiry": "2024-03-16T17:33:28.60106188+08:00",
    "verification_at": "0001-01-01T00:00:00Z"
  }
}
```

curl：

```bash
curl -X POST "http://localhost:8080/register" -H "accept: application/json" -H "Content-Type: application/json" -d '{"email":"test@test.com","password":"Password123!"}'
```

### 驗證email

在呼叫完註冊API後，在程式執行的畫面上(stdout)會看到一個驗證碼，這個驗證碼是用來驗證email的。

驗證email的API endpoint為`/verify-email`，使用POST method。

有針對這隻API的單元測試，可以在`services/user_service_test.go`中找到。

Request body：

```json
{
    "email": "test@test.com",
    "code": "1234"
}
```

Response：

```json
{
  "code": 200,
  "message": "success",
  "data": null
}
```

curl：

```bash
curl -X POST "http://localhost:8080/verify-email" -H "accept: application/json" -H "Content-Type: application/json" -d '{"email":"test@test.com","code":"1234"}'
```

### 登入

登入的API endpoint為`/login`，使用POST method。

Request body：

```json
{
    "email": "test@test.com",
    "password": "Password123!"
}
```

Response：

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJiQGFhLmNvbSIsImV4cCI6MTcxMDY2NzEyMiwiaWF0IjoxNzEwNTgwNzIyLCJ1c2VyX2lkIjoxMCwidmVyaWZpZWQiOnRydWV9.MEyFjyJUenk_Xgd0XoLrUmaA9zOwoJOrnGHeihcg74w"
  }
}
```

curl：

```bash
curl -X POST "http://localhost:8080/login" -H "accept: application/json" -H "Content-Type: application/json" -d '{"email":"test@test.com","password":"Password123!"}'
```

### 取得推薦商品資料

在登入後會取得一組JWT Token，這個Token需要放在Request的Header中，Key為`Authorization`。

取得推薦商品資料的API endpoint為`/recommendation`，使用GET method。

因題目限制的300秒查詢時間，所以就不針對資料庫查詢做優化，直接假設查詢就是需要300秒（time.Sleep）

所以這隻API設計成非同步模式

在快取裡面找不到資料時，會回傳`{"code":200,"message":"success","data":{"message":"recommendation is generating, please wait"}}`，並且在背景執行查詢，查詢完後會放到快取中。

在查詢完成前，再次呼叫這隻API會回傳`{"code":200,"message":"success","data":{"message":"recommendation is generating, please wait"}}`。

查詢完成後，再次呼叫這隻API會回傳推薦商品資料。

在查詢並寫入快取時，會有一個mutex lock，避免短時間同一人有多個請求時，同時多次查詢與寫入快取的問題。

curl：

```bash
curl -X GET "http://localhost:8080/recommendation" \
     -H "accept: application/json" \
     -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJiQGFhLmNvbSIsImV4cCI6MTcxMDY2MDExMSwiaWF0IjoxNzEwNTczNzExLCJ1c2VyX2lkIjoxMCwidmVyaWZpZWQiOnRydWV9.XZcCTuFh5POeSgWCTnLTgxdT-ESBkubHgH_jWp6aiCE'
```
