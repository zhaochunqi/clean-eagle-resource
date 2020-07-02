package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "Clean Eagle Resources"
	app.Usage = "Sometimes eagle resources not match *.json in folder, this tools will help fix it."
	app.Authors = []*cli.Author{{
		"ZHAOCHUNQI", "zcq.qiqi@gmail.com",
	}}
	app.Version = "1.0.0"
}

func commands() {
}


func options(){
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name: "dryrun",
			Usage: "Running programme with dryrun.",
			Value: false,
			DefaultText: "false",
		},
		&cli.StringFlag{
			Name: "file",
			Aliases: []string{"f"},
			Required: true,
			Value: "file",
			Usage: "File path this commands will run." ,
		},
	}
}

func actions() {
	app.Action = func(context *cli.Context) error {
		var output string

		dryRunMode := context.Bool("dryrun")

		if dryRunMode{
			output = "Working in DryRun mode"
			fmt.Println(output)
		}


		if len(context.String("file"))!=0 {
			path:= context.String("file")

			FixImagePath(path, dryRunMode)
		}


		return nil
	}
}

func main() {
	info()
	commands()
	options()
	actions()

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}


