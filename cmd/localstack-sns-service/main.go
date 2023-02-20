package main

import (
	"fmt"
	"log"

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

	// Publish a message to the SNS topic
	message := "Hello, world!"
	result, err := snsClient.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: topic.TopicArn,
	})
	if err != nil {
		log.Fatalf("failed to publish message: %v", err)
	}

	fmt.Printf("Published message with ID %s\n", *result.MessageId)
}
