package mypkg

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func MakeCloudWatchSVC(region string) *cloudwatch.Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)

	svc := cloudwatch.NewFromConfig(cfg)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return svc
}
