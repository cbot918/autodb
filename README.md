# AutoDB
a toy level db automation tool to boost small side-project

</br>

# Dependeicies:
1. go1.19
2. docker
2. linux like system

</br>

# Getting Started
1. 看一下 .env 內容 ( 確認一下 PORT 跟 CONTAINER_NAME 沒有跟本地衝突 )
2. `go run .`

</br>

# Verify
1. db container 正確啟動
```bash
docker ps | grep autodb
```
2. 資料庫有8張表
```bash
docker exec -it autodb mysql -uroot -p12345 autodb -e "SELECT COUNT(table_name)
	FROM information_schema.tables
	WHERE table_schema = 'autodb';"
```
3. 讀一下 t_goods 表
```bash
docker exec -it autodb mysql -uroot -p12345  --default-character-set=utf8 autodb -e "SELECT * FROM t_goods;"
```

</br>

# Features:
- [x] Create database
- [x] Migrate sql
- [x] Read db and table metadata
- [　] Generate model
- [　] Generate dao ( 目前應該做不太出來 )

</br>

# Clean Up
```bash
docker container stop autodb
docker container rm autodb
```

</br>

# Notes:
- 實驗性質, 功能都極簡略
- 只支援 MySQL