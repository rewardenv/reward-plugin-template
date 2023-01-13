package root

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdpkg "github.com/rewardenv/reward-greeter/cmd"
	"github.com/rewardenv/reward-greeter/cmd/greet"
	"github.com/rewardenv/reward-greeter/internal/config"
	"github.com/rewardenv/reward-greeter/internal/logic"
)

func NewCmdRoot(conf *config.Config) *cmdpkg.Command {
	conf.Init()

	cmd := &cmdpkg.Command{
		Command: &cobra.Command{
			Use:   "greeter [command]",
			Short: "greeter is a cli tool which helps you to create fantastic greetings",
			Long: `
██████╗ ███████╗██╗    ██╗ █████╗ ██████╗ ██████╗      ██████╗ ██████╗ ███████╗███████╗████████╗███████╗██████╗ 
██╔══██╗██╔════╝██║    ██║██╔══██╗██╔══██╗██╔══██╗    ██╔════╝ ██╔══██╗██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██████╔╝█████╗  ██║ █╗ ██║███████║██████╔╝██║  ██║    ██║  ███╗██████╔╝█████╗  █████╗     ██║   █████╗  ██████╔╝
██╔══██╗██╔══╝  ██║███╗██║██╔══██║██╔══██╗██║  ██║    ██║   ██║██╔══██╗██╔══╝  ██╔══╝     ██║   ██╔══╝  ██╔══██╗
██║  ██║███████╗╚███╔███╔╝██║  ██║██║  ██║██████╔╝    ╚██████╔╝██║  ██║███████╗███████╗   ██║   ███████╗██║  ██║
╚═╝  ╚═╝╚══════╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
                                                                                                                
`,
			Version:       conf.AppVersion(),
			SilenceErrors: conf.SilenceErrors(),
			SilenceUsage:  true,
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (
				[]string, cobra.ShellCompDirective,
			) {
				return nil, cobra.ShellCompDirectiveNoFileComp
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				err := logic.New(conf).RunCmdRoot(cmd)
				if err != nil {
					return fmt.Errorf("error running root command: %w", err)
				}

				return nil
			},
		},
		Config: conf,
	}

	addFlags(cmd)
	// Reinitialize config after command line flags are added
	conf.Init()

	cmd.AddCommands(greet.NewCmdGreet(conf))

	return cmd
}

func addFlags(cmd *cmdpkg.Command) {
	// --log-level
	cmd.PersistentFlags().String(
		"log-level", "info", "logging level (options: trace, debug, info, warning, error)",
	)
	_ = cmd.Config.BindPFlag("log_level", cmd.PersistentFlags().Lookup("log-level"))

	// --debug
	cmd.PersistentFlags().Bool(
		"debug", false, "enable debug mode (same as --log-level=debug)",
	)
	_ = cmd.Config.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	// --disable-colors
	cmd.PersistentFlags().Bool(
		"disable-colors", false, "disable colors in output",
	)
	_ = cmd.Config.BindPFlag("disable_colors", cmd.PersistentFlags().Lookup("disable-colors"))
}
