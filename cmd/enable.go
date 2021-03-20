package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tapir/mypkg"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables CloudWatchAlarm actions",
	Long:  "Enables CloudWatchAlarm actions",

	Run: func(cmd *cobra.Command, args []string) {
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
			fmt.Println(color.Green.Sprint("Success"))
		}
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)
}
