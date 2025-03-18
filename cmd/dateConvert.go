/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var dateConvertCmd = &cobra.Command{
	Use:   "date-convert",
	Short: "Convert a list of dates from stdin in a specified date format and spits to stdout in another format",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		src, err := cmd.Flags().GetString("source-format")
		if err != nil {
			panic(err)
		}

		dst, err := cmd.Flags().GetString("destination-format")
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			v := scanner.Text()
			t, err := time.Parse(src, v)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println(t.Format(dst))
		}
	},
}

func init() {
	rootCmd.AddCommand(dateConvertCmd)

	dateConvertCmd.Flags().StringP("source-format", "s", "2006/01/02 15:04:05", "Datetime format for the input data")
	dateConvertCmd.Flags().StringP("destination-format", "d", time.RFC3339, "Datetime format for the output data")
}
