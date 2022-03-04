package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Command struct {
	Command  string
	Desc     string
	Aliases  []string
	Exec     func(ctx context.Context, flags *pflag.FlagSet, args []string) error
	Timeout  int
	SetFlags func(cmd *cobra.Command)
	Args     cobra.PositionalArgs
}

func (c *Command) ToCobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     c.Command,
		Short:   c.Desc,
		Aliases: c.Aliases,
		Args:    c.Args,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
			defer cancel()

			errCh := make(chan error)

			go func() {
				errCh <- c.Exec(ctx, cmd.Flags(), args)
			}()

			select {
			case <-ctx.Done():
				return fmt.Errorf("unknown error")
			case err := <-errCh:
				return err
			}
		},
	}
	if c.SetFlags != nil {
		c.SetFlags(cmd)
	}
	return cmd
}

func (c *Command) Apply(conf *config.Config) *Command {
	c.Timeout = conf.General.Timeout
	return c
}
