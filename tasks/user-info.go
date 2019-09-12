package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	"collector/components"
	m "collector/modeles"
	"github.com/sirupsen/logrus"
	"strconv"
)

type userInfoTask struct {
	cab *m.Cabinet
	log *logrus.Entry
	api *components.ApiClient
}

func UserInfoTask(cab *m.Cabinet, log *logrus.Entry) *userInfoTask {
	t := userInfoTask{}
	t.cab = cab
	t.log = log.WithFields(logrus.Fields{"task": "user-info", "cab": cab.CabinetID})
	t.api = app.ApiClient(t.log)
	return &t
}

func (t *userInfoTask) Run() {
	//defer cltk.Toolkit().Recover(t.log)
	temp := *t.cab
	c := &temp

	resp, err := t.api.User(t.cab)
	if err != nil {
		t.log.Errorln(err, "api user error")
		return
	}
	c.CabinetID = resp.ID
	c.AccountID = resp.Account.ID
	c.Username = resp.Username
	c.Name = resp.Firstname + " " + resp.Lastname
	c.Balance, err = strconv.ParseFloat(resp.Account.Balance, 64)
	if err != nil {
		t.log.Panic(err)
	}
	c.Currency = resp.Currency
	c.Email = resp.Email

	apiRows := []cltk.IRow{c}
	dbRows := []cltk.IRow{t.cab}
	t.log.Info("Finish for all cabs")
	insertRows, _ := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
	t.log.Infof("inserted-updated-rows:%d", len(insertRows))
	cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
}
