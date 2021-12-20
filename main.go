package main

import (
	"contabo.com/cli/cntb/cmd"
	_ "contabo.com/cli/cntb/cmd/images"
	_ "contabo.com/cli/cntb/cmd/instanceActions"
	_ "contabo.com/cli/cntb/cmd/instances"
	_ "contabo.com/cli/cntb/cmd/roles"
	_ "contabo.com/cli/cntb/cmd/secrets"
	_ "contabo.com/cli/cntb/cmd/snapshots"
	_ "contabo.com/cli/cntb/cmd/tagAssignment"
	_ "contabo.com/cli/cntb/cmd/tags"
	_ "contabo.com/cli/cntb/cmd/users"
)

func main() {
	cmd.Execute()
}
