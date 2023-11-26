/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// decrCmd represents the decr command
var decrCmd = &cobra.Command{
	Use:   "decr",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		readFile, err := os.Open(".data.txt")

		defer readFile.Close()
		if err != nil {
			fmt.Println(err)
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)

		var counter string
		for fileScanner.Scan() {
			counter = fileScanner.Text()
			break
		}

		split := strings.Split(counter, "=")
		name := split[0]
		count := split[1]
		atoi, err := strconv.Atoi(count)
		if err != nil {
			log.Fatal(err)
		}
		newcount := atoi - 1
		err = os.WriteFile(".data.txt", []byte(fmt.Sprintf("%s=%d", name, newcount)), 0644)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(decrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
