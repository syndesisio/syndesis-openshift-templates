// To install, run:
//
//   go get -u github.com/spf13/cobra github.com/spf13/pflag github.com/hoisie/mustache
//   go install ipaas-template.go
//
// Or to just test locally:
//
//     go run ipaas-template.go <args>
//
package main

import (
	//"flag"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/hoisie/mustache"
	"io/ioutil"
	"flag"
)

func main() {
	err := installCommand.Execute()
	check(err)
}

var installCommand = &cobra.Command{
	Use:   "ipaas-install",
	Short: "ipaas-install is a tool for creating the OpenShift resources needed to install the redhat-ipass.",
	Long:  `ipaas-install is a tool for creating the OpenShift resources needed to install the redhat-ipass.`,
	Run: install,
}

type Context struct {
	ImageStreams bool;
	ClusterAdmin bool;
}

var context = Context{}

func init() {
	installCommand.PersistentFlags().BoolVar(&context.ImageStreams, "image-streams", true, "Should the template use image streams?")
	installCommand.PersistentFlags().BoolVar(&context.ClusterAdmin, "cluster-admin", true, "Should the templaae crate cluster resources?")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

func install(cmd *cobra.Command, args []string) {
	template, err := ioutil.ReadFile("redhat-ipaas.yml.mustache")
	check(err)
	fmt.Print(mustache.Render(string(template), context))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}