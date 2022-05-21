/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mailvalidator/model"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Mailvalidator, Send.",
	Long: `Yubi Mailvalidator, Send Terminal... Started...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send called")
		fmt.Println(viper.GetViper().GetString("DB_WRITER_HOST"))

		jsonFile, err := os.Open("payload.json")

		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var payload model.Payload
		json.Unmarshal([]byte(byteValue), &payload)

		//fmt.Println(payload)

		fmt.Println(exist(payload.Recipients[0].Email))
		Send(payload.Recipients[0].Email)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

func Send(email string){
	fmt.Print("sending queue")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"mailvalidator_queue", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare a queue")

	body := email
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
		

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}