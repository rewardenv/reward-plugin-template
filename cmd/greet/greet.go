package greet

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdpkg "github.com/rewardenv/reward-greeter/cmd"
	"github.com/rewardenv/reward-greeter/internal/config"
	"github.com/rewardenv/reward-greeter/internal/logic"
)

func NewCmdGreet(conf *config.Config) *cmdpkg.Command {
	cmd := &cmdpkg.Command{
		Command: &cobra.Command{
			Use:     "greet name",
			Short:   "greet command description",
			Long:    `greet command longer description`,
			Version: conf.AppVersion(),
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (
				[]string, cobra.ShellCompDirective,
			) {
				return nil, cobra.ShellCompDirectiveNoFileComp
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				err := logic.New(conf).RunCmdGreet(cmd, args)
				if err != nil {
					return fmt.Errorf("error running greet command: %w", err)
				}

				return nil
			},
		},
		Config: conf,
	}

	cmd.Flags().Bool("add-cakes", false, "enable this flag to add cakes")
	_ = conf.BindPFlag("add_cakes", cmd.Flags().Lookup("add-cakes"))

	return cmd
}
