package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"flag"
)

type BrainConfig struct {
	Redis struct {
		RedisHost string `yaml:"redisHost"`
		RedisPort int    `yaml:"redisPort"`
		RedisPwd  string `yaml:"redisPwd"`
		RedisDB   int    `yaml:"redisDB"`
	}
	Kafka struct {
		KafkaBrokers []string `yaml:"kafkaBrokers,flow"`
	}
	Mysql struct{
		DbUser string `yaml:dbUser`
		DbPasswd string `yaml:dbPasswd`
		DbIp string `yaml:dbIp`
		DbPort int `yaml:dbPort`
		DbName string `yaml dbName`
	}
	DataPath string `yaml:"dataPath"`
	LogPath  string `yaml:"logPath"`
}

var c BrainConfig
var configFile string

func init() {
	flag.StringVar(&configFile, "c", "./etc/brain_config.yaml", "set configuration file")
	flag.Parse()
	c = BrainConfig{}
	b, e := ioutil.ReadFile(configFile)
	if e != nil {
		panic(e)
	}
	err := yaml.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
}

func Get() BrainConfig {
	return c
}
