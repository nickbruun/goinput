package input

import (
	"errors"
	"fmt"
	"golang.org/x/net/idna"
	"regexp"
	"strings"
)

var (
	// Invalid e-mail address.
	ErrInvalidEmail = errors.New("invalid e-mail address")

	// E-mail address user part pattern.
	emailUserPartPattern = regexp.MustCompile(`^[-!#$%&'*+/=?^_` + "`" + `{}|~0-9A-Za-z]+(\.[-!#$%&'*+/=?^_` + "`" + `{}|~0-9A-Za-z]+)*$`)

	// E-mail address domain part pattern.
	emailDomainPartPattern = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+([a-zA-Z0-9\-]{2,63})$`)
)

// Parse e-mail.
//
// Lower-cases the entire e-mail address and normalizes the domain part with
// IDNA encoding. This method does not allow IP address domain parts, nor
// quoted user parts.
func ParseEmail(in string) (out string, err error) {
	// First, attempt to split the e-mail address into user and domain parts.
	atIndex := strings.IndexByte(in, '@')
	if atIndex == -1 {
		err = ErrInvalidEmail
		return
	}

	userPart := strings.TrimSpace(in[:atIndex])
	domainPart := strings.TrimSpace(in[atIndex+1:])

	// Validate the user part.
	if !emailUserPartPattern.MatchString(userPart) {
		err = ErrInvalidEmail
		return
	}

	// Validate the domain part.
	if !emailDomainPartPattern.MatchString(domainPart) {
		// Try for possible IDN domain-part.
		var idnaErr error
		if domainPart, idnaErr = idna.ToASCII(domainPart); idnaErr == nil {
			if !emailDomainPartPattern.MatchString(domainPart) {
				err = ErrInvalidEmail
				return
			}
		} else {
			err = ErrInvalidEmail
			return
		}
	}

	return fmt.Sprintf("%s@%s", strings.ToLower(userPart), strings.ToLower(domainPart)), nil
}
