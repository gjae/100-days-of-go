package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/gjae/debtor/internal/repository"
	models "github.com/gjae/debtor/pkg"
	"github.com/gjae/debtor/pkg/tabler"
	"github.com/spf13/cobra"
)

var amount float32
var description string
var AccountStatusOpened = models.AccountStatusOpened
var AccountStatusClosed = models.AccountStatusClose

var AccountLoanSubcommad = &cobra.Command{
	Use:     "loan",
	Short:   "Add new loan to client",
	Example: "debtor accounts add loans --ced 0000000",
	Aliases: []string{"loans", "l"},
	Run: func(cmd *cobra.Command, args []string) {
		clientCed := strings.Trim(ced, " ")
		if clientCed == "" {
			color.Red("Need client ced")
			return
		}
		_, err := repository.NewAccount(ced, amount, description, models.AccountLoan)

		if err == nil {
			color.GreenString("New account has been created")
			return
		}

		color.Red(err.Error())
	},
}

var AccountAddCommand = &cobra.Command{
	Use: "add",
}

var AccountOpenedListCommand = &cobra.Command{
	Use:   "list-open",
	Short: "List all accounts created",
	Run: func(cmd *cobra.Command, args []string) {
		accs := repository.ListAccounts(AccountStatusOpened)
		tabulator := tabler.New(tabler.CornerCharacter("+"), tabler.LineCharacter("-"))
		tabulator.SetColumnsName(tabler.ColumnNames([]string{"ID", "Persona", "Descripcion", "Monto"}))
		// success := color.New(color.FgGreen)

		for _, acc := range accs {
			tabulator.AddRow(&tabler.Row{
				Id: int(acc.Amount),
				Coumns: []tabler.Column{
					tabler.Column{
						Id:    int(acc.ID),
						Value: fmt.Sprint(acc.ID),
					},
					tabler.Column{
						Id:    int(acc.ID) + 1,
						Value: fmt.Sprintf("%s %s", acc.Client.Name, acc.Client.Lastname),
					},
					tabler.Column{
						Id:    int(acc.ID) + 2,
						Value: acc.Description,
					},
					tabler.Column{
						Id:    int(acc.ID) + 3,
						Value: fmt.Sprintf("%.2f", acc.Amount),
					},
				},
			})
		}

		tabulator.Show()
	},
}

var AccountCommand = &cobra.Command{
	Use:   "accounts",
	Short: "Management account data for loans and incomings",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
