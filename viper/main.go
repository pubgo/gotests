package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func main() {
	os.Setenv("oss_hello1111", "hello")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("oss")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", ".", "-", ".", "/", "."))
	viper.BindEnv()

	fmt.Println(viper.ReadInConfig())
	fmt.Println(viper.Get("oss.hello"))
	fmt.Println(viper.Get("hello"))

	fmt.Println(viper.AllSettings())
}
