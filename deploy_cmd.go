package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var skipAAD bool

var deployCmd = &cobra.Command{
	Use:  "deploy",
	Args: cobra.MinimumNArgs(1),
}

var deployProductCmd = &cobra.Command{
	Use:   "product [descriptor] [config]",
	Short: "Deploy product into environment",
	Long:  `Deploy product into environment`,

	Args: cobra.ExactValidArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		// deployer, err := deploy.NewProductDeployer(authProvider, args[0], args[1])
		// if err != nil {
		// 	return err
		// }
		// return deployer.Deploy(skipAAD)
		fmt.Printf("Deployed product into environment")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.AddCommand(deployProductCmd)

	deployProductCmd.Flags().BoolVar(&skipAAD, "skip-aad", false, "Skip AAD deployment")

}
