/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var langs = []string{"javascript", "python", "java", "ruby", "php", "c++", "c#", "typescript", "shell", "go", "swift", "objective-c", "html", "css", "r", "powershell", "kotlin", "rust", "matlab", "vue", "perl", "scala", "groovy", "julia", "lua", "hcl", "dart", "coffeescript", "objective-c++", "erlang", "haskell", "apex", "emacs lisp", "scheme", "clojure", "f#", "vim script", "crystal", "elixir", "assembly", "jupyter notebook", "ocaml", "purescript", "reason", "tcl", "puppet", "fortran", "abap", "sas", "common-lisp", "hack", "racket", "awk", "cobol", "webassembly", "xslt", "autohotkey", "smalltalk", "nim", "vala", "cuda"}

// trendCmd represents the trend command
var trendCmd = &cobra.Command{
	Use:   "trending",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://github.com/trending/%s?since=daily"
		for _, lang := range langs {
			wf.NewItem(lang).Title(lang).Subtitle(fmt.Sprintf("Looking for %s trending", lang)).Arg(fmt.Sprintf(url, lang)).Valid(true)
		}
		if len(args) > 0 {
			wf.Filter(args[0])
		}
		wf.SendFeedback()
	},
}

func init() {
	rootCmd.AddCommand(trendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
