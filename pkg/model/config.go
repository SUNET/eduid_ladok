package model

// Cfg is the main configuration structure for this application
type Cfg struct {
	APIServer struct {
		Host string `yaml:"host" validate:"required"`
	} `yaml:"api_server"`

	Production bool   `yaml:"production"`
	HTTPProxy  string `yaml:"http_proxy"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`

	Schools map[string]struct {
		SwamidName string `yaml:"swamid_name" validate:"required"`
	} `yaml:"schools"`

	// SchoolInformation holds information of schools
	SchoolInformation map[string]SchoolInfo `yaml:"school_information"`

	Ladok struct {
		URL         string `yaml:"url"`
		Certificate struct {
			Folder string `yaml:"folder"`
		} `yaml:"certificate"`
		Atom struct {
			Periodicity int `yaml:"periodicity"`
		} `yaml:"atom"`
	} `yaml:"ladok"`

	EduID struct {
		IAM struct {
			URL string `yaml:"url" validate:"required,url"`
		} `yaml:"iam"`
	} `yaml:"eduid"`

	Sunet struct {
		Auth struct {
			URL string `yaml:"url" validate:"required,url"`
		} `yaml:"auth"`
	} `yaml:"sunet"`

	Redis struct {
		DB                  int      `yaml:"db" validate:"required"`
		Addr                string   `yaml:"host" validate:"required_without_all=SentinelHosts SentinelServiceName"`
		SentinelHosts       []string `yaml:"sentinel_hosts" validate:"required_without=Addr,omitempty,min=2"`
		SentinelServiceName string   `yaml:"sentinel_service_name" validate:"required_with=SentinelHosts"`
	} `yaml:"redis"`
}

// Config represent the complete config file structure
type Config struct {
	EduID struct {
		Worker struct {
			Ladok Cfg `yaml:"ladok"`
		} `yaml:"worker"`
	} `yaml:"eduid"`
}
