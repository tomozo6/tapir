package mypkg

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

func getMetricAlarms(svc *cloudwatch.Client, prefix string) []types.MetricAlarm {

	var metricAlarms []types.MetricAlarm
    var nextToken *string

	// 全てのレコードを取得するためにfor文でループさせる(nextToken)
	for {
		// DescribeAlarmsのパラメーター設定
		params := &cloudwatch.DescribeAlarmsInput{}
		params.MaxRecords = aws.Int32(100)

		// Prefix
		if prefix == "" {
			params.AlarmNamePrefix = nil
		} else {
			params.AlarmNamePrefix = aws.String(prefix)
		}

		// NextToken(最初は必ずnil)
		if nextToken == nil {
			params.NextToken = nil
		} else {
			params.NextToken = nextToken
		}

		resp, err := svc.DescribeAlarms(context.TODO(), params)

		if err != nil {
			fmt.Printf("メトリクスの取得に失敗しました。 %v", err)
		}

		for _, v := range resp.MetricAlarms {
			metricAlarms = append(metricAlarms, v)
		}

		if resp.NextToken == nil {
			break
		}

		nextToken = resp.NextToken
	}

	return metricAlarms
}
