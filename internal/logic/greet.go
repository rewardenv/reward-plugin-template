package logic

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func (c *Client) RunCmdGreet(cmd *cobra.Command, args []string) error {
	log.Debugln("Creating a greeting...")

	if len(args) == 0 {
		return fmt.Errorf("no name provided")
	}

	if c.Config.GetBool("add_cakes") {
		log.Debugln("The command line flag was set to true. Adding cakes to the greeting.")
		//nolint:forbidigo
		fmt.Printf("Hello %s, here are some cakes 🎂!\n", strings.Join(args, " "))
	} else {
		//nolint:forbidigo
		fmt.Printf("Hello %s!\n", strings.Join(args, " "))
	}

	log.Debugln("...greeting created.")

	return nil
}
