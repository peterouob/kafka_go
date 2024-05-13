package main

import (
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"

	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("error to init viper : ", err)
	}
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println("error to open mysql : ", err)
	}
	return db
}

func main() {
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	db := initDatabase()
	accountRepo := repositories.NewAccountRepository(db)
	accountEvent := services.NewAccountEventHandler(accountRepo)
	accountConsume := services.NewConsumerHandler(accountEvent)
	fmt.Println("Account Consumer started ...")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsume)
	}
}

// func main() {
// 	servers := []string{"localhost:9092"}
// 	consumer, err := sarama.NewConsumer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer consumer.Close()
// 	partition, err := consumer.ConsumePartition("bondhi", 0, sarama.OffsetNewest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer partition.Close()

// 	fmt.Println("Consumer start.")
// 	for {
// 		select {
// 		case err := <-partition.Errors():
// 			fmt.Println(err)
// 		case msg := <-partition.Messages():
// 			println(string(msg.Value))
// 		}
// 	}
// }
