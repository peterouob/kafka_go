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
### 創建consumer
    - kafka-console-consumer --bootstrap-server=localhost:9092 --include="CloseAccountEvent|DepositFundEvent|OpenAccountEvent|WithdrawFundEvent" --group=log
