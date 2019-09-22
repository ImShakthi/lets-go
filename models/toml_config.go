package models

import "time"

type Config struct {
	Servers map[string]Server
	RBMS    RBMS `toml:"rdbms"`
	//Age  int  `toml:"age"`
}
type RBMS struct {
	Database Database `toml:"DB"`
	User     User     `toml:"user"`
}

type Database struct {
	Name string `toml:"name"`
	Port int    `toml:"port"`
}

type User struct {
	Name     string `toml:"name"`
	Password string `toml:"password"`
}

type TomlConfig struct {
	Title   string
	Owner   OwnerInfo
	DB      DB `toml:"DB"`
	Servers map[string]Server
	Clients Clients
}

type OwnerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type DB struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type Server struct {
	IP string `toml:"ip"`
	DC string `toml:"dc"`
}

type Clients struct {
	Data  [][]interface{}
	Hosts []string
}
