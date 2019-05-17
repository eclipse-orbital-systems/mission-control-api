package echo

import (
	"github.com/eclipse-orbital-systems/mission-control-api/dal"
)

type Api struct {
	dalImpl *dal.Dal
}

func New(dalImpl *dal.Dal) *Api {
	obj := Api{}
	obj.dalImpl = dalImpl

	return &obj
}
