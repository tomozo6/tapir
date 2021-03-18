package mypkg


import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

func getArnsContainsTags(svc *cloudwatch.Client, metricAlarms []types.MetricAlarm, searchTags map[string]string) []types.MetricAlarm {
	var wg sync.WaitGroup
	var matchMetricAlarms []types.MetricAlarm

	for _, v := range metricAlarms {
		wg.Add(1)

		// 並列処理のために無名関数化。for文のvをそのまま渡している
		go func(v types.MetricAlarm) {
			defer wg.Done()

			params := &cloudwatch.ListTagsForResourceInput{
				ResourceARN: v.AlarmArn,
			}
			// CloudWatchAlarmのタグ情報を取得
			resp, err := svc.ListTagsForResource(context.TODO(), params)

			if err != nil {
				fmt.Printf("CloudWatchAlarmのタグ情報取得に失敗しました。 %v", err)
			}

			// アラームのタグ情報をマップ形式に詰め直し
			alarmTags := map[string]string{}
			for _, v := range resp.Tags {
				alarmTags[*v.Key] = *v.Value
			}
			// fmt.Printf("alarmTags: %v", alarmTags)
			// fmt.Printf("searchTags: %v", searchTags)

			// CloudWatchAlarmに設定されているタグをリストに詰める
			// AlarmTags := []Tag{}
			// for _, v := range resp.Tags {
			// 	AlarmTags = append(AlarmTags, Tag{Key: *v.Key, Value: *v.Value})
			// }

			// CloudWatchAlarmのタグにsearchTagsが全て含まれていたら
			// metricAlarm を matchMetricAlarms に詰める
			if tagContains(searchTags, alarmTags) {
				matchMetricAlarms = append(matchMetricAlarms, v)
			}
		}(v)
	}
	wg.Wait()
	return matchMetricAlarms
}
