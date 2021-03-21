package mypkg

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

func GetMatchMetricAlams(svc *cloudwatch.Client, prefix string, tags map[string]string) []types.MetricAlarm {

	// 各アラートの詳細情報を配列として取得
	metricAlarms := getMetricAlarms(svc, prefix)

	// Filter by tags
	if len(tags) > 0 {
		metricAlarms, _ = getContainsTagsAlarms(svc, metricAlarms, tags)
	}

	return metricAlarms
}
