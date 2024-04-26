package environments

import (
	"github.com/openshift-online/rh-trex/pkg/dao"
	"github.com/openshift-online/rh-trex/pkg/services"
)

type KindNameServiceLocator func() services.KindNameService

func NewKindNameServiceLocator(env *Env) KindNameServiceLocator {
	return func() services.KindNameService {
		return services.NewKindNameService(
			dao.NewKindNameDao(&env.Database.SessionFactory),
			env.Services.Events(),
		)
	}
}
