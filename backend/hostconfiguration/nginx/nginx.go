package nginx

import (
	"deployed/datastore"
	"fmt"
)

func buildFileForwardSection(domain string, directory string) string {
	return fmt.Sprintf(`server {
    server_name %s;

    root /var/www/%s;

    index index.html index.htm;

    location / {
        try_files $uri $uri/ /index.html;
    }
}
`, domain, directory)
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
func BuildSitesEnabled(domains []*datastore.DomainConfiguration) string {
	fileContents := ""
	for _, domainConfig := range domains {
		if domainConfig.GetPort() != "" {
			fileContents += buildPortForwardSection(domainConfig.GetDomain(), domainConfig.GetPort())
			continue
		}
		if domainConfig.GetForwardDirectory() != "" {
			fileContents += buildFileForwardSection(domainConfig.GetDomain(), domainConfig.GetForwardDirectory())
			continue
		}
	}
	return fileContents
}
