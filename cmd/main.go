package main

import (
	"fmt"
	"github.com/gost-c/gost-cli/colors"
	"github.com/gost-c/gost-cli/utils"
	"github.com/zcong1993/tools/md5"
	"github.com/zcong1993/tools/resolver"
	"github.com/zcong1993/tools/ulid"
	utils2 "github.com/zcong1993/tools/utils"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	// Version is cli version
	Version = "v0.1.0"
)

var (
	app        = kingpin.New("tools", "Command line tool for tools.")
	dnsCommand = app.Command("dns", "sub command for dns. ")
	dnsHost    = dnsCommand.Arg("host", "host for dns resolver. ").Required().String()

	ulidCommand = app.Command("ulid", "sub command for ulid. ")
	ulidNums    = ulidCommand.Arg("nums", "nums for ulid command. ").Required().Int64()

	md5Command = app.Command("md5", "sub command for md5. ")
	md5Input   = md5Command.Arg("input", "input string for md5 command. ").Required().String()

	version = app.Command("version", "Show tools cli version.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case version.FullCommand():
		showVersion()
	case dnsCommand.FullCommand():
		utils2.Output(resolver.Resolve(*dnsHost))
	case ulidCommand.FullCommand():
		nums := *ulidNums
		if *ulidNums > 100 {
			nums = 100
		}
		utils2.Output(ulid.GenUlids(nums), nil)
	case md5Command.FullCommand():
		utils2.Output(md5.GetMd5([]byte(*md5Input)), nil)
	}
}

func showVersion() {
	version := fmt.Sprintf("%s version %s", colors.Cyan(app.Name), colors.Purple(Version))
	utils.LogPad(version)
}
