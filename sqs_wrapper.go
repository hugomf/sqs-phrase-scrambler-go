package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSWrapper struct {
	sqsClient *sqs.Client
	queueURL  string
	queueName string
	response  sqs.ReceiveMessageInput
}

func NewSQSWrapper(queueName string) (*SQSWrapper, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Print("configuration error, " + err.Error())
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	result, err := client.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})

	if err != nil {
		return nil, err
	}

	return &SQSWrapper{
		sqsClient: client,
		queueURL:  *result.QueueUrl,
		queueName: queueName,
	}, nil
}

func (s *SQSWrapper) SendMessage(message string) (*sqs.SendMessageOutput, error) {
	sMInput := &sqs.SendMessageInput{
		DelaySeconds: 2,
		MessageBody:  aws.String(message),
		QueueUrl:     aws.String(s.queueURL),
	}

	resp, err := s.sqsClient.SendMessage(context.TODO(), sMInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SQSWrapper) ReceiveMessage() (string, error) {

	gMInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            &s.queueURL,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   int32(5),
	}

	msgResult, err := s.sqsClient.ReceiveMessage(context.TODO(), gMInput)
	if err != nil {
		return "", err
	}

	if msgResult.Messages != nil {
		// fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
		// fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
		// fmt.Println("Message: " + *msgResult.Messages[0].Body)
		return *msgResult.Messages[0].Body, nil
	}

	return "", nil
}

func (s *SQSWrapper) PopMessage() (string, error) {

	gMInput := &sqs.ReceiveMessageInput{
		QueueUrl:            &s.queueURL,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   int32(5),
		WaitTimeSeconds:     int32(5),
	}
	msgResult, err := s.sqsClient.ReceiveMessage(context.TODO(), gMInput)
	if err != nil {
		return "", err
	}

	if msgResult.Messages != nil {
		// fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
		// fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
		// fmt.Println("Message: " + *msgResult.Messages[0].Body)
		dMInput := &sqs.DeleteMessageInput{
			QueueUrl:      &s.queueURL,
			ReceiptHandle: *&msgResult.Messages[0].ReceiptHandle,
		}
		_, err := s.sqsClient.DeleteMessage(context.TODO(), dMInput)
		if err != nil {
			return "", err
		}

		return *msgResult.Messages[0].Body, nil

	}

	return "", nil
}
