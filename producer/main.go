package main

import (
	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"log"
	"producer/controllers"
	"producer/services"
	"strings"
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

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		log.Println("error to created producer :", err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountControllers := controllers.NewAccountControllers(accountService)

	app := fiber.New()

	app.Post("/openAccount", accountControllers.OpenAccount)
	app.Post("/depositFund", accountControllers.DepositFund)
	app.Post("/withdrawFund", accountControllers.WithdrawFund)
	app.Post("/closeAccount", accountControllers.CloseAccount)

	app.Listen(":8083")
}

//func main() {
//	servers := []string{"localhost:9092"}
//	producer, err := sarama.NewSyncProducer(servers, nil)
//	if err != nil {
//		panic(err)
//	}
//	defer producer.Close()
//
//	msg := sarama.ProducerMessage{
//		Topic: "bondhi",
//		Value: sarama.StringEncoder("Hello world"),
//	}
//	p, o, err := producer.SendMessage(&msg)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("partition=%v,offset=%v", p, o)
//}
