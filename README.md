# 廣告投放服務

作者：簡蔚驊 jimmyhealer

## 使用方法

將 .env.example 複製一份並命名為 .env，並且修改裡面的環境變數。

```bash
docker-compose up -d
```

## 系統架構與想法

架構採用 Clean Architecture，並且使用 Wire 來管理依賴注入。
其中，因爲此專案商業邏輯較簡易，因此沒有使用到 UseCase 的部分，而是直接將 Repository 注入到 Controller 中。

### 架構圖

```
api.v1 (Controller) -> usecase (virtual) -> models (Entity) <- usecase (virtual) <- repositories <- db (Datasource)
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

### Redis Cache [未實現]

- 時間分段緩存
	- 按照時間分段緩存，將 startAt 與 endAt 這兩個欄位的時間分段，並且將這個時間段的廣告存入 Redis 中，這樣可以減少資料庫的查詢次數。
- 布隆過濾器
	- 用於 Condictions 判斷條件，有 age, gender, country, platfrom 四個條件
	- 總可能的組合為 100 * 4 * 8 * 8 = 25600 種
	- 因此可以使用布隆過濾器來判斷是否有這個條件的廣告，如果沒有的話就可以確定資料庫裡面沒有這個條件的廣告，不用再去查詢資料庫。