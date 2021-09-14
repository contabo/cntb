package config

type Configuration struct {
	Debug              string `yaml:"debug"`
	Oauth2TokenUrl     string `yaml:"oauth2-tokenurl"`
	Oauth2ClientId     string `yaml:"oauth2-clientid"`
	Oauth2ClientSecret string `yaml:"oauth2-client-secret"`
	Oauth2User         string `yaml:"oauth2-user"`
	Oauth2Password     string `yaml:"oauth2-password"`
	Api                string `yaml:"api"`
}

var Conf Configuration = Configuration{}

func Configure(oauth2TokenUrl, oauth2ClientId, oauth2ClientSecret, oauth2User, oauth2Password, apiUrl string) {
	Conf.Oauth2TokenUrl = oauth2TokenUrl
	Conf.Oauth2ClientId = oauth2ClientId
	Conf.Oauth2ClientSecret = oauth2ClientSecret
	Conf.Oauth2User = oauth2User
	Conf.Oauth2Password = oauth2Password
	Conf.Api = apiUrl
}
