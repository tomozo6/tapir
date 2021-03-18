package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tapir/mypkg"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables CloudWatchAlarms",
	Long:  "Enables CloudWatchAlarms",

	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		prefix := viper.GetString("prefix")
		region := viper.GetString("region")
		tags := viper.GetStringMapString("tags")

		svc := mypkg.MakeCloudWatchSVC(region)
		metricAlarms := mypkg.GetMatchMetricAlams(svc, prefix, tags)

		// Enable
		alarmNames := []string{}
		for _, v := range metricAlarms {
			alarmNames = append(alarmNames, *v.AlarmName)
		}

		params := &cloudwatch.EnableAlarmActionsInput{
			AlarmNames: alarmNames,
		}
		_, err := svc.EnableAlarmActions(context.TODO(), params)

		if err != nil {
			log.Fatalf("Failure Enable Alarm Actions. %v", err)
		} else {
			fmt.Println("Success")
		}

		end := time.Now()
		fmt.Println(end.Sub(start))
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
