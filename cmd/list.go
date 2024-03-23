package cmd

// listCmd represents the list command
// var listCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "list all subcommands",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		sub := cmd.Commands()
// 		for _, c := range sub {
// 			if !c.Hidden {
// 				wf.NewItem(c.Name()).Valid(false).Icon(&aw.Icon{Value: c.Example}).Title(c.Name()).Subtitle(c.Short).Autocomplete(c.Name())
// 			}
// 		}
// 		if len(args) > 0 {
// 			wf.Filter(args[0])
// 		}
// 		wf.WarnEmpty("No matching commands", "Try a different query?")
// 		wf.SendFeedback()
// 	},
// }

func init() {
	// rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
