package cmd

import (
	"encoding/json"
	"strconv"

	"contabo.com/cli/cntb/openapi"
	"github.com/spf13/viper"
)

type PrivateNetwork struct {
	PrivateNetworkId int64    `json:"privateNetworkId"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Region           string   `json:"region"`
	DataCenter       string   `json:"dataCenter"`
	Cidr             string   `json:"cidr"`
	AvailableIps     int64    `json:"availableIps"`
	Instances        []string `json:"instances"`
}

func SetPrivateNetwork(
	PrivateNetworkId int64,
	Name string,
	Description string,
	Region string,
	DataCenter string,
	Cidr string,
	AvailableIps int64) *PrivateNetwork {
	privateNetwork := new(PrivateNetwork)
	privateNetwork.PrivateNetworkId = PrivateNetworkId
	privateNetwork.Name = Name
	privateNetwork.Description = Description
	privateNetwork.Region = Region
	privateNetwork.DataCenter = DataCenter
	privateNetwork.Cidr = Cidr
	privateNetwork.AvailableIps = AvailableIps
	privateNetwork.Instances = []string{}
	return privateNetwork
}

func mapInstancesToIds(data []openapi.PrivateNetworkResponse) ([]byte, error) {
	if viper.GetString("output") != "json" && viper.GetString("output") != "yaml" {
		vpn := data[0]
		var privateNetworks []*PrivateNetwork
		privateNetwork := SetPrivateNetwork(vpn.PrivateNetworkId, vpn.Name, vpn.Description, vpn.Region, vpn.DataCenter, vpn.Cidr, vpn.AvailableIps)
		var len = len(vpn.Instances)
		// if instance list has elements
		if len > 0 {
			var instanceIds []string
			for _, instance := range vpn.Instances {
				instanceIds = append(instanceIds, strconv.FormatInt(instance.InstanceId, 10))
			}
			privateNetwork.Instances = instanceIds
		}
		privateNetworks = append(privateNetworks, privateNetwork)
		return json.Marshal(privateNetworks)
	} else {
		return json.Marshal(data)
	}
}
