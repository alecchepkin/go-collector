package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	"collector/components"
	"collector/modeles"
	m "collector/modeles"
	"github.com/sirupsen/logrus"
)

type bannersTask struct {
	cab *m.Cabinet
	log *logrus.Entry
	api *components.ApiClient
}

func BannersTask(cab *m.Cabinet, log *logrus.Entry) *bannersTask {
	t := &bannersTask{}
	t.log = log.WithFields(logrus.Fields{"task": "banner-task", "cab": cab.CabinetID})
	t.cab = cab
	t.api = app.ApiClient(t.log)
	return t
}

func (t *bannersTask) Run() {
	defer cltk.Toolkit().Recover(t.log)
	resp, err := t.api.Banners(t.cab)
	if err != nil {
		t.log.Panic(err)
	}

	var apiRows []cltk.IRow
	for _, item := range resp.Items {

		ban := modeles.Banner{}
		ban.BannerID = item.ID
		ban.CabinetID = t.cab.CabinetID
		ban.CampaignID = item.CampaignID
		ban.ModerationStatus = item.ModerationStatus
		apiRows = append(apiRows, ban)

	}

	var regs []*modeles.Banner
	err = app.Db().Model(&regs).Where("cabinet_id=?", t.cab.CabinetID).Select()
	if err != nil {
		t.log.Panic(err)
	}

	var dbRows []cltk.IRow
	for _, acc := range regs {
		dbRows = append(dbRows, acc)
	}
	regs = nil

	insertRows, deleteRows := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
	t.log.Infof("inserted-updated-rows:%d; deleted-rows: %d", len(insertRows), len(deleteRows))
	cltk.Toolkit().DeleteRowsByChunk(t.log, app.Db(), deleteRows, 1000)
	cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
}
