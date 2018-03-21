package cmd


import (
	"superBrain/cli"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "superBrainServer",
	Short: "superBrainServer App cmd",
	Long: `superBrainServer is a tool to manage golang service`,
	Example:`
		superBrainServer cjqc
	`,
}

func init(){
	RootCmd.AddCommand(cli.TestDataCmd("cjqc"))
}
