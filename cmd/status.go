package cmd

import (
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
		prefix := viper.GetString("prefix")
		region := viper.GetString("region")
		tags := viper.GetStringMapString("tags")

		svc := mypkg.MakeCloudWatchSVC(region)
		metricAlarms := mypkg.GetMatchMetricAlams(svc, prefix, tags)

		// テーブル形式で出力
		mypkg.OutputTable(metricAlarms)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("prefix", "p", "", "Alam Name Preix")
	statusCmd.Flags().StringP("region", "r", "us-east-1", "Target AWS Region")
	statusCmd.Flags().StringToStringP("tags", "t", nil, "Input the tags you want to filter. ex) project=test,env=dev")

	viper.BindPFlag("prefix", statusCmd.Flags().Lookup("prefix"))
	viper.BindPFlag("region", statusCmd.Flags().Lookup("region"))
	viper.BindPFlag("tags", statusCmd.Flags().Lookup("tags"))
}
