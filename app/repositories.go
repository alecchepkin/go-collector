package app

import (
	"collector/repositories"
	"github.com/sirupsen/logrus"
)

func CabinetRepo(log *logrus.Entry) *repositories.CabinetRepo {
	r := repositories.NewCabinetRepo(Db(), log)

	return r
}
