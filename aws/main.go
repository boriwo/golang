//
// Golang Workshop 2024
//

package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"os"
)

const (
	region   = "us-west-2"
	queueURL = "https://sqs.us-west-2.amazonaws.com/447701116110/bwtest"
)

type Message struct {
	Payload string
	Count   int
}

func sendMessage(message *Message) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return "", err
	}
	svc := sqs.New(sess)
	buf, err := json.Marshal(message)
	if err != nil {
		return "", err
	}
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody:  aws.String(string(buf)),
		QueueUrl:     aws.String(queueURL),
		DelaySeconds: aws.Int64(0),
	})
	if err != nil {
		return "", err
	}
	return *result.MessageId, nil
}

func receiveMessage() (*Message, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}
	svc := sqs.New(sess)
	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: aws.Int64(1),  // Maximum number of messages to receive (1 in this case)
		VisibilityTimeout:   aws.Int64(30), // Visibility timeout in seconds (message will be invisible for 30 seconds)
		WaitTimeSeconds:     aws.Int64(20), // Long polling wait time in seconds (20 in this case)
	})
	if err != nil {
		return nil, err
	}
	// note there can only be one message in the response
	for _, message := range result.Messages {
		var msg Message
		err = json.Unmarshal([]byte(*message.Body), &msg)
		if err != nil {
			return nil, err
		}
		_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: message.ReceiptHandle,
		})
		if err != nil {
			return &msg, err
		} else {
			return &msg, nil
		}
	}
	return nil, nil
}

func receive(ch chan Message, done chan struct{}) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return err
	}
	svc := sqs.New(sess)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
					QueueUrl:            aws.String(queueURL),
					MaxNumberOfMessages: aws.Int64(1),
					VisibilityTimeout:   aws.Int64(30),
					WaitTimeSeconds:     aws.Int64(20),
				})
				if err != nil {
					logger.Println(err)
				}
				for _, message := range result.Messages {
					var msg Message
					err = json.Unmarshal([]byte(*message.Body), &msg)
					if err != nil {
						logger.Println(err)
					}
					ch <- msg
					_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
						QueueUrl:      aws.String(queueURL),
						ReceiptHandle: message.ReceiptHandle,
					})
					if err != nil {
						logger.Println(err)
					}
				}
			}
		}
	}()
	return nil
}

func receiveForever(ch chan Message, done chan struct{}) error {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return err
	}
	svc := sqs.New(sess)
	go func() {
		for {
			result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(queueURL),
				MaxNumberOfMessages: aws.Int64(1),
				VisibilityTimeout:   aws.Int64(30),
				WaitTimeSeconds:     aws.Int64(20),
			})
			if err != nil {
				logger.Println(err)
			}
			for _, message := range result.Messages {
				var msg Message
				err = json.Unmarshal([]byte(*message.Body), &msg)
				if err != nil {
					logger.Println(err)
				}
				ch <- msg
				_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
					ReceiptHandle: message.ReceiptHandle,
				})
				if err != nil {
					logger.Println(err)
				}
			}
		}
	}()
	return nil
}

func main() {
	id, err := sendMessage(&Message{"go workshop 2024 go", 0})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
	ch := make(chan Message)
	done := make(chan struct{})
	receive(ch, done)
	msg := <-ch
	fmt.Println(msg.Payload)
	done <- struct{}{}
	/*msg, err := receiveMessage()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg.Payload)
	}*/
}
