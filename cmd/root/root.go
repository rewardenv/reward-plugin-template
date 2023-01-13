package root

import (
	"fmt"

	"github.com/spf13/cobra"

	cmdpkg "github.com/rewardenv/reward-greeter/cmd"
	"github.com/rewardenv/reward-greeter/cmd/greet"
	"github.com/rewardenv/reward-greeter/internal/config"
	"github.com/rewardenv/reward-greeter/internal/logic"
)

func NewCmdRoot(c *config.Config) *cmdpkg.Command {
	var cmd = &cmdpkg.Command{
		Command: &cobra.Command{
			Use:   fmt.Sprintf("greeter [command]"),
			Short: fmt.Sprintf("greeter is a cli tool which helps you to create fantastic greetings"),
			Long: `
██████╗ ███████╗██╗    ██╗ █████╗ ██████╗ ██████╗      ██████╗ ██████╗ ███████╗███████╗████████╗███████╗██████╗ 
██╔══██╗██╔════╝██║    ██║██╔══██╗██╔══██╗██╔══██╗    ██╔════╝ ██╔══██╗██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██████╔╝█████╗  ██║ █╗ ██║███████║██████╔╝██║  ██║    ██║  ███╗██████╔╝█████╗  █████╗     ██║   █████╗  ██████╔╝
██╔══██╗██╔══╝  ██║███╗██║██╔══██║██╔══██╗██║  ██║    ██║   ██║██╔══██╗██╔══╝  ██╔══╝     ██║   ██╔══╝  ██╔══██╗
██║  ██║███████╗╚███╔███╔╝██║  ██║██║  ██║██████╔╝    ╚██████╔╝██║  ██║███████╗███████╗   ██║   ███████╗██║  ██║
╚═╝  ╚═╝╚══════╝ ╚══╝╚══╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝      ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝
                                                                                                                
`,
			Version: c.AppVersion(),
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) (
				[]string, cobra.ShellCompDirective,
			) {
				return nil, cobra.ShellCompDirectiveNoFileComp
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return logic.New(c).RunCmdRoot(cmd)
			},
		},
		Config: c,
	}

	addFlags(cmd)
	c.Init()

	cmd.AddCommands(greet.NewCmdGreet(c))

	return cmd
}

func addFlags(cmd *cmdpkg.Command) {
	// --log-level
	cmd.PersistentFlags().String(
		"log-level", "info", "logging level (options: trace, debug, info, warning, error)",
	)
	cmd.Config.BindPFlag("log_level", cmd.PersistentFlags().Lookup("log-level"))

	// --debug
	cmd.PersistentFlags().Bool(
		"debug", false, "enable debug mode (same as --log-level=debug)",
	)
	cmd.Config.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	// --disable-colors
	cmd.PersistentFlags().Bool(
		"disable-colors", false, "disable colors in output",
	)
	cmd.Config.BindPFlag("disable_colors", cmd.PersistentFlags().Lookup("disable-colors"))
}
