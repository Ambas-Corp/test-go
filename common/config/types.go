package config

//Configuration Struct
type Configuration struct {
	Debug    bool
	Server   Server
	Context  Context
	Database Database
	URL      URL
	Redis    Redis
}

//Database Struct
type Database struct {
	Driver string
	Host   string
	Port   string
	User   string
	Pass   string
	Name   string
}

//Server Struct
type Server struct {
	Port string
	Name string
}

//Context Struct
type Context struct {
	Timeout int
}

//URL struct
type URL struct {
}

//Redis struct
type Redis struct {
	Host string
	Port string
}
