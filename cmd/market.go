package cmd

// marketCmd represents the market command
// var marketCmd = &cobra.Command{
// 	Use:     "market",
// 	Short:   "Directly open github market page",
// 	Args:    cobra.RangeArgs(0, 1),
// 	Example: "icons/market.svg",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		markets := []Metadata{
// 			{item: "Actions", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://github.com/marketplace?type=actions"},
// 			{item: "Apps", icon: &aw.Icon{Value: "icons/default.svg"}, url: "https://github.com/marketplace?type=apps"},
// 		}
// 		for _, m := range markets {
// 			wf.NewItem(m.item).Icon(m.icon).Subtitle(m.subtitle).Arg(m.url).Valid(true).UID(m.item).Autocomplete(m.item).IsFile(true)
// 		}
// 		if len(args) > 0 {
// 			wf.Filter(args[0])
// 		}
// 		wf.SendFeedback()
// 	},
// }
//
// func init() {
// 	listCmd.AddCommand(marketCmd)
// }
