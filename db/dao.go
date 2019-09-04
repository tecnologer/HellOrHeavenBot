package db

import (
	"github.com/recoilme/pudge"
	log "github.com/sirupsen/logrus"
	m "github.com/tecnologer/HellOrHeavenBot/model"
)

const (
	stats = "stats"
)

//InsertStat inserts or update registers for stats
func InsertStat(doomed string, t m.StatsType) error {
	db, err := pudge.Open("./db/dbFiles/Stats", nil)

	if err != nil {
		return err
	}
	defer db.Close()
	var doomedStats m.Stats
	err = db.Get(doomed, &doomedStats)

	if err != nil {
		doomedStats = m.Stats{
			UserName: doomed,
		}

		if t == m.StatsHell {
			doomedStats.Hell = 1
		}
		if t == m.StatsHeaven {
			doomedStats.Heaven = 1
		}
	} else {
		if t == m.StatsHell {
			doomedStats.Hell++
		}

		if t == m.StatsHeaven {
			doomedStats.Heaven++
		}
	}

	err = db.Set(doomed, &doomedStats)
	if err != nil {
		return err
	}

	// id, err := db.Counter(stats, 0)

	// if err != nil {
	// 	return err
	// }

	return nil
}

func GetStats(doomed string) *m.Stats {
	db, err := pudge.Open("./db/dbFiles/Stats", nil)

	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()

	var doomedStats m.Stats
	err = db.Get(doomed, &doomedStats)

	if err != nil {
		log.Println(err)
		return nil
	}
	return &doomedStats
}
