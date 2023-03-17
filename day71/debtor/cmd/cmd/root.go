package cmd

import "github.com/spf13/cobra"

var rootCommand = &cobra.Command{
	Use:   "debtor",
	Short: "Debtor it's an easier way to controll loans",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	InstallCommand.Flags().BoolVarP(&force, "force", "f", false, "Forece reinstall app")
	NewClientCommand.Flags().StringVarP(&name, "name", "n", "", "Client name")
	NewClientCommand.Flags().StringVarP(&lastname, "lastname", "l", "", "Client Lastname")
	NewClientCommand.Flags().StringVarP(&ced, "ced", "c", "", "Client ced")
	AccountLoanSubcommad.Flags().StringVarP(&ced, "ced", "c", "", "Client ced")
	AccountLoanSubcommad.Flags().Float32VarP(&amount, "amount", "a", 0.0, "Amount of new account")

	AccountLoanSubcommad.Flags().StringVarP(&description, "description", "m", "", "Description")

	rootCommand.AddCommand(VersionCmd)
	rootCommand.AddCommand(InstallCommand)
	rootCommand.AddCommand(ClientCommand)
	rootCommand.AddCommand(AccountCommand)

	ClientCommand.AddCommand(NewClientCommand)
	ClientCommand.AddCommand(ListClientsCommand)

	AccountCommand.AddCommand(AccountAddCommand)
	AccountAddCommand.AddCommand(AccountLoanSubcommad)
	AccountCommand.AddCommand(AccountOpenedListCommand)
}

func Executable() {
	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}
