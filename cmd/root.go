package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use: "toolbox",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(toolboxArt())
			cmd.Help()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func toolboxArt() string {
	return `
 ________                   __ __                         
|        \                 |  \  \                        
 \▓▓▓▓▓▓▓▓ ______   ______ | ▓▓ ▓▓____   ______  __    __ 
   | ▓▓   /      \ /      \| ▓▓ ▓▓    \ /      \|  \  /  \
   | ▓▓  |  ▓▓▓▓▓▓\  ▓▓▓▓▓▓\ ▓▓ ▓▓▓▓▓▓▓\  ▓▓▓▓▓▓\\▓▓\/  ▓▓
   | ▓▓  | ▓▓  | ▓▓ ▓▓  | ▓▓ ▓▓ ▓▓  | ▓▓ ▓▓  | ▓▓ >▓▓  ▓▓ 
   | ▓▓  | ▓▓__/ ▓▓ ▓▓__/ ▓▓ ▓▓ ▓▓__/ ▓▓ ▓▓__/ ▓▓/  ▓▓▓▓\ 
   | ▓▓   \▓▓    ▓▓\▓▓    ▓▓ ▓▓ ▓▓    ▓▓\▓▓    ▓▓  ▓▓ \▓▓\
    \▓▓    \▓▓▓▓▓▓  \▓▓▓▓▓▓ \▓▓\▓▓▓▓▓▓▓  \▓▓▓▓▓▓ \▓▓   \▓▓`
}
