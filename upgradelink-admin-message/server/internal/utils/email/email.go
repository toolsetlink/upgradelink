// Package email implements utility routines for email convert
package email

import "upgradelink-admin-message/server/internal/enum/emailauthtype"

func ConvertAuthTypeToString(data uint8) string {
	switch data {
	case emailauthtype.Plain:
		return "plain"
	case emailauthtype.CRAMMD5:
		return "CRAMMD5"
	}

	return "plain"
}
