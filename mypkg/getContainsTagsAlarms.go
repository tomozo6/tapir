package mypkg

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type TagsOfAlarm struct {
	AlarmName string
	AlarmTags map[string]string
}

func getContainsTagsAlarms(svc *cloudwatch.Client, metricAlarms []types.MetricAlarm, searchTags map[string]string) ([]types.MetricAlarm, []TagsOfAlarm) {
	var wg sync.WaitGroup
	var matchMetricAlarms []types.MetricAlarm
	var tagsOfAlarms []TagsOfAlarm

	limit := 10
	slots := make(chan struct{}, limit)

	for _, v := range metricAlarms {

		wg.Add(1)

		// 並列処理のために無名関数化。for文のvをそのまま渡している
		go func(v types.MetricAlarm) {
			slots <- struct{}{}
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

			// CloudWatchAlarmのタグにsearchTagsが全て含まれていたら
			//   - metricAlarm を matchMetricAlarms に詰める
			//   - tagsofalarms に 情報を詰める
			if tagContains(searchTags, alarmTags) {
				matchMetricAlarms = append(matchMetricAlarms, v)

				tagsOfAlarm := TagsOfAlarm{
					AlarmName: *(v.AlarmName),
					AlarmTags: alarmTags,
				}
				tagsOfAlarms = append(tagsOfAlarms, tagsOfAlarm)
			}
			<-slots
		}(v)
		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
	return matchMetricAlarms, tagsOfAlarms
}
