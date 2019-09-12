package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	comp "collector/components"
	m "collector/modeles"
	"collector/responses"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type regionsTask struct {
	cabs []*m.Cabinet
	log  *logrus.Entry
	api  *comp.ApiClient
}

func RegionsTask(cabs []*m.Cabinet, log *logrus.Entry) *regionsTask {
	t := regionsTask{}
	t.cabs = cabs
	t.log = log.WithFields(logrus.Fields{"task": "regions"})
	t.api = app.ApiClient(t.log)
	return &t
}

func (t *regionsTask) Run() {
	//defer cltk.Toolkit().Recover(t.log)
	l := len(t.cabs)
	// we need update regions with any cab, and break immediately
	for i, cab := range t.cabs {
		resp, err := t.api.Regions(cab)
		if err != nil {
			t.log.Errorln(err, "cab:", cab.CabinetID, "; will try with next cab; hope:", i+1, "/", l)
			continue
		}

		apiRows := t.apiRows(resp)

		dbRows := t.dbRows()

		insertRows, deleteRows := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
		t.log.Infof("inserted-updated-rows:%d; deleted-rows: %d", len(insertRows), len(deleteRows))
		cltk.Toolkit().DeleteRowsByChunk(t.log, app.Db(), deleteRows, 1000)
		cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
		t.log.Info("packages successfully updated")
		break
	}
	t.log.Panic("Failed to update regions")
}

func (t *regionsTask) apiRows(resp responses.RegionsResp) []cltk.IRow {
	var regions []cltk.IRow
	for _, item := range resp.Items {
		reg := &m.Region{}
		reg.RegionID = item.ID
		reg.Name = item.Name
		reg.ParentID = item.ParentID
		reg.Flags = func() string {
			d, err := json.Marshal(item.Flags)
			if err != nil {
				t.log.Panic(err)
			}
			return string(d)
		}()

		regions = append(regions, reg)
	}
	return regions

}

func (t *regionsTask) dbRows() []cltk.IRow {
	var regs []*m.Region
	err := app.Db().Model(&regs).Select()
	if err != nil {
		t.log.Panic(err)
	}

	var rows []cltk.IRow
	for _, acc := range regs {
		rows = append(rows, acc)
	}
	regs = nil
	return rows
}
