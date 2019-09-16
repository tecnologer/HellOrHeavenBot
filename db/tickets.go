package db

import (
	"github.com/recoilme/pudge"
	log "github.com/sirupsen/logrus"
	m "github.com/tecnologer/HellOrHeavenBot/model"
)

const (
	tableNameStats        = "Stats"
	queryCreateTableStats = `CREATE TABLE IF NOT EXISTS [%s] (
		[Id] integer not null primary key AUTOINCREMENT,
		[HellCount] integer not null,
		[HeavenCount] integer not null,
		[UserName] text not null,
		[UserId] integer
	)`
	queryGetStatsByUsername    = "SELECT [HellCount], [HeavenCount], [UserName] FROM [%s] WHERE [UserName]='%s'"
	queryUpdateStatsByUsername = "UPDATE [%s] SET [HellCount]=%d, [HeavenCount]=%d WHERE [UserName]='%s';"
	queryInsertStatsByUsername = "INSERT INTO [%s] (HellCount, HeavenCount, UserName) VALUES (%d, %d, '%s');"
)

var tableStatsIsCreated = false

func init() {
	createTableStats()
}

//InsertStat inserts or update registers for stats
func InsertStat(doomed string, t m.StatsType) error {
	createTableStats()

	getStatsByUser := queryf(queryGetStatsByUsername, tableNameStats, doomed)

	rows, err := execQuery(getStatsByUser)

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
	var tmpQuery query

	if isInsert {
		hell, heaven := 0, 0
		if t == m.StatsHeaven {
			heaven = 1
		} else {
			hell = 1
		}

		tmpQuery = queryf(queryInsertStatsByUsername, tableNameStats, hell, heaven, doomed)
	} else {
		if t == m.StatsHell {
			doomedStats.Hell++
		}

		if t == m.StatsHeaven {
			doomedStats.Heaven++
		}

		tmpQuery = queryf(queryUpdateStatsByUsername, tableNameStats, doomedStats.Hell, doomedStats.Heaven, doomedStats.UserName)
	}

	err = execQueryNoResult(tmpQuery)
	if err != nil {
		return err
	}
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

func createTableStats() {
	log.Printf("creating table %s\n", tableNameStats)
	if tableStatsIsCreated {
		return
	}

	err := execQueryNoResult(queryf(queryCreateTableStats, tableNameStats))
	tableStatsIsCreated = err == nil
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("table %s is created\n", tableNameStats)
	}

}
