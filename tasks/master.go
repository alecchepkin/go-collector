package tasks

import (
	cltk "collector-toolkit"
	"collector/app"
	"collector/components"
	m "collector/modeles"
	"collector/repositories"
	"github.com/sirupsen/logrus"
	"time"
)

const masterTaskName = "master"

type masterTask struct {
	log        *logrus.Entry
	cabRepo    *repositories.CabinetRepo
	proxy      *components.Proxy
	finishChan chan int
}

type Runnable interface {
	Run()
}

func MasterTask() *masterTask {
	t := &masterTask{}
	t.log = app.Log().WithField("task", "master")
	t.cabRepo = app.CabinetRepo(t.log)
	t.proxy = app.Proxy(t.log)
	t.finishChan = make(chan int)
	return t
}

type subMasterTask struct {
	log      *logrus.Entry
	master   *masterTask
	cab      *m.Cabinet
	dateFrom time.Time
	dateTo   time.Time
	api      *components.ApiClient
}

func newSubMasterTask(
	master *masterTask,
	cab *m.Cabinet,
	dateFrom time.Time,
	dateTo time.Time,
) *subMasterTask {
	s := &subMasterTask{
		log:      master.log.WithField("cab", cab.CabinetID),
		master:   master,
		cab:      cab,
		dateFrom: dateFrom,
		dateTo:   dateTo,
	}
	s.api = app.ApiClient(s.log)
	return s
}

func (t *masterTask) ScheduledRun(dateFromStr string, dateToStr string) {

	dateFrom, dateTo := cltk.Toolkit().DatePeriodStrToTime(t.log, dateFromStr, dateToStr)
	t.Run(dateFrom, dateTo)
}

func (t *masterTask) Run(dateFrom time.Time, dateTo time.Time) {
	startTime := time.Now()
	t.log = cltk.Toolkit().ExpandLog(t.log, false, masterTaskName, "", dateFrom, dateTo, "")
	//defer cltk.Toolkit().Recover(t.log)
	defer func() { t.log.Infof("time execution:%v", time.Now().Sub(startTime)) }()
	skipKey := cltk.Toolkit().ParamsToKey(masterTaskName, dateFrom.String(), dateTo.String())
	queueKey := cltk.Toolkit().ParamsToKey(masterTaskName)

	if !cltk.Toolkit().TryToStartTask(t.log, skipKey, queueKey) {
		return
	}
	defer cltk.Toolkit().FinishTask(skipKey, queueKey)

	CabinetsTask(t.log).Run()

	cabs := t.cabRepo.ClientCabs()

	t.proxy.UpdateTokens(cabs)

	cabs = t.cabRepo.ValidClients()

	RegionsTask(cabs, t.log).Run()

	//os.Exit(23)
	for _, cab := range cabs {
		t.log.Infof("process cab: %d", cab.CabinetID)
		sbat := newSubMasterTask(t, cab, dateFrom, dateTo)
		go sbat.process()
	}

	fin := 0
	l := len(cabs)
	for cabId := range t.finishChan {
		fin++
		t.log.Info(fin, "/", l, "; finish for cab:", cabId)
		if l == fin {
			t.log.Info("finish processing all cabs:", l)
			return
		}
	}

}

func (s *subMasterTask) process() {
	defer cltk.Toolkit().Recover(s.log)
	defer func() { s.master.finishChan <- s.cab.CabinetID }()

	tasks := []Runnable{
		UserInfoTask(s.cab, s.log),
		CampaignsTask(s.cab, s.log),
		PackagesTask(s.cab, s.log),
		BannersTask(s.cab, s.log),
		StatBannersTask(s.cab, s.log, s.dateFrom, s.dateTo),
	}

	c := make(chan int)
	for i, task := range tasks {
		go func(r Runnable, i int) {
			defer func() { c <- i }()
			r.Run()
		}(task, i)
	}

	fin := 0
	l := len(tasks)
	for i := range c {
		fin++
		s.log.Info(fin, "/", l, " key:", i)
		if fin == l {
			close(c)
		}
	}

}
