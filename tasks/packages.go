package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	"collector/components"
	"collector/modeles"
	m "collector/modeles"
	rsp "collector/responses"
	"github.com/sirupsen/logrus"
)

type packagesTask struct {
	cab *m.Cabinet
	log *logrus.Entry
	api *components.ApiClient
}

func PackagesTask(cab *m.Cabinet, log *logrus.Entry) *packagesTask {
	t := packagesTask{}
	t.cab = cab
	t.log = log.WithFields(logrus.Fields{"task": "packages", "cab": cab.CabinetID})
	t.api = app.ApiClient(t.log)
	return &t
}

func (t *packagesTask) Run() {
	defer cltk.Toolkit().Recover(t.log)
	resp, err := t.api.Packages(t.cab)
	if err != nil {
		t.log.Panic(err)
	}

	apiRows := t.apiRows(resp)

	dbRows := t.dbRows()

	insertRows, deleteRows := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
	t.log.Infof("inserted-updated-rows:%d; deleted-rows: %d", len(insertRows), len(deleteRows))
	cltk.Toolkit().DeleteRowsByChunk(t.log, app.Db(), deleteRows, 1000)
	cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
}

func (t *packagesTask) apiRows(resp rsp.PackagesResp) []cltk.IRow {

	var rows []cltk.IRow

	for _, item := range resp.Items {

		pack := &modeles.Package{}
		pack.PackageID = item.ID
		pack.CabinetID = t.cab.CabinetID
		pack.Name = item.Name
		pack.Created = item.Created
		pack.Updated = item.Updated
		pack.BannerFormatID = item.BannerFormatID
		pack.Description = item.Description
		pack.Flags = item.Flags
		pack.MaxPricePerUnit = item.MaxPricePerUnit
		pack.MaxUniqShowsLimit = item.MaxUniqShowsLimit
		pack.Objective = item.Objective
		pack.PadsTreeID = item.PadsTreeID
		pack.PaidEventType = item.PaidEventType
		pack.Price = item.Price
		pack.PricedEventType = item.PricedEventType
		pack.RelatedPackageIds = item.RelatedPackageIds
		pack.Status = item.Status
		pack.UrlType = item.UrlType
		pack.UrlTypes = func() string {
			ut, err := components.ToJson(item.UrlTypes)
			if err != nil {
				t.log.Panic(err)
			}
			return ut
		}()

		rows = append(rows, pack)
	}

	return rows

}

func (t *packagesTask) dbRows() []cltk.IRow {
	var regs []*modeles.Package
	err := app.Db().Model(&regs).Where("cabinet_id=?", t.cab.CabinetID).Select()
	if err != nil {
		t.log.Panic(err)
	}

	var rows []cltk.IRow
	for _, cab := range regs {
		rows = append(rows, cab)
	}
	regs = nil
	return rows
}
