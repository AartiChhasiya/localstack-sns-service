# localstack-sns-service
## Publisher of Pub-sub model implementation in golang using AWS SNS service with localstack

#### In a pub-sub model using SNS and SQS, messages are published to an SNS topic, and then delivered to subscribed SQS queues. So, in order to read messages in a pub-sub model, you need to read messages from the SQS queue that is subscribed to the SNS topic.

#### When a message is published to an SNS topic, the message is forwarded to all SQS queues that are subscribed to that topic. 

#### First, make sure you have the AWS SDK for Go installed. You can install it using the following command:
$ go get github.com/aws/aws-sdk-go

#### This example creates an SNS topic named "my-topic" and publishes a message to the topic.

#### I've implemented 2 ways to publish the message to the topic:
##### 1. Publish one-time message
##### 2. Publish messages continuously on the interval of 5 seconds.

#### Note that you'll need to have valid AWS credentials set up in order to run this example. You can set them up using the AWS CLI or by setting the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables.





