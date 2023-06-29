package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type CamundaConfig struct {
	ZEEBE_ADDRESS                  string `mapstructure:"ZEEBE_ADDRESS"`
	ZEEBE_CLIENT_ID                string `mapstructure:"ZEEBE_CLIENT_ID"`
	ZEEBE_CLIENT_SECRET            string `mapstructure:"ZEEBE_CLIENT_SECRET"`
	ZEEBE_AUTHORIZATION_SERVER_URL string `mapstructure:"ZEEBE_AUTHORIZATION_SERVER_URL"`
	ZEEBE_TOKEN_AUDIENCE           string `mapstructure:"ZEEBE_TOKEN_AUDIENCE"`
	CAMUNDA_CLUSTER_ID             string `mapstructure:"CAMUNDA_CLUSTER_ID"`
	CAMUNDA_CLUSTER_REGION         string `mapstructure:"CAMUNDA_CLUSTER_REGION"`
	CAMUNDA_CREDENTIALS_SCOPES     string `mapstructure:"CAMUNDA_CREDENTIALS_SCOPES"`
	CAMUNDA_OAUTH_URL              string `mapstructure:"CAMUNDA_OAUTH_URL"`
}

func newCamundaConfig() *CamundaConfig {

	camundaConfig := CamundaConfig{}
	viper.SetConfigFile("camunda.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file camunda.env : ", err)
	}

	err = viper.Unmarshal(&camundaConfig)
	if err != nil {
		log.Fatal("Camunda environment can't be loaded : ", err)
	}

	return &camundaConfig
}
