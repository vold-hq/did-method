package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/bryk-io/id/client/store"
	"github.com/bryk-io/x/did"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Example: "bryk-id list",
	Short:   "List registered DIDs",
	RunE:    runListCmd,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runListCmd(cmd *cobra.Command, args []string) error {
	// Get store handler
	st, err := store.NewLocalStore(viper.GetString("home"))
	if err != nil {
		return err
	}
	list := st.List()
	if len(list) == 0 {
		fmt.Println("No DIDs registered for the moment")
		return nil
	}

	// Show list of registered entries
	table := tabwriter.NewWriter(os.Stdout, 8, 0, 4, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintf(table,"%s\t%s\t%s\n", "Reference Name", "Recovery Mode", "DID")
	for _, e := range list {
		id := &did.Identifier{}
		if err := id.Decode(e.Contents); err != nil {
			continue
		}
		_, _ = fmt.Fprintf(table,"%s\t%s\t%s\n", e.Name, e.Recovery, id.GetDID())
	}
	return table.Flush()
}
