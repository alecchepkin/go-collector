package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	"collector/components"
	m "collector/modeles"
	"collector/repositories"
	"github.com/sirupsen/logrus"
)

type cabinetsTask struct {
	log     *logrus.Entry
	cabRepo *repositories.CabinetRepo
	proxy   *components.Proxy
	api     *components.ApiClient
}

func CabinetsTask(log *logrus.Entry) *cabinetsTask {
	t := cabinetsTask{}
	t.log = log.WithField("task", "cabinets")
	t.cabRepo = app.CabinetRepo(t.log)
	t.proxy = app.Proxy(t.log)
	t.api = app.ApiClient(t.log)
	return &t
}

func (t *cabinetsTask) Run() {
	defer cltk.Toolkit().Recover(t.log)
	t.addAgenciesCabs()
	t.addExternalAccounts()
}

func (t *cabinetsTask) addAgenciesCabs() {
	defer cltk.Toolkit().Recover(t.log)
	agencies := t.cabRepo.FindAllAgencies()
	t.proxy.UpdateTokens(agencies)
	t.log.Infof("get api-agency-clients")
	for _, agency := range agencies {

		resp, err := t.api.AgencyClients(agency)
		if err != nil {
			t.log.Panic(err)
		}

		dbAccs, dbRows := func() (cabs []*m.Cabinet, rows []cltk.IRow) {
			err := app.Db().Model(&cabs).Where("parent_id=?", agency.CabinetID).Select()
			if err != nil {
				t.log.Panic(err)
			}

			for _, cab := range cabs {
				rows = append(rows, cab)
			}
			return
		}()

		var apiRows []cltk.IRow
		for _, item := range resp.Items {

			acc := &m.Cabinet{}
			for _, dba := range dbAccs {
				if item.User.ID == dba.CabinetID {
					acc = dba
				}
			}
			acc.CabinetID = item.User.ID
			acc.AccountID = item.User.Account.ID
			acc.Username = item.User.Username
			acc.ParentID = agency.CabinetID

			apiRows = append(apiRows, acc)
		}

		insertRows, deleteRows := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
		t.log.Infof("agencies cabs: inserted-updated-rows:%d; deleted-rows: %d; total:%d", len(insertRows), len(deleteRows), len(apiRows))
		cltk.Toolkit().DeleteRowsByChunk(t.log, app.Db(), deleteRows, 1000)
		cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
	}

}
func (t *cabinetsTask) addExternalAccounts() {
	var clients []m.ExtClient
	err := app.DbService().Model(&clients).Select()
	if err != nil {
		t.log.Panic(err)
	}

	dbAccs, dbRows := func() (cabs []*m.Cabinet, rows []cltk.IRow) {
		err := app.Db().Model(&cabs).Where("is_external").Select()
		if err != nil {
			t.log.Panic(err)
		}
		for _, cab := range cabs {
			rows = append(rows, cab)
		}
		return
	}()

	var apiRows []cltk.IRow
	for _, client := range clients {

		cab := &m.Cabinet{CabinetID: client.ClientCabinetId}

		for _, dba := range dbAccs {
			if cab.CabinetID == dba.CabinetID {
				cab = dba
			}
		}
		cab.Username = client.ClientName
		cab.AccessToken = client.DailyAccessToken
		cab.RefreshToken = client.RefreshToken
		cab.IsExternal = true
		apiRows = append(apiRows, cab)
	}

	insertRows, deleteRows := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
	t.log.Infof("external cabs: inserted-updated-rows:%d; deleted-rows: %d; total:%d", len(insertRows), len(deleteRows), len(apiRows))
	cltk.Toolkit().DeleteRowsByChunk(t.log, app.Db(), deleteRows, 1000)
	cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
}
