// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"github.com/greper/ulogin/internal/dao/internal"
)

// internalOauthClientDao is an internal type for wrapping the internal DAO implementation.
type internalOauthClientDao = *internal.OauthClientDao

// oauthClientDao is the data access object for the table oauth_client.
// You can define custom methods on it to extend its functionality as needed.
type oauthClientDao struct {
	internalOauthClientDao
}

var (
	// OauthClient is a globally accessible object for table oauth_client operations.
	OauthClient = oauthClientDao{
		internal.NewOauthClientDao(),
	}
)

// Add your custom methods and functionality below.
