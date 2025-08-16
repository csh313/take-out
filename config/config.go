package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const configFile = "config/application-dev.yaml"

func InitConfig() *Config {
	viper.SetConfigName("application-dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper.ReadInConfig() error(%v)", err)
	}
	AppConfig := &Config{}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("viper.Unmarshal() error(%v)", err)
	}
	fmt.Println(AppConfig)
	return AppConfig
}

type Config struct {
	Server     Server
	Datasource Datasource
	Redis      Redis
	Log        Log
	Jwt        Jwt
	AliOss     AliOss
	Wechat     Wechat
	Kafka      Kafka
	MysqlConf  MysqlConf `mapstructure:"mysql_conf"`
}

type Server struct {
	Port  string
	Level string
}

type Datasource struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string `mapstructure:"db_name"`
	Config   string
}

func (d *Datasource) Dsn() string {
	return d.Username + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DBName + "?" + d.Config
}

type Redis struct {
	Password string
	Host     string
	Port     string
	Database int
}
type Log struct {
	Level    string
	FilePath string
}

type JwtOption struct {
	Secret string
	TTL    int
	Name   string
}

type Jwt struct {
	Admin JwtOption
	User  JwtOption
}

type AliOss struct {
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	Endpoint        string
	BucketName      string `mapstructure:"bucket_name"`
	Region          string `yaml:"region" json:"region"`
}

type Wechat struct {
	AppId  string `mapstructure:"appid"`
	Secret string `mapstructure:"secret"`
}

type Kafka struct {
	Brokers []string `mapstructure:"brokers"`
	Topic   string   `mapstructure:"topic"`
	GroupId string   `mapstructure:"group_id"`
}

type SqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string `mapstructure:"db_name"`
	Config   string
}

type Strategy struct {
	Read  string
	Write string
}
type MysqlConf struct {
	Master   SqlConfig
	Slave    []SqlConfig
	Strategy Strategy
}

func (s *SqlConfig) Dsn() string {
	return s.Username + ":" + s.Password + "@tcp(" + s.Host + ":" + s.Port + ")/" + s.DBName + "?" + s.Config
}
