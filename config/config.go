package config

import "github.com/spf13/viper"

func InitConfig() {
	// Configuração de inicialização, leitura de arquivos de configuração, definição de variáveis de ambiente, etc.
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}