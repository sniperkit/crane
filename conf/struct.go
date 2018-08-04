package conf

// Config the application's configuration
// IMPORTANT!
// Casing of the `Config` struct properties is important to note
// All properties must be cased as capital letter first, followed by all lowercase
// eg. `Customcompose` (correct), `CustomCompose` (incorrect)
// This is to ensure they both conform to the golang convention
// and that they are able to be properly parsed by the `tok config-x` commands
type Config struct {
	Tokaido struct {
		Force            bool   `yaml:"force,omitempty"`
		Customcompose    bool   `yaml:"customcompose"`
		Debug            bool   `yaml:"debug,omitempty"`
		Config           string `yaml:"config,omitempty"`
		Enableemoji      bool   `yaml:"enableemoji"`
		Betacontainers   bool   `yaml:"betacontainers"`
		Dependencychecks bool   `yaml:"dependencychecks"`
		Project          struct {
			Name string `yaml:"name"`
			Path string `yaml:"path,omitempty"`
		} `yaml:"project"`
	} `yaml:"tokaido"`
	Drupal struct {
		Path         string `yaml:"path,omitempty"`
		Majorversion string `yaml:"majorversion,omitempty"`
	} `yaml:"drupal,omitempty"`
	System struct {
		Xdebug struct {
			Port      string `yaml:"port,omitempty"`
			Logpath   string `yaml:"logpath,omitempty"`
			Enabled   bool   `yaml:"enabled"`
			Autostart bool   `yaml:"autostart"`
		} `yaml:"xdebug,omitempty"`
		Syncsvc struct {
			Systemdpath string `yaml:"systemdpath,omitempty"`
			Launchdpath string `yaml:"launchdpath,omitempty"`
			Enabled     bool   `yaml:"enabled"`
		} `yaml:"syncsvc,omitempty"`
	} `yaml:"system,omitempty"`
	Services Services `yaml:"services,omitempty"`
}

// Services ...
type Services struct {
	Unison struct {
		Image       string   `yaml:"image,omitempty"`
		Hostname    string   `yaml:"hostname,omitempty"`
		Ports       []string `yaml:"ports,omitempty"`
		Entrypoint  []string `yaml:"entrypoint,omitempty"`
		User        string   `yaml:"user,omitempty"`
		Cmd         string   `yaml:"cmd,omitempty"`
		Volumesfrom []string `yaml:"volumes_from,omitempty"`
		Dependson   []string `yaml:"depends_on,omitempty"`
		Environment []string `yaml:"environment,omitempty"`
		Volumes     []string `yaml:"volumes,omitempty"`
	} `yaml:"unison,omitempty"`
	Syslog struct {
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Volumesfrom []string          `yaml:"volumes_from,omitempty"`
		Dependson   []string          `yaml:"depends_on,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
		Volumes     []string          `yaml:"volumes,omitempty"`
	} `yaml:"syslog,omitempty"`
	Haproxy struct {
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Volumesfrom []string          `yaml:"volumes_from,omitempty"`
		Dependson   []string          `yaml:"depends_on,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"haproxy,omitempty"`
	Varnish struct {
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Volumesfrom []string          `yaml:"volumes_from,omitempty"`
		Dependson   []string          `yaml:"depends_on,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"varnish,omitempty"`
	Nginx struct {
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Volumesfrom []string          `yaml:"volumes_from,omitempty"`
		Dependson   []string          `yaml:"depends_on,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"nginx,omitempty"`
	Fpm struct {
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Volumesfrom []string          `yaml:"volumes_from,omitempty"`
		Workingdir  string            `yaml:"working_dir,omitempty"`
		Dependson   []string          `yaml:"depends_on,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"fpm,omitempty"`
	Memcache struct {
		Enabled     bool              `yaml:"enabled,omitempty"`
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"memcache,omitempty"`
	Mysql struct {
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"mysql,omitempty"`
	Drush struct {
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Workingdir  string            `yaml:"working_dir,omitempty"`
		Volumesfrom []string          `yaml:"volumes_from,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"drush,omitempty"`
	Solr struct {
		Enabled     bool              `yaml:"enabled,omitempty"`
		Image       string            `yaml:"image,omitempty"`
		Hostname    string            `yaml:"hostname,omitempty"`
		Ports       []string          `yaml:"ports,omitempty"`
		Entrypoint  []string          `yaml:"entrypoint,omitempty"`
		User        string            `yaml:"user,omitempty"`
		Cmd         string            `yaml:"cmd,omitempty"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"solr,omitempty"`
	Kishu struct {
		Enabled     bool              `yaml:"enabled,omitempty"`
		Image       string            `yaml:"image,omitempty"`
		Volumesfrom []string          `yaml:"volumes_from"`
		Environment map[string]string `yaml:"environment,omitempty"`
	} `yaml:"kishu,omitempty"`
}

// ComposeDotTok ...
type ComposeDotTok struct {
	Version  string   `yaml:"version,omitempty"`
	Services Services `yaml:"services,omitempty"`
}
