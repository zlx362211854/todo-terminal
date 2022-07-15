package ctx

import (
	"terminal/service/auth"
	"terminal/service/data"
)

type Ctx struct {
	Auth *auth.Auth
	Data *data.TODOData
}
