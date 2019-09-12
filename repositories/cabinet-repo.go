package repositories

import (
	"collector/modeles"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

type CabinetRepo struct {
	db  *pg.DB
	log *logrus.Entry
}

//
//func (t *CabinetRepo) SetLog(log *logrus.Entry) *CabinetRepo {
//	t.log = log
//	return t
//}

func NewCabinetRepo(db *pg.DB, log *logrus.Entry) *CabinetRepo {
	return &CabinetRepo{db: db, log: log.WithField("repo", "cabinet")}
}

func (t *CabinetRepo) FindAllAccounts() []*modeles.Cabinet {
	var cabs []*modeles.Cabinet
	err := t.db.Model(&cabs).Select()
	if err != nil {
		t.log.Panic(err)
	}

	t.withAgencies(cabs)
	return cabs
}

func (t *CabinetRepo) withAgencies(cabs []*modeles.Cabinet) {
	agencies := t.FindAllAgencies()
	agencyMap := make(map[int]*modeles.Cabinet, 0)
	for _, agency := range agencies {
		agencyMap[agency.CabinetID] = agency
	}
	for _, cab := range cabs {
		if cab.ParentID == 0 {
			continue
		}
		if agency, ok := agencyMap[cab.ParentID]; ok {
			cab.Parent = agency
			continue
		}
		t.log.Panic("Parent with AccountId:", cab.ParentID, "not found")
	}
}

func (t *CabinetRepo) ClientCabs() []*modeles.Cabinet {
	var cabs []*modeles.Cabinet
	err := t.db.Model(&cabs).Where("t.is_agency=false").Select()
	if err != nil {
		t.log.Panic(err)
	}
	t.withAgencies(cabs)

	return cabs
}

func (t *CabinetRepo) ValidClients() []*modeles.Cabinet {
	var cabs []*modeles.Cabinet
	err := t.db.Model(&cabs).Where("t.is_agency=false and cabinet_id>0").Select()
	if err != nil {
		t.log.Panic(err)
	}
	t.withAgencies(cabs)

	return cabs
}

func (t *CabinetRepo) FindAllAgencies() []*modeles.Cabinet {
	var cabs []*modeles.Cabinet
	err := t.db.Model(&cabs).Where("is_agency=true").Select()
	if err != nil {
		t.log.Panic(err)
	}
	return cabs
}

func (t *CabinetRepo) UpdateAccounts(accounts []modeles.Cabinet) {
	err := t.db.Update(&accounts)
	if err != nil {
		t.log.Panic(err)
	}
}
