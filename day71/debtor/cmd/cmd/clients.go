package cmd

import (
	"log"
	"strconv"

	"github.com/fatih/color"
	"github.com/gjae/debtor/internal/repository"
	"github.com/gjae/debtor/pkg/tabler"
	"github.com/spf13/cobra"
)

var name string
var lastname string
var ced string

var NewClientCommand = &cobra.Command{
	Use:   "add",
	Short: "Creates a new client if not exists",
	Long:  "This subcommand creates new clients, require flags with the data",
	Run: func(cmd *cobra.Command, args []string) {
		clientExists, err := repository.NewClient(name, lastname, ced)

		if err != nil {
			log.Fatal(err.Error())
			return
		}

		(color.New(color.FgWhite, color.BgGreen, color.Bold)).Printf("Client: %s was been created", clientExists.Name)
	},
}

var ListClientsCommand = &cobra.Command{
	Use:   "list",
	Short: "List all clients registered",
	Run: func(cmd *cobra.Command, args []string) {
		clients := repository.ListClients()
		tabulator := tabler.New(tabler.CornerCharacter("+"), tabler.LineCharacter("-"))
		tabulator.SetColumnsName(tabler.ColumnNames([]string{"ID", "Nombre", "Apellido", "Cedula"}))

		for _, client := range clients {
			column := []tabler.Column{
				tabler.Column{Id: int(client.ID) + 1, Value: strconv.Itoa(int(client.ID))},
				tabler.Column{Id: int(client.ID) + 1, Value: client.Name},
				tabler.Column{Id: int(client.ID) + 1, Value: client.Lastname},
				tabler.Column{Id: int(client.ID) + 1, Value: client.Ced},
			}
			tabulator.AddRow(&tabler.Row{Id: int(client.ID), Coumns: column})
		}

		tabulator.Show()
	},
}

var ClientCommand = &cobra.Command{
	Use:   "clients",
	Short: "Management clients records in app",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
