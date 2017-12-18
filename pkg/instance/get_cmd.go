package instance

import (
	"github.com/Azure/service-catalog-cli/pkg/command"
	"github.com/Azure/service-catalog-cli/pkg/output"
	"github.com/spf13/cobra"
)

type getCmd struct {
	*command.Context
	ns string
}

// NewGetCmd builds a "svcat get instances" command
func NewGetCmd(cxt *command.Context) *cobra.Command {
	getCmd := &getCmd{Context: cxt}
	cmd := &cobra.Command{
		Use:     "instances [name]",
		Aliases: []string{"instance", "inst"},
		Short:   "List instances, optionally filtered by name",
		Example: `
  svcat get instances
  svcat get instance wordpress-mysql-instance
  svcat get instance -n ci concourse-postgres-instance
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getCmd.run(args)
		},
	}
	cmd.Flags().StringVarP(
		&getCmd.ns,
		"namespace",
		"n",
		"default",
		"The namespace in which to get the ServiceInstance",
	)
	return cmd
}

func (c *getCmd) run(args []string) error {
	if len(args) == 0 {
		return c.getAll()
	} else {
		name := args[0]
		return c.get(name)
	}
}

func (c *getCmd) getAll() error {
	instances, err := retrieveAll(c.Client, c.ns)
	if err != nil {
		return err
	}

	output.WriteInstanceList(c.Output, instances.Items...)
	return nil
}

func (c *getCmd) get(name string) error {
	instance, err := retrieveByName(c.Client, c.ns, name)
	if err != nil {
		return err
	}

	output.WriteInstanceList(c.Output, *instance)
	return nil
}
