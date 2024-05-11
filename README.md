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
