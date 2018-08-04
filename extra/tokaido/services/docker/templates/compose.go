package dockertmpl

// DrupalSettings ...
func DrupalSettings(drupalRoot string, projectName string) []byte {
	return []byte(`services:
  fpm:
    environment:
      DRUPAL_ROOT: ` + drupalRoot + `
  nginx:
    environment:
      DRUPAL_ROOT: ` + drupalRoot + `
  drush:
    environment:
      DRUPAL_ROOT: ` + drupalRoot + `
      PROJECT_NAME: ` + projectName + `
  kishu:
      environment:
        DRUPAL_ROOT: ` + drupalRoot)
}

// EdgeContainers ...
func EdgeContainers() []byte {
	return []byte(`services:
  syslog:
    image: tokaido/syslog:edge
  haproxy:
    image: tokaido/haproxy:edge
  varnish:
    image: tokaido/varnish:edge
  nginx:
    image: tokaido/nginx:edge
  fpm:
    image: tokaido/fpm:edge
  drush:
    image: tokaido/drush-heavy:edge`)
}

// EnableSolr ...
func EnableSolr(version string) []byte {
	return []byte(`services:
  solr:
    image: tokaido/solr:` + version + `
    ports:
      - "8983"
    entrypoint:
      - solr-precreate
      - drupal
      - /opt/solr/server/solr/configsets/search-api-solr/`)
}

// SetUnisonVersion ...
func SetUnisonVersion(version string) []byte {
	return []byte(`services:
  unison:
    image: tokaido/unison:` + version)
}

// EnableMemcache ...
func EnableMemcache(version string) []byte {
	return []byte(`services:
  memcache:
    image: memcached:` + version)
}

// WindowsAjustments ...
func WindowsAjustments() []byte {
	return []byte(`services:
  unison:
    image: onnimonni/unison:2.48.4`)
}

// ModWarning - Displayed at the top of `docker-compose.tok.yml`
var ModWarning = []byte(`
# WARNING: THIS FILE IS MANAGED DIRECTLY BY TOKAIDO.
# DO NOT MAKE MODIFICATIONS HERE, THEY WILL BE OVERWRITTEN

`)

// ComposeTokDefaults - Template byte array for `docker-compose.tok.yml`
var ComposeTokDefaults = []byte(`
version: "2"
services:
  unison:
    image: tokaido/unison:2.51.2
    environment:
      - UNISON_DIR=/tokaido/site
      - UNISON_UID=1001
      - UNISON_GID=1001
    ports:
      - "5000"
    volumes:
      - /tokaido/site
  syslog:
    image: tokaido/syslog:latest
    volumes:
      - /tokaido/logs
  haproxy:
    user: "1005"
    image: tokaido/haproxy:latest
    ports:
      - "8080"
      - "8443"
    depends_on:
      - varnish
      - nginx
  varnish:
    user: "1004"
    image: tokaido/varnish:latest
    depends_on:
      - nginx
    volumes_from:
      - syslog
  nginx:
    user: "1002"
    image: tokaido/nginx:latest
    volumes_from:
      - unison
      - syslog
    depends_on:
      - fpm
    ports:
      - "8082"
    environment:
      DRUPAL_ROOT: docroot
  fpm:
    user: "1001"
    image: tokaido/fpm:latest
    working_dir: /tokaido/site/
    volumes_from:
      - unison
      - syslog
    depends_on:
      - syslog
    ports:
      - "9000"
    environment:
      PHP_DISPLAY_ERRORS: "yes"
  mysql:
    image: mysql:5.7
    volumes_from:
      - syslog
    ports:
      - "3306"
    environment:
      MYSQL_DATABASE: tokaido
      MYSQL_USER: tokaido
      MYSQL_PASSWORD: tokaido
      MYSQL_ROOT_PASSWORD: tokaido
  drush:
    image: tokaido/drush-heavy:latest
    hostname: 'tokaido'
    ports:
      - "22"
    working_dir: /tokaido/site
    volumes_from:
      - unison
      - syslog
    environment:
      SSH_AUTH_SOCK: /ssh/auth/sock
      APP_ENV: local
      PROJECT_NAME: tokaido
  kishu:
    image: tokaido/kishu:latest
    volumes_from:
      - unison
    environment:
      DRUPAL_ROOT: docroot
`)
