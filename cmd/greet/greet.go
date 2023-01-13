package greet

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdpkg "github.com/rewardenv/reward-greeter/cmd"
	"github.com/rewardenv/reward-greeter/internal/config"
	"github.com/rewardenv/reward-greeter/internal/logic"
)

func NewCmdGreet(c *config.Config) *cmdpkg.Command {
	var cmd = &cmdpkg.Command{
		Command: &cobra.Command{
			Use:     "greet name",
			Short:   "greet command description",
			Long:    `greet command longer description`,
			Version: c.AppVersion(),
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (
				[]string, cobra.ShellCompDirective,
			) {
				return nil, cobra.ShellCompDirectiveNoFileComp
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				err := logic.New(c).RunCmdGreet(cmd, args)
				if err != nil {
					return fmt.Errorf("error running greet command: %w", err)
				}

				return nil
			},
		},
		Config: c,
	}

	cmd.Flags().Bool("add-cakes", false, "enable this flag to add cakes")
	c.BindPFlag("add_cakes", cmd.Flags().Lookup("add-cakes"))

	return cmd
}
