//
// Copyright (c) 2021 Stefaan Coussement - MIT License
// more info: https://github.com/stefaanc/modelctl/LICENSE
//
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var version string = "0.0.1"

func init() {
    rootCmd.AddCommand(&cobra.Command{
        Use:   "version",
        Short: "Get the version number of modelctl",
        Long:  `Get the version number of modelctl`,
        Run: printVersion,
    })
}

func printVersion(cmd *cobra.Command, args []string) {
    fmt.Printf("modelctl v%s\n", version)
}
