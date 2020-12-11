package hostconfiguration

// DefaultDomains are the domains that are not managed
// directly by Deployed that must be maintained
var DefaultDomains []DomainConfiguration = []DomainConfiguration{
	DomainConfiguration{
		Domain: "brad.coffee",
	},
	DomainConfiguration{
		Domain: "code.brad.coffee",
		Port:   "2212",
	},
	DomainConfiguration{
		Domain: "dev.brad.coffee",
		Port:   "2212",
	},
	DomainConfiguration{
		Domain: "app.brad.coffee",
		Port:   "4200",
	},
	DomainConfiguration{
		Domain: "api.brad.coffee",
		Port:   "3000",
	},
}
