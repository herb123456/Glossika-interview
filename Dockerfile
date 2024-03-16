FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用scratch作為最終映像基底
FROM scratch

# 複製從第一階段建構的執行檔
COPY --from=builder /app/main .

# 複製設定
COPY --from=builder /app/config /config

# 執行Go程式
CMD ["./main"]