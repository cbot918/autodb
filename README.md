# AutoDB
a toy level db automation tool to boost small side-project

</br>

# Dependeicies:
1. go1.19
2. docker
2. linux like system

</br>

# Getting Started
### 用 go 執行
```bash
git clone https://github.com/cbot918/autodb
cd autodb && go run .
```
verify scripts
```bash
# 看資料庫有 8 張表
docker exec -it autodb mysql -uroot -p12345 autodb -e "SELECT COUNT(table_name)
	FROM information_schema.tables
	WHERE table_schema = 'autodb';"
```
```bash
# 撈資料出來看
docker exec -it autodb mysql -uroot -p12345  --default-character-set=utf8 autodb -e "SELECT * FROM t_goods;"
```

### 下載 cli 來用
1. 下載安裝
```bash
curl -OL https://github.com/cbot918/autodb/releases/download/v0.0.1/odb && sudo chmod +x odb && sudo mv odb /usr/local/bin
```
輸入 `odb` 這樣應該會看到 cli help

2. 開始
```bash
mkdir testodb && cd testodb
odb init      # init .env
odb createdb  # create mysql container
odb createsql # download sample.sql
odb migrate   # odb migrate
odb verifydb  # see table numbers
odb clean     # clean up container
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


# Reference:
### [簡易秒殺系統-Go語言實現](https://github.com/Nobodiesljh/seckill-golang)

### [祁老師SpringBoot秒殺實戰](https://www.bilibili.com/video/BV1i84y1i7zF?p=2&spm_id_from=pageDriver&vd_source=5e33dcdca19327cd3f2787b83dddbd6c)