載入包

- go get github.com/IBM/sarama

producer

- kafka-console-producer --bootstrap-server=localhost:9092 --topic=bondhi

consumer

- kafka-console-consumer --bootstrap-server=localhost:9092 --topic=bondhi

## 引入包

可以使用

` replace events => ../events`

來新增目錄文件當作包引用，接著使用go get events即可使用

## 使用kafka指令

### 查看kafka topic
    - kafka-topics --bootstrap-server=localhost:9092 --list
        - CloseAccountEvent
        - DepositFundEvent
        - OpenAccountEvent
        - WithdrawFundEvent

### 查看consumer
    - kafka-consumer-groups --bootstrap-server=localhost:9092 --list
### 創建producer
    - kafka-console-producer --bootstrap-server=localhost:9092 --topic=OpenAccountEvent
### 監聽consumer
    - kafka-console-consumer --bootstrap-server=localhost:9092 --include="CloseAccountEvent|DepositFundEvent|OpenAccountEvent|WithdrawFundEvent" --group=log

### 完整使用流程
1. 啟動docker compose
2. 在consumer 底下 go run main.go
3. 監聽consumer kafka-console-consumer --bootstrap-server=localhost:9092 --include="CloseAccountEvent|DepositFundEvent|OpenAccountEvent|WithdrawFundEvent" --group=log
4. 在producer 底下 go run main.go
5. 使用api test or curl 去測試api

### 操作
1. openAccount
   - curl -H 'content-type:application/json' localhost:8083/openAccount -d '{
     "AccountHolder":"Bond123",
     "AccountType": 1,
     "OpeningBalance":10000
    }' -i

2. depositFund
    - curl -H 'content-type:application/json' localhost:8083/depositFund -d '{
      "ID":"da3f0db0-fc1d-49fd-ad39-1562484aee0e",
      "Amount": 50
    }' -i

3. withdrawFund
    - curl -H 'content-type:application/json' localhost:8083/withdrawFund -d '{
      "ID":"da3f0db0-fc1d-49fd-ad39-1562484aee0e",
      "Amount": 5000
    }' -i
4. closeAccount
    - curl -H 'content-type:application/json' localhost:8083/closeAccount -d '{
      "ID":"da3f0db0-fc1d-49fd-ad39-1562484aee0e"
   }' -i