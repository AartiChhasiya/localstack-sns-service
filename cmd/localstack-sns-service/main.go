package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	// Create an AWS session on Localstack
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint:         aws.String("http://localhost:4566"),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
		},
	})
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	// Create an SNS client on Localstack
	snsClient := sns.New(sess)

	// Create a new SNS topic
	topic, err := snsClient.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String("test-topic"),
	})
	if err != nil {
		log.Fatalf("failed to create SNS topic: %v", err)
	}

	// fmt.Printf("Type of Topic ARN is: %T\n", topic.TopicArn)
	fmt.Printf("Topic ARN is: %s\n", *topic.TopicArn)

	// Publish a message to the SNS topic
	publishMessage(snsClient, *topic.TopicArn)

	// Publish a new message every 5 seconds continuously
	publishMessageContinuouslyOnInterval(snsClient, *topic.TopicArn, 5)

	// err = godotenv.Load(".env-local")
	// if err != nil {
	// 	log.Fatal("Error loading .env-local file")
	// }
	// envValue := os.Getenv("PUBLISH_MSG_CONTINUOUSLY")

	// if envValue == "" {
	// 	fmt.Println("Environment variable is empty")
	// 	return
	// }

	// // Convert the environment variable value to a boolean
	// postMsgContinuously, err := strconv.ParseBool(envValue)
	// if err != nil {
	// 	fmt.Println("Failed to parse environment variable:", err)
	// 	return
	// }

	// if postMsgContinuously {
	// 	// Publish a message to the SNS topic
	// 	publishMessage(snsClient, *topic.TopicArn)
	// } else {

	// 	interval, err := time.ParseDuration(os.Getenv("PUBLISH_MSG_INTERVAL"))
	// 	if err != nil {
	// 		fmt.Println("Failed to parse environment variable:", err)
	// 		return
	// 	}

	// 	// Publish a new message every 5 seconds continuously
	// 	publishMessageContinuouslyOnInterval(snsClient, *topic.TopicArn, interval)
	// }
}

func publishMessage(snsClient *sns.SNS, topicARN string) {
	message := "Hello, from Aarti Chhasiya!"
	result, err := snsClient.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topicARN),
	})
	if err != nil {
		log.Fatalf("failed to publish message: %v", err)
	}

	fmt.Printf("Published message with ID %s\n", *result.MessageId)
}

func publishMessageContinuouslyOnInterval(snsClient *sns.SNS, topicARN string, interval time.Duration) {
	for {
		// Generate a new message
		message := fmt.Sprintf("Hello, from Aarti Chhasiya! ==> Message published at %s", time.Now().Format(time.RFC3339))

		// Publish the message to the SNS topic
		_, err := snsClient.Publish(&sns.PublishInput{
			Message:  aws.String(message),
			TopicArn: aws.String(topicARN),
		})
		if err != nil {
			fmt.Println("Failed to publish message:", err)
		} else {
			fmt.Println("Message published:", message)
		}

		// Wait for 5 seconds before publishing the next message
		time.Sleep(interval * time.Second)
	}
}
