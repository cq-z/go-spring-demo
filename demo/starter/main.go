package main

import (
	_ "github.com/cq-z/go-spring-demo/demo/admin/cmd"
	SpringBoot "github.com/go-spring/spring-boot"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string

	rootCmd = &cobra.Command{
		Short: "A based Applications",
	}
)
func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

}
func main() {

	for _ , cmd := range SpringBoot.GetPrefixProperties("cmd") {
		rootCmd.AddCommand(cmd.(*cobra.Command))
	}
	rootCmd.Execute()
}