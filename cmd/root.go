//
// Copyright (c) 2021 Stefaan Coussement - MIT License
// more info: https://github.com/stefaanc/modelctl/LICENSE
//
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    homedir "github.com/mitchellh/go-homedir"
    "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
    Use:   "modelctl",
    Short: "Building a model for infrastructure deployment and maintenance",
    Long: `Building a model for infrastructure deployment and maintenance`,
    Version: "0.0.1",
}

var cfgFile string

func init() {
    cobra.OnInitialize(initConfig)

    rootCmd.SetVersionTemplate(`{{.Name}} v{{.Version}}`)

    rootCmd.SetHelpTemplate(`
{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}{{if or .Runnable .HasSubCommands}}
{{.UsageString}}{{end}}
`)

    rootCmd.SetUsageTemplate(`
Usage: {{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.modelctl.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile != "" {
        // Use config file from the flag.
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory.
        home, err := homedir.Dir()
        cobra.CheckErr(err)

        // Search config in home directory with name ".modelctl" (without extension).
        viper.AddConfigPath(home)
        viper.SetConfigName(".modelctl")
    }

    viper.AutomaticEnv() // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
    }
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}
