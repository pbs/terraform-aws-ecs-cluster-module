package test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/sts"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func DeleteClusterInstances(cluster string) {
	ecsSvc := ecs.New(session.New())
	ec2Svc := ec2.New(session.New())
	input := &ecs.ListContainerInstancesInput{
		Cluster: aws.String(cluster),
	}
	result, err := ecsSvc.ListContainerInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	describeInput := &ecs.DescribeContainerInstancesInput{
		Cluster:            aws.String(cluster),
		ContainerInstances: result.ContainerInstanceArns,
	}
	containerInstances, err := ecsSvc.DescribeContainerInstances(describeInput)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecs.ErrCodeServerException:
				fmt.Println(ecs.ErrCodeServerException, aerr.Error())
			case ecs.ErrCodeClientException:
				fmt.Println(ecs.ErrCodeClientException, aerr.Error())
			case ecs.ErrCodeInvalidParameterException:
				fmt.Println(ecs.ErrCodeInvalidParameterException, aerr.Error())
			case ecs.ErrCodeClusterNotFoundException:
				fmt.Println(ecs.ErrCodeClusterNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	for _, containerInstance := range containerInstances.ContainerInstances {
		instanceID := *containerInstance.Ec2InstanceId
		input := &ec2.TerminateInstancesInput{
			InstanceIds: []*string{
				aws.String(instanceID),
			},
		}
		_, err := ec2Svc.TerminateInstances(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}
	}
}

func getAWSAccountID(t *testing.T) string {
	session, err := session.NewSession()
	if err != nil {
		t.Fatalf("Failed to create AWS session: %v", err)
		return ""
	}
	svc := sts.New(session)
	result, err := svc.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatalf("Failed to get AWS Account ID: %v", err)
		return ""
	}
	return *result.Account
}

func getAWSRegion(t *testing.T) string {
	session, err := session.NewSession()
	if err != nil {
		t.Fatalf("Failed to create AWS session: %v", err)
		return ""
	}
	return *session.Config.Region
}
