package nginx

import (
	"deployed/hostconfiguration"
	"fmt"
)

func buildFileForwardSection() string {
	return `server {
    server_name brad.coffee;

    root /var/www/brad.coffee/html;

    index index.html index.htm;

    location / {
        try_files $uri $uri/ =404;
    }
}
`
}

func buildPortForwardSection(domain string, port string) string {
	return fmt.Sprintf(
		`server {
    server_name %s;
    location / {
        proxy_pass http://localhost:%s;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
    }
}
`, domain, port)
}

// BuildSitesEnabled returns the nginx config file
// to write to /etc/nginx/sites-enabled
func BuildSitesEnabled(domains []hostconfiguration.DomainConfiguration) string {
	fileContents := buildFileForwardSection()
	for _, domainConfig := range domains {
		if domainConfig.Port == "" {
			continue
		}
		fileContents += buildPortForwardSection(domainConfig.Domain, domainConfig.Port)
	}
	return fileContents
}
