package bootstrap

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/database"
	"github.com/goravel/framework/log"
	"github.com/goravel/sqlserver"
)

func Providers() []foundation.ServiceProvider {
	return []foundation.ServiceProvider{
		&log.ServiceProvider{},
		&database.ServiceProvider{},
		&sqlserver.ServiceProvider{},
	}
}
