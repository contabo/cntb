module contabo.com/cli/cntb

go 1.16

require (
	contabo.com/cli/cntb/openapi v0.0.0-00010101000000-000000000000
	github.com/PaesslerAG/gval v1.1.0 // indirect
	github.com/PaesslerAG/jsonpath v0.1.1
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hprose/hprose-go v0.0.0-20161031134501-83de97da5004
	github.com/magefile/mage v1.11.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.8.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	gopkg.in/yaml.v2 v2.4.0
	sigs.k8s.io/yaml v1.2.0
)

replace contabo.com/cli/cntb/openapi => ./openapi
