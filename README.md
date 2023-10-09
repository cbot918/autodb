# autodb
a database automation tool to boost small side-project

## Features:
- [x] Create database
- [x] Migrate sql
- [x] Read db and table metadata
- [　] Generate model
- [　] Generate dao ( 目前應該做不太出來 )

## Dependeicies:
1. go1.19
2. docker
2. linux like system

## Getting Started
1. 看一下 .env 內容 ( 確認一下 PORT 跟 CONTAINER_NAME 沒有跟本地衝突 )
2. `go run .`

## Verify
1. db container 正確啟動
2. 資料庫有8張表及對應的資料

## Clean Up
```bash
docker container stop autodb
docker container rm autodb
```

## Notes:
- 實驗性質, 功能都極簡略
- 只支援 MySQL