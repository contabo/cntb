package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// historyCmd represents the history command
var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Shows history of your resources",
	Long:  `Displays and filters what happend to your resources over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if viper.GetString("page") != "" && viper.GetString("page") != "1" {
			Page = viper.GetInt64("page")
		}
		if viper.GetString("size") != "" && viper.GetString("size") != "100" {
			Size = viper.GetInt64("size")
		}
		if viper.GetString("orderBy") != "" && viper.GetString("orderBy") != "name:asc" {
			OrderBy = viper.GetString("orderBy")
		}
		if viper.GetString("requestId") != "" {
			RequestIdFilter = viper.GetString("requestId")
		}
		if viper.GetString("changedBy") != "" {
			ChangedByFilter = viper.GetString("changedBy")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(HistoryCmd)

	HistoryCmd.PersistentFlags().Int64VarP(&Page, "page", "p", int64(DefaultPage),
		`Page number to display.`)
	viper.BindPFlag("page", HistoryCmd.Flags().Lookup("page"))

	HistoryCmd.PersistentFlags().Int64VarP(&Size, "size", "s", int64(DefaultPageSize),
		`Number of elements per page.`)
	viper.BindPFlag("size", HistoryCmd.Flags().Lookup("size"))

	HistoryCmd.PersistentFlags().StringVarP(&OrderBy, "orderBy", "b", "timestamp:desc",
		`Order results by these fields. E.g. name:asc.`)
	viper.BindPFlag("orderBy", HistoryCmd.Flags().Lookup("orderBy"))

	HistoryCmd.PersistentFlags().StringVar(&RequestIdFilter, "requestId", "",
		`Filter by the requestId which led to the change.`)
	viper.BindPFlag("requestId", HistoryCmd.Flags().Lookup("requestId"))

	HistoryCmd.PersistentFlags().StringVarP(&ChangedByFilter, "changedBy", "c", "",
		`Filter by the user id who performed the change.`)
	viper.BindPFlag("changedBy", HistoryCmd.Flags().Lookup("changedBy"))
}
