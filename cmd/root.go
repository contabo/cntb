package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"contabo.com/cli/cntb/config"
	"github.com/minio/minio-go/v7"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile                       string
	DebugLevel                    string
	ApiBaseUrl                    string
	oauth2ClientId                string
	oauth2ClientIdDesignation     = "oauth2-clientid"
	oauth2ClientSecret            string
	oauth2ClientSecretDesignation = "oauth2-client-secret"
	oauth2TokenUrl                string
	oauth2TokenUrlDesignation     = "oauth2-tokenurl"
	oauth2User                    string
	oauth2UserDesignation         = "oauth2-user"
	oauth2Password                string
	oauth2PasswordDesignation     = "oauth2-password"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cntb",
	Short: "cntb is the command-line interface (CLI) for managing resources at Contabo (http://contabo.com)",
	Long: `                             _
                   _        | |
  ____ ___  ____ _| |_ _____| |__   ___
 / ___) _ \|  _ (_   _|____ |  _ \ / _ \
( (__| |_| | | | || |_/ ___ | |_) ) |_| |
 \____)___/|_| |_| \__)_____|____/ \___/

cntb is the command-line interface (CLI) for managing resources at Contabo (http://contabo.com).
It allows to view, create, change, delete various resources. Furthermore it allows you to
automate tasks easily without programming.

Examples:
 * create, delete, control compute resources
 * manage custom images
 * create subusers
 * assign tags
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setDebugLevel()
	},
	SilenceUsage:  true, // prevent showing usage when showing "Did you mean this?"
	SilenceErrors: true, // preventing showing "Did you mean this?" twice
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	SuggestionsMinimumDistance: 1,
	Args: func(cmd *cobra.Command, args []string) error {
		parsedApiBaseUrl, err := url.ParseRequestURI(ApiBaseUrl)
		if err != nil {
			return errors.New("api is not a well formed URL")
		}
		ApiBaseUrl = parsedApiBaseUrl.String()

		parsedOauth2TokenUrl, err := url.ParseRequestURI(oauth2TokenUrl)
		if err != nil {
			return errors.New("oauth2-tokenurl is not a well formed URL")
		}
		oauth2TokenUrl = parsedOauth2TokenUrl.String()

		return CheckSuggestions(args, cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	minio.MaxRetry = 50
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (Looks up /etc/cntb/.cntb.yaml then $HOME/.cntb.yaml)")

	rootCmd.PersistentFlags().StringVarP(&DebugLevel, "debug", "d", "warn",
		"debug level [fatal|error|warn|info|debug|trace]")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().StringVar(&ApiBaseUrl, "api", "https://api.contabo.com",
		"API base url to be used")
	viper.BindPFlag("api", rootCmd.PersistentFlags().Lookup("api"))

	rootCmd.PersistentFlags().StringVar(&oauth2ClientId, oauth2ClientIdDesignation, "",
		"OAuth2 ClientId. Please refer to you customer panel to retrieve this information")
	viper.BindPFlag(oauth2ClientIdDesignation, rootCmd.PersistentFlags().Lookup(oauth2ClientIdDesignation))

	rootCmd.PersistentFlags().StringVar(&oauth2ClientSecret, oauth2ClientSecretDesignation, "",
		"OAuth2 Secret. Please refer to you customer panel to retrieve this information")
	viper.BindPFlag(oauth2ClientSecretDesignation, rootCmd.PersistentFlags().Lookup(oauth2ClientSecretDesignation))

	rootCmd.PersistentFlags().StringVar(&oauth2TokenUrl, oauth2TokenUrlDesignation,
		"https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token",
		"OAuth2 Token URL. Please refer to you customer panel to retrieve this information")
	viper.BindPFlag(oauth2TokenUrlDesignation, rootCmd.PersistentFlags().Lookup(oauth2TokenUrlDesignation))

	rootCmd.PersistentFlags().StringVar(&oauth2User, oauth2UserDesignation, "",
		"OAuth2 User. Please refer to you customer panel to retrieve this information")
	viper.BindPFlag(oauth2UserDesignation, rootCmd.PersistentFlags().Lookup(oauth2UserDesignation))

	rootCmd.PersistentFlags().StringVar(&oauth2Password, oauth2PasswordDesignation, "",
		"OAuth2 User Password. Please refer to you customer panel to retrieve this information")
	viper.BindPFlag(oauth2PasswordDesignation, rootCmd.PersistentFlags().Lookup(oauth2PasswordDesignation))

	rootCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "normal",
		`output format could be json|yaml|normal(=delimiter)|wide(=delimiter)|jsonpath=...|
	See jsonpath [http://goessner.net/articles/JsonPath/]. `+
			`Delimiter defaults to horizontally aligned spaces, you could also use ',' `+
			`for csv format.`)
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("cntb")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		cfgFile = home + "/.cntb.yaml"
		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath("/etc/cntb/")
		viper.AddConfigPath(home)
		viper.SetConfigName(".cntb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	log.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          false,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		DebugLevel = viper.GetString("debug")
		setDebugLevel()
		log.Info(fmt.Sprintf("Using config file %v", viper.ConfigFileUsed()))
		cfgFile = viper.ConfigFileUsed()
	}
	// taking over value from config or env variables
	DebugLevel = viper.GetString("debug")
	ApiBaseUrl = viper.GetString("api")
	oauth2ClientId = viper.GetString(oauth2ClientIdDesignation)
	oauth2ClientSecret = viper.GetString(oauth2ClientSecretDesignation)
	oauth2TokenUrl = viper.GetString(oauth2TokenUrlDesignation)
	oauth2User = viper.GetString(oauth2UserDesignation)
	oauth2Password = viper.GetString(oauth2PasswordDesignation)

	config.Configure(oauth2TokenUrl, oauth2ClientId, oauth2ClientSecret, oauth2User, oauth2Password, ApiBaseUrl)
}

func setDebugLevel() {
	level, err := log.ParseLevel(DebugLevel)
	if err != nil {
		log.Error("Value for --debug not allowed, please choose a valid one. Use --help to get list.")
		os.Exit(1)
	}
	log.SetLevel(level)
	log.Debug(fmt.Sprintf("Setting debug level to %v", DebugLevel))
}

func CheckSuggestions(args []string, cmd *cobra.Command) error {
	suggestionsString := ""
	if len(args) == 0 {
		return nil
	}
	arg := args[0]
	if suggestions := cmd.SuggestionsFor(arg); len(suggestions) > 0 {
		suggestionsString += "\n\nDid you mean this?\n"
		for _, s := range suggestions {
			suggestionsString += fmt.Sprintf("\t%v\n", s)
		}
	}
	if suggestionsString != "" {
		return fmt.Errorf("invalid argument %q for %q%s", arg, cmd.CommandPath(), suggestionsString)
	}
	return nil
}
