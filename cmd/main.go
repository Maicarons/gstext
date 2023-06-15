package main

import (
	"github.com/Maicarons/gstext"
	"github.com/Maicarons/gstext/cmd/ui"
	"log"
	"os"
)

var Version = gstext.Version

var Outfile []byte
var Ostr gstext.GameText

func main() {
	if len(os.Args) < 2 || os.Args[1] == "help" {
		log.Fatalln(`You must specify one or more arguments.

example:
gstext json s.json to xml s.xml`)
	}
	if len(os.Args) >= 3 {
		if os.Args[3] == "to" {
			Out()
			return
		}

	}
	if len(os.Args) >= 3 {
		if os.Args[3] == "convert" {
			return
		}

	}
	if len(os.Args) >= 1 {
		if os.Args[1] == "version" {
			println("gstext v", Version)
			return
		}

	}
	if len(os.Args) >= 1 {
		if os.Args[1] == "ui" {
			ui.Start()
			return
		}

	}
	log.Fatalln(`You must specify one or more arguments.

example:
gstext json s.json to xml s.xml`)

}

func Out() {
	if len(os.Args) < 6 || os.Args[2] == "help" {
		log.Fatalln(`You must specify one or more arguments.
Format:
gstext [json/xml/yaml/toml] filename to [json/xml/yaml/toml] filename 
    convert your script to other format

example:
gstext json s.json to xml s.xml`)
	}

	OrgFile, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatalln(err.Error())
	}

	if os.Args[1] == "json" {
		var err error
		Ostr, err = gstext.JSONUnmarshalGameText(OrgFile)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else if os.Args[1] == "xml" {
		var err error
		Ostr, err = gstext.XMLUnmarshalGameText(OrgFile)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else if os.Args[1] == "toml" {
		var err error
		Ostr, err = gstext.TOMLUnmarshalGameText(OrgFile)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else if os.Args[1] == "yaml" {
		var err error
		Ostr, err = gstext.YAMLUnmarshalGameText(OrgFile)
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		log.Fatalln("Error input format\nSee \"help\" for more details.")
	}

	if os.Args[4] == "json" {
		var err error
		Outfile, err = Ostr.JSONMarshal()
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else if os.Args[4] == "yaml" {
		var err error
		Outfile, err = Ostr.YAMLMarshal()
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else if os.Args[4] == "xml" {
		var err error
		Outfile, err = Ostr.XMLMarshal()
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else if os.Args[4] == "toml" {
		var err error
		Outfile, err = Ostr.TOMLMarshal()
		if err != nil {
			log.Fatalln(err.Error())
		}
	} else {
		log.Fatalln("Error output format\nSee \"help\" for more details.")
	}

	err = os.WriteFile(os.Args[5], Outfile, 0777)
	if err != nil {
		log.Fatalln(err.Error())
	}
	println("Success")
}
