package mypkg

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
)

func GetMatchMetricAlamsWithTags(svc *cloudwatch.Client, prefix string, tags map[string]string) ([]TagsOfAlarm) {
	var tagsOfAlarms []TagsOfAlarm

	// 各アラートの詳細情報を配列として取得
	metricAlarms := getMetricAlarms(svc, prefix)

	_, tagsOfAlarms = getContainsTagsAlarms(svc, metricAlarms, tags)

	return tagsOfAlarms
}
