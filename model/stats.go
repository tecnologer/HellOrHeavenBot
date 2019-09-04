package model

//StatsType identifier for type of stats
type StatsType byte

//Stats is the struct for stats data, used to register the /hell and /heaven commands
type Stats struct {
	ID       int64
	Hell     uint   `json:"hell"`
	Heaven   uint   `json:"heaven"`
	UserID   int    `json:"user_id"`
	UserName string `json:"user"`
}

const (
	StatsHell StatsType = iota
	StatsHeaven
)
