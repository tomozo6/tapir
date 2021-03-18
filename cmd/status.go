package cmd

import (
	"fmt"
	"time"

	"tapir/mypkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tags map[string]string

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Displays the 'ActinsEnabled' and 'State' status of the CloudWatchAlarms",
	Long:  "Displays the 'ActinsEnabled' and 'State' status of the CloudWatchAlarms",

	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		// fmt.Println(args)
		// fmt.Println(tags["project"])
		// fmt.Println(tags["env"])

		prefix := viper.GetString("prefix")
		region := viper.GetString("region")
		tags := viper.GetStringMapString("tags")

		svc := mypkg.MakeCloudWatchSVC(region)
		metricAlarms := mypkg.GetMatchMetricAlams(svc, prefix, tags)

		// テーブル形式で出力
		mypkg.OutputTable(metricAlarms)

		end := time.Now()
		fmt.Println(end.Sub(start))
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	// statusCmd.Flags().StringToStringVarP(&tags, "tags", "t", nil, "Input the tags you want to filter. ex) project=test,env=dev")
}
