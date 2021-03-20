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

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables CloudWatchAlarm actions",
	Long:  "Disables CloudWatchAlarm actions",

	Run: func(cmd *cobra.Command, args []string) {
		prefix := viper.GetString("prefix")
		region := viper.GetString("region")
		tags := viper.GetStringMapString("tags")

		svc := mypkg.MakeCloudWatchSVC(region)
		metricAlarms := mypkg.GetMatchMetricAlams(svc, prefix, tags)

		// Disables
		alarmNames := []string{}
		for _, v := range metricAlarms {
			alarmNames = append(alarmNames, *v.AlarmName)
		}

		params := &cloudwatch.DisableAlarmActionsInput{
			AlarmNames: alarmNames,
		}
		_, err := svc.DisableAlarmActions(context.TODO(), params)

		if err != nil {
			log.Fatalf("Failure Disable Alarm Actions. %v", err)
		} else {
			fmt.Println(color.Green.Sprint("Success"))
		}
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)
}
