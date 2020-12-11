package certbot

import (
	"deployed/hostconfiguration"
	"os/exec"
)

// UpdateCertificates runs certbot to support the domains specified
func UpdateCertificates(domains []hostconfiguration.DomainConfiguration) error {
	certbotArgs := []string{"--nginx", "--non-interactive", "--agree-tos", "--expand", "--redirect", "--hsts"}
	for _, domainConfig := range domains {
		certbotArgs = append(certbotArgs, "-d")
		certbotArgs = append(certbotArgs, domainConfig.Domain)
	}
	cmd := exec.Command("certbot", certbotArgs...)
	return cmd.Run()
}
