package hostconfiguration

// DomainConfiguration represents an nginx domain to configure
// in the sites-enabled file
type DomainConfiguration struct {
	Domain string
	Port   string
}
