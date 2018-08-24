package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"path/filepath"
	"errors"
)

const desc = `
Automatic deploy war to tomcat at specific timeFilter out template files
	$ deploy myapp.war --at "2018-08-24 15:01" --tomcat "/path/to/tomcat"
`

func main() {
	c := deployCmd{}

	cmd := &cobra.Command{
		Use:   "deploy [flags] WAR",
		Short: fmt.Sprintf("automatic deploy war to tomcat at specific time"),
		Long:  desc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("war is required")
			}
			// verify war path exists
			if _, err := os.Stat(args[0]); err == nil {
				if c.warPath, err = filepath.Abs(args[0]); err != nil {
					return err
				}
			} else {
				return err
			}

			return c.run()
		},
	}
	f := cmd.Flags()
	f.StringVarP(&c.at, "at", "", "", "specific time")
	f.StringVarP(&c.layout, "layout", "", "2006-01-02 15:04", "specific time layout to parse")
	f.StringVarP(&c.tomcatPath, "tomcat", "t", "", "specific tomcat path")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
