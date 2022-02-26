package adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/KatsuyaAkasaka/nt/pkg/domain/config"
	"github.com/spf13/cobra"
)

type Command struct {
	Command string
	Desc    string
	Aliases []string
	Exec    func(ctx context.Context, args []string) error
	Timeout int
}

func (c *Command) ToCobraCommand() *cobra.Command {
	return &cobra.Command{
		Use:     c.Command,
		Short:   c.Desc,
		Aliases: c.Aliases,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
			defer cancel()

			errCh := make(chan error)

			go func() {
				errCh <- c.Exec(ctx, args)
			}()

			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return fmt.Errorf("unknown error")
			case err := <-errCh:
				return err
			}
		},
	}
}

func (c *Command) Apply(conf *config.Config) *Command {
	if c.Timeout == 0 {
		c.Timeout = conf.Note_cli.Todo.Timeout
	}
	return c
}
