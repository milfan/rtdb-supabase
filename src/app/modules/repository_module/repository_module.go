package repository_module

import (
	"github.com/milfan/neoten-lib/lib_http_request"
	"github.com/milfan/rtdb-supabase/src/infra/configs/app_config"
	infra_repositories "github.com/milfan/rtdb-supabase/src/infra/repositories"
)

type (
	RepositoryModules struct {
		InfraRepo InfraRepositoryModule
	}

	InfraRepositoryModule struct {
		RTDBRepository infra_repositories.IRTDBRepository
	}
)

func LoadRepositoryModules(
	rtdbConf app_config.RTDBConfig,
	httpRequestUtil lib_http_request.IHttpRequestUtils,
) RepositoryModules {
	return RepositoryModules{
		InfraRepo: InfraRepositoryModule{
			RTDBRepository: infra_repositories.NewRTDBRepository(
				rtdbConf.Host,
				rtdbConf.ClientID,
				rtdbConf.ClientSecret,
				httpRequestUtil,
			),
		},
	}
}
