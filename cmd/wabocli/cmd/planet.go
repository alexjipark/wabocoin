package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var GreetPlanetGmd = &cobra.Command{
	Use: "mars",
	Run: Greets,

}

func Greets (cmd *cobra.Command, args []string) {
	fmt.Println("Hello Mars :)")
}