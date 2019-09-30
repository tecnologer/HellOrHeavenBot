package db

import (
	log "github.com/sirupsen/logrus"
	hpr "github.com/tecnologer/HellOrHeavenBot/db/dbhelp"
	"github.com/tecnologer/HellOrHeavenBot/model"
	m "github.com/tecnologer/HellOrHeavenBot/model"
)

var ticketsTable = &hpr.SQLTable{
	Name: "Stats",
	Columns: []*hpr.SQLColumn{
		hpr.NewPKCol("Id"),
		hpr.NewIntCol("HellCount"),
		hpr.NewIntCol("HeavenCount"),
		hpr.NewTextCol("UserName"),
		hpr.NewIntNilCol("UserId"),
	},
}

//InsertStat inserts or update registers for stats
func InsertStat(doomed string, t m.StatsType) error {
	err := ticketsTable.Create()

	if err != nil {
		return err
	}

	conditions := []*hpr.ConditionGroup{
		&hpr.ConditionGroup{
			ConLeft: &hpr.Condition{
				Column: ticketsTable.GetColByName("UserName"),
				RelOp:  hpr.Eq,
				Value:  doomed,
			},
		},
	}

	rows, err := ticketsTable.ExecSelectCols([]string{"HellCount", "HeavenCount", "UserName"}, conditions)

	if err != nil {
		return err
	}
	defer rows.Close()
	var doomedStats *m.Stats

	for rows.Next() {
		doomedStats = new(m.Stats)
		err = rows.Scan(&doomedStats.Hell, &doomedStats.Heaven, &doomedStats.UserName)
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	var isInsert = doomedStats == nil

	if isInsert {
		hell, heaven := 0, 0
		if t == m.StatsHeaven {
			heaven = 1
		} else {
			hell = 1
		}

		err = ticketsTable.Insert(hell, heaven, doomed, nil)
	} else {
		if t == m.StatsHell {
			doomedStats.Hell++
		}

		if t == m.StatsHeaven {
			doomedStats.Heaven++
		}

		valuesToUpdate := map[string]interface{}{
			"HellCount":   doomedStats.Hell,
			"HeavenCount": doomedStats.Heaven,
		}

		err = ticketsTable.Update(valuesToUpdate, conditions)
	}

	if err != nil {
		return err
	}
	return nil
}

//GetStats returns the statistic for the user who request it
func GetStats(doomed string) *m.Stats {
	err := ticketsTable.Create()

	if err != nil {
		log.WithField("GetStatsByName", doomed).WithError(err)
		return nil
	}

	conditions := []*hpr.ConditionGroup{
		&hpr.ConditionGroup{
			ConLeft: &hpr.Condition{
				Column: ticketsTable.GetColByName("UserName"),
				RelOp:  hpr.Eq,
				Value:  doomed,
			},
		},
	}

	var doomedStats *model.Stats
	rows, err := ticketsTable.ExecSelectCols([]string{"HellCount", "HeavenCount", "UserName"}, conditions)
	if err != nil {
		log.WithField("GetStatsByName", doomed).WithError(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		doomedStats = new(model.Stats)
		err = rows.Scan(&doomedStats.Hell, &doomedStats.Heaven, &doomedStats.UserName)
		if err != nil {
			log.WithField("GetStatsByName", doomed).WithError(err)
			return nil
		}
	}
	err = rows.Err()

	if err != nil {
		log.WithField("GetStatsByName", doomed).WithError(err)
		return nil
	}
	return doomedStats
}
