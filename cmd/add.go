/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var num1 int32
var num2 int32

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A command you can use to perform addition of 2 numbers :)",
	Long: `This command help you get the addition of 2 numbers. It expects the
numbers to passed as value to the flags num1 and num2.
Below is an example of how you can use this command: 

- kubectl plugin add --num1=3 --num2=4

The default value of the flags if not provided is '0'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		fmt.Println("Performing the addition of the following numbers: ")
		fmt.Printf("Num1: %v\n", num1)
		fmt.Printf("Num2: %v\n", num2)
		fmt.Printf("Addition of those 2 numbers is: %v\n", num1+num2)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().Int32Var(&num1, "num1", 0, "First digit of the addition operation")
	addCmd.Flags().Int32Var(&num2, "num2", 0, "Second digit of the addition operation")
}
