package structs

// Config is a structure containing global website configuration.
//
// See the comments for Server, Database and PageContext for more details.
type Config struct {
	Server      Server      `toml:"server"`
	Database    Database    `toml:"database"`
	PageContext PageContext `toml:"pageContext"`
}

// Server is a structure containing server configuration.
type Server struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
	Timeout int    `toml:"timeout"`
}

type Database struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
}

// PageContext is a structure containing static information to provide
// to all page templates.
//
// This contains the website's long and short names, as well as a directory
// of pages for navigation.
type PageContext struct {
	LongName        string `toml:"longName"`
	ShortName       string `toml:"shortName"`
	SiteDescription string `toml:"siteDescription"`
	URLPrefix       string `toml:"urlPrefix"`
	FullURL         string `toml:"fullURL"`
	MainTwitter     string `toml:"mainTwitter"`
	MainFacebook    string `toml:"mainFacebook"`
}
