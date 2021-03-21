package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tapir/mypkg"
)

var showtagsCmd = &cobra.Command{
	Use:   "showtags",
	Short: "show tags for CloudWatchAlarm",
	Long:  "show tags for CloudWatchAlarm",

	Run: func(cmd *cobra.Command, args []string) {
		prefix := viper.GetString("prefix")
		region := viper.GetString("region")
		tags := viper.GetStringMapString("tags")

		svc := mypkg.MakeCloudWatchSVC(region)
		tagsOfAlarms := mypkg.GetMatchMetricAlamsWithTags(svc, prefix, tags)

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U0001F449 {{ .AlarmName | red}}",
			Inactive: "  {{ .AlarmName | cyan }}",
			Selected: "\U0001F449 {{ .AlarmName | faint}}",
			Details: `
--------- Tags ----------
{{ range $key, $value := .AlarmTags }}
{{"{"}}{{ $key }}: {{ $value }}{{"}"}}
{{ end }}`,
		}

		// 一覧結果絞り込み処理を自前で定義できます。
		// 絞り込みの条件など細かく処理をかけますが、特に定義しない場合は絞り込み機能が無効になります。
		// searcher := func(input string, index int) bool {
		// 	pepper := tagsOfAlarms[index]
		// 	// name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		// 	// input = strings.Replace(strings.ToLower(input), " ", "", -1)
		// 	return strings.Contains(name, input)
		// }

		// 本体の定義
		prompt := promptui.Select{
			Label:     "Which alarms tags do you want to show?",
			Items:     tagsOfAlarms, //  Itemsに定義した配列のデータをそのまま渡せばOKです。
			Templates: templates,
			Size:      10, // 一覧に表示するデータの件数を定義、この数を超えた場合はスクロールします。
			// Searcher:  searcher,
		}

		// 実行処理、Prompt同様、特定の操作でエラーが返ってくるので、同様のエラー処理が必要になります。
		// 戻り値は、第1は選択したデータのindex、第2は選択した行の文字列が返ってきます。
		_, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// fmt.Printf("You choose number %d: %s\n", i+1, aws.String(tagsOfAlarms[i].AlarmName))
		fmt.Printf("\n")

	},
}

func init() {
	rootCmd.AddCommand(showtagsCmd)
}
