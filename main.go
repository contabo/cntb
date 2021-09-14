/*
Copyright © 2021 Contabo GmbH <support@contabo.de>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
