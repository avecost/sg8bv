package raffle

import (
	"github.com/avecost/sg8bv/db"
	"time"
)

type Raffle struct {
	Id           int
	TerminalAcct string
	JackpotAt    time.Time
	JackpotAmt   float32
}

func GetAllPendingBaccaratResultsByGameId(db *db.DB, providerId, gameId int, dateTo string) ([]Raffle, error) {
	rows, err := db.Query("SELECT id, terminal_acct, jackpot_at, jackpot_amt "+
		"FROM raffles "+
		"WHERE valid = ? "+
		"AND provider_id = ? "+
		"AND game_id = ? "+
		"AND jackpot_at < ? "+
		"ORDER BY jackpot_at", 1, providerId, gameId, dateTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rs []Raffle
	for rows.Next() {
		r := Raffle{}
		err = rows.Scan(&r.Id, &r.TerminalAcct, &r.JackpotAt, &r.JackpotAmt)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return rs, nil
}

// status:
// 0 invalid / 1 pending / 2 valid
func UpdateRaffleStatus(db *db.DB, raffleId, status int) error {
	_, err := db.Exec("UPDATE raffles SET valid = ?, validated = ? WHERE id = ?", status, 1, raffleId)
	if err != nil {
		return err
	}

	updateEntriesStatus(db, raffleId, status)

	return nil
}

// status:
// 0 invalid / 1 pending / 2 valid
func updateEntriesStatus(db *db.DB, raffleId, status int) error {
	_, err := db.Exec("UPDATE entries SET valid = ? WHERE player_id = ?", status, raffleId)
	if err != nil {
		return err
	}

	return nil
}
