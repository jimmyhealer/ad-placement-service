# 廣告投放服務

作者：簡蔚驊 jimmyhealer

## 使用方法

將 .env.example 複製一份並命名為 .env，並且修改裡面的環境變數。

```bash
docker-compose up -d
```

## 系統架構與想法

架構採用 Clean Architecture，並且使用 Wire 來管理依賴注入。

```
├── Dockerfile
├── README.md
├── api
│   └── v1
│       ├── ad.go
│       ├── ad_test.go
│       └── utils.go
├── apitest.sh
├── coverage.out
├── db
│   ├── db.go
│   ├── db_test.go
│   └── interface.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── models
│   ├── advertisement.go
│   ├── conditions.go
│   └── enum.go
├── repositories
│   ├── adRepository.go
│   ├── adRepository_test.go
│   └── interface.go
├── router.go
├── test_utils
│   └── db.go
├── wire.go
└── wire_gen.go
```

## 系統效能優化

### 資料庫查詢效能

首先可以分析出這個系統的瓶頸在於資料庫的查詢，因此我們可以透過建立索引來提升查詢效能。\
而其中 StartAt 與 EndAt 這兩個欄位是我們最常用來查詢的欄位，因此我們可以建立索引來提升效能。\
並採用複合索引的方式，將 StartAt 與 EndAt 這兩個欄位建立索引，並且設定優先順序，讓查詢效能更好。

```go
type Advertisement struct {
	gorm.Model
	ID         uint         `gorm:"primaryKey" `
	Title      string       `gorm:"type:varchar(255);not null"`
	StartAt    time.Time    `gorm:"type:timestamp with time zone;not null;index:idx_member,priority:2"`
	EndAt      time.Time    `gorm:"type:timestamp with time zone;not null;index:idx_member,priority:1"`
	Conditions []Conditions
}
```

