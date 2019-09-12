package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	"collector/components"
	m "collector/modeles"
	"github.com/sirupsen/logrus"
)

type campaignsTask struct {
	cab *m.Cabinet
	log *logrus.Entry
	api *components.ApiClient
}

func CampaignsTask(cab *m.Cabinet, log *logrus.Entry) *campaignsTask {
	t := campaignsTask{}
	t.log = log.WithFields(logrus.Fields{"task": "campaigns", "cab": cab.CabinetID})
	t.cab = cab
	t.api = app.ApiClient(t.log)
	return &t
}

func (t *campaignsTask) Run() {
	defer cltk.Toolkit().Recover(t.log)
	resp, err := t.api.Campaigns(t.cab)
	if err != nil {
		t.log.Panic(err)
	}

	var apiRows []cltk.IRow
	for _, item := range resp.Items {

		camp := m.Campaign{}
		camp.CampaignID = int64(item.ID)
		camp.CabinetID = t.cab.CabinetID
		camp.Name = item.Name
		camp.BudgetLimit = 0.0
		camp.BudgetLimitDay = 0.0
		camp.Created = item.Created
		camp.Updated = item.Updated
		camp.Status = item.Status
		apiRows = append(apiRows, camp)
	}

	var dbRows []cltk.IRow
	var loadedRows []*m.Campaign
	err = app.Db().Model(&loadedRows).Where("cabinet_id=?", t.cab.CabinetID).Select()
	if err != nil {
		t.log.Panic(err)
	}
	for _, row := range loadedRows {
		dbRows = append(dbRows, row)
	}
	loadedRows = nil

	insertRows, deleteRows := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
	t.log.Infof("inserted-updated-rows:%d; deleted-rows: %d", len(insertRows), len(deleteRows))
	cltk.Toolkit().DeleteRowsByChunk(t.log, app.Db(), deleteRows, 1000)
	cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
}
