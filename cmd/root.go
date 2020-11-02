package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "herocli",
	Short: "a brief description",
	Long:  "A longer description",
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(Config)

}
func Config() {
	//TODO para que essas configurações funcionem dentro de um container será necessário utiliza do $(PWD)
	viper.SetConfigName("project")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
	}
}
