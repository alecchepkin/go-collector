package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
	cltk "collector-toolkit"
	"collector/app"
	"collector/tasks"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	webServer *http.Server
	crontab   *cron.Cron
)

func main() {
	stopch := make(chan os.Signal, 1)
	signal.Notify(stopch, syscall.SIGINT, syscall.SIGTERM)

	tasks.MasterTask().ScheduledRun(cltk.MonthAgo, cltk.Today)
	tasks.MasterTask().ScheduledRun(cltk.Yesterday, cltk.Today)
	scheduler()
	webApi()

	<-stopch
	shutdown()
}

func webApi() {
	router := mux.NewRouter()
	router.HandleFunc("/stat/{date-from}/{date-to}/", statHandler)
	webServer = &http.Server{Addr: app.Config().ListenAddr, Handler: router}

	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Log().Panic(err)
		}
	}()

	app.Log().Infof("Run web webApi server: %v", app.Config().ListenAddr)
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var dateFrom, dateTo time.Time
	ok := cltk.Toolkit().DatesCheckOrSetFailResp(app.Log(), w, vars["date-from"], vars["date-to"], &dateFrom, &dateTo)
	if !ok {
		return
	}

	go tasks.MasterTask().Run(dateFrom, dateTo)
	cltk.Toolkit().SetSuccessResp(app.Log(), w, r)
}
func scheduler() {

	var errs []error

	crontab = cron.New()
	errs = append(errs, crontab.AddFunc("0 */1 * * * *", func() {
		tasks.MasterTask().ScheduledRun(cltk.Yesterday, cltk.Today)
	}))
	errs = append(errs, crontab.AddFunc("0 0 * * * *", func() {
		tasks.MasterTask().ScheduledRun(cltk.MonthAgo, cltk.Today)
	}))
	crontab.Start()

	for _, err := range errs {
		if err != nil {
			app.Log().Error(err)
		}
	}

	app.Log().Info("process scheduler")
}

func shutdown() {
	app.Log().Info("Service shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = webServer.Shutdown(ctx)
	app.Log().Info("Web server down")

	crontab.Stop()
	app.Log().Info("Scheduler down")

	_ = app.Db().Close()
	_ = app.DbService().Close()
	app.Log().Info("DB connection was closed")
}
