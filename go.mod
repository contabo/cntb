module contabo.com/cli/cntb

go 1.16

require (
	contabo.com/cli/cntb/openapi v0.0.0-00010101000000-000000000000
	github.com/PaesslerAG/gval v1.1.2 // indirect
	github.com/PaesslerAG/jsonpath v0.1.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hprose/hprose-go v0.0.0-20161031134501-83de97da5004
	github.com/minio/minio-go/v7 v7.0.23
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	gopkg.in/yaml.v2 v2.4.0
	sigs.k8s.io/yaml v1.3.0
)

replace contabo.com/cli/cntb/openapi => ./openapi
