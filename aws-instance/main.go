package main

import (
	"context"
	"fmt"
	"strings"

	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

var (
	instanceID  string
	err         error
	ctx         context.Context
	region      string = "eu-central-1"
	keyPairName string = "go-aws-demo"
	pemFilename string = "go-aws.pem"
)

func main() {

	ctx = context.Background()
	instanceID, err = createEC2(region, ctx)
	if err != nil {
		fmt.Printf("create EC2 instance error %s", err)
		os.Exit(1)
	}

	fmt.Printf("Instance ID: %s\n", instanceID)
}

func createEC2(region string, ctx context.Context) (string, error) {

	// loading  AWS sdk fo go
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)

	keyPair, err := ec2Client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		KeyNames: []string{keyPairName},
	})
	if err != nil && !strings.Contains(err.Error(), "InvalidKeyPair.NotFound") {
		return "", fmt.Errorf("describe key pairs error, %v", err)
	}

	if keyPair == nil || len(keyPair.KeyPairs) == 0 {

		newKeyPair, err := ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
			KeyName: aws.String(keyPairName),
		})
		if err != nil {
			return "", fmt.Errorf("unable to create key pair, %v", err)
		}

		err = os.WriteFile(pemFilename, []byte(*newKeyPair.KeyMaterial), 0400)
		if err != nil {
			return "", fmt.Errorf("create key pair error %s", err)
		}

	}

	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})

	if err != nil {
		return "", fmt.Errorf("describe images error, %v", err)
	}
	if len(imageOutput.Images) == 0 {
		return "", fmt.Errorf("images is of 0 length, %v", err)
	}

	instance, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      imageOutput.Images[0].ImageId,
		KeyName:      aws.String("go-aws-demo"),
		InstanceType: types.InstanceTypeT3Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})

	if err != nil {
		return "", fmt.Errorf("describe images error, %v", err)
	}

	if len(instance.Instances) == 0 {
		return "", fmt.Errorf("there are 0 instances, %v", err)
	}
	return *instance.Instances[0].InstanceId, nil

}
