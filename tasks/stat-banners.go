package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	"collector/components"
	m "collector/modeles"
	"collector/responses"
	"github.com/sirupsen/logrus"
	"time"
)

type statBannersTask struct {
	cab      *m.Cabinet
	log      *logrus.Entry
	dateFrom time.Time
	dateTo   time.Time
	api      *components.ApiClient
}

func StatBannersTask(cab *m.Cabinet, log *logrus.Entry, dateFrom time.Time, dateTo time.Time) *statBannersTask {
	t := &statBannersTask{}
	t.log = log.WithFields(logrus.Fields{"task": "stat-banner", "cab": cab.CabinetID})
	t.cab = cab
	t.dateFrom = dateFrom
	t.dateTo = dateTo
	t.api = app.ApiClient(t.log)

	return t
}

func (t *statBannersTask) Run() {
	defer cltk.Toolkit().Recover(t.log)
	var resps []responses.StatBannersDayResp
	bannerIdsChunks, l := t.bannerIdsChunks()
	t.log.Info("banners total:", l, "chunks total:", len(bannerIdsChunks))
	for i, chunk := range bannerIdsChunks {
		t.log.Info("api", "chunk:", i)
		resp, err := t.api.StatBanners(t.cab, chunk, t.dateFrom, t.dateTo)
		if err != nil {
			t.log.Panic(err)
		}
		resps = append(resps, resp)

	}
	apiRows := t.apiRows(resps)
	dbRows := t.dbRows()

	insertRows, deleteRows := cltk.Toolkit().DetectInsertAndDeleteRows(apiRows, dbRows)
	cltk.Toolkit().DeleteRowsByChunk(t.log, app.Db(), deleteRows, 1000)
	cltk.Toolkit().InsertUpdateRowsByChunk(t.log, app.Db(), insertRows, 1000)
	t.log.Infof("Inserted/updated rows: %v", len(insertRows))
	t.log.Infof("Deleted rows: %v", len(deleteRows))
}

func (t *statBannersTask) bannerIdsChunks() (chunks [][]int64, l int) {
	var bannerIDs []int64
	err := app.Db().Model(&m.Banner{}).Where("cabinet_id=?", t.cab.CabinetID).Column("banner_id").Select(&bannerIDs)
	if err != nil {
		t.log.Panic(err)
	}
	l = len(bannerIDs)
	const max = 10000
	for i := 0; i < l; i += max {
		r := i + max
		if r > l {
			r = l
		}
		chunks = append(chunks, bannerIDs[i:r])
	}
	return
}

func (t *statBannersTask) apiRows(resps []responses.StatBannersDayResp) []cltk.IRow {
	var stats []cltk.IRow
	for _, resp := range resps {
		for _, item := range resp.Items {
			for _, row := range item.Rows {
				s := &m.StatBanner{}
				s.Date = row.Date
				s.CabinetID = t.cab.CabinetID
				s.Shows = row.Base.Shows
				s.Clicks = row.Base.Clicks
				s.Goals = row.Base.Goals
				s.Spent = row.Base.Spent
				s.Reach = row.Uniques.Reach
				s.Total = row.Uniques.Total
				s.Increment = row.Uniques.Increment

				s.BannerID = item.ID
				stats = append(stats, s)
			}

		}

	}
	return stats
}

func (t *statBannersTask) dbRows() []cltk.IRow {
	var stats []*m.StatBanner
	err := app.Db().Model(&stats).Where("cabinet_id=?", t.cab.CabinetID).Select()
	if err != nil {
		t.log.Panic(err)
	}
	var iRows []cltk.IRow
	for _, s := range stats {
		iRows = append(iRows, s)
	}
	return iRows
}
