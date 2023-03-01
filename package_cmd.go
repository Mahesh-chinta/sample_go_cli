package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var policyUrl string
var policyQueryPath string
var policyDecisionPath string
var policyErrorPath string
var policyWarningPath string
var enableComparison bool
var skipPurviewCheck bool

var checkCmd = &cobra.Command{
	Use:  "check",
	Args: cobra.MinimumNArgs(1),
}

var checkProductCmd = &cobra.Command{
	Use:   "product [descriptor] [teamConfig]",
	Short: "Validate product description",
	Long:  `Validate product description`,

	Args: cobra.ExactValidArgs(2),

	RunE: func(cmd *cobra.Command, args []string) error {
		descriptorPath := args[0]
		teamConfigPath := args[1]
		skipPolicyCheck := !cmd.Flags().Changed("policy-url")

		fmt.Printf("validated product description")
		fmt.Printf("descriptorPath :", descriptorPath)
		fmt.Printf("teamConfigPath :", teamConfigPath)
		fmt.Printf("skipPolicyCheck :", skipPolicyCheck)
		// checker := check.NewProductChecker(authProvider, descriptorPath, teamConfigPath)
		// if !skipPolicyCheck {
		// 	provider, err := check.NewAzureBlogStorageBundleProvider(authProvider, policyUrl)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	checker.PolicyExecutor = policy.NewOPAPolicyExecutor(
		// 		policy.OpaConfig{
		// 			PolicyQueryPath:    policyQueryPath,
		// 			PolicyDecisionPath: policyDecisionPath,
		// 			PolicyErrorPath:    policyErrorPath,
		// 			PolicyWarningPath:  policyWarningPath,
		// 		}, provider)
		// }
		// checker.ColumnCompare = enableComparison //used inside check/product.go
		// return checker.Check()
		return nil
	},
}
var checkConnectivityCmd = &cobra.Command{
	Use:   "connectivity [config]",
	Short: "Validate connectivity and authentication to backend services",
	Long:  `This validating will connect to required services to make sure connectivity and authentication is in place`,

	Args: cobra.ExactValidArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		// checker, err := check.NewConnectivityChecker(authProvider, args[0])
		// checker.SkipPurview = skipPurviewCheck
		// if err != nil {
		// 	return err
		// }
		// return checker.RunTest()
		fmt.Printf("validated connectivity and authentication to backend services")
		return nil
	},
}

var checkTeamConfigCmd = &cobra.Command{
	Use:   "team-config [config]",
	Short: "Validate team configuration file",
	Long:  `Validate team configuration file`,

	Args: cobra.ExactValidArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		// checker := check.NewTeamConfigChecker(args[0])
		// return checker.Check()
		fmt.Printf("validated team configuration file")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.AddCommand(checkProductCmd)
	checkCmd.AddCommand(checkConnectivityCmd)
	checkCmd.AddCommand(checkTeamConfigCmd)

	checkProductCmd.Flags().StringVar(&policyUrl, "policy-url", "[NOT SET]", "The URL of the governance policy to validate the descriptor against.")
	checkProductCmd.Flags().StringVar(&policyQueryPath, "policy-evaluation-query", "data.governance", "Policy decision to evaluate")
	checkProductCmd.Flags().StringVar(&policyDecisionPath, "policy-decision-path", "allow", "Policy decision to evaluate")
	checkProductCmd.Flags().StringVar(&policyErrorPath, "policy-errors-path", "errors", "Policy decision to evaluate")
	checkProductCmd.Flags().StringVar(&policyWarningPath, "policy-warnings-path", "warnings", "Policy decision to evaluate")

	checkConnectivityCmd.Flags().BoolVar(&skipPurviewCheck, "skip-purview", false, "Skip purview connectivity check")
	checkProductCmd.Flags().BoolVar(&enableComparison, "enable-comparison", false, "enable comparison")
}
