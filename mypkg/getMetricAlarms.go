package mypkg

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

func getMetricAlarms(svc *cloudwatch.Client, prefix string) []types.MetricAlarm {

	// DescribeAlarmsのパラメーター設定
	params := &cloudwatch.DescribeAlarmsInput{}
	params.MaxRecords = aws.Int32(100)

	if prefix == "" {
		params.AlarmNamePrefix = nil
	} else {
		params.AlarmNamePrefix = aws.String(prefix)
	}

	// params := &cloudwatch.DescribeAlarmsInput{
	// 	MaxRecords:      aws.Int32(100),
	// 	AlarmNamePrefix: nil,
	// 	// AlarmNamePrefix: aws.String(prefix),
	// 	// AlarmNamePrefix: aws.String("HTTP Monitoring"),
	// }
	// fmt.Println(prefix)

	resp, err := svc.DescribeAlarms(context.TODO(), params)

	if err != nil {
		fmt.Printf("メトリクスの取得に失敗しました。 %v", err)
	}

	metricAlarms := []types.MetricAlarm{}

	for _, v := range resp.MetricAlarms {
		metricAlarms = append(metricAlarms, v)
	}

	return metricAlarms
}
