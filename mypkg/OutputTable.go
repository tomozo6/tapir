package mypkg

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/gookit/color"
	"github.com/olekukonko/tablewriter"
)

func OutputTable(metricAlarms []types.MetricAlarm) {

	data := [][]string{}
	for _, v := range metricAlarms {

		// ActionEnabledの色設定
		var enabled string
		switch aws.ToBool(v.ActionsEnabled) {
		case true:
			enabled = color.Green.Sprint("enable")
		default:
			enabled = color.Gray.Sprint("disable")
		}

		// StateValueの色設定
		var state string
		switch v.StateValue {
		case "ALARM":
			state = color.Red.Sprint(v.StateValue)
		case "OK":
			state = color.Green.Sprint(v.StateValue)
		default:
			state = color.Gray.Sprint(v.StateValue)
		}

		data = append(data, []string{aws.ToString(v.AlarmName), enabled, state})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"AlarmName", "ActionsEnabled", "AlarmsState"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output

}
