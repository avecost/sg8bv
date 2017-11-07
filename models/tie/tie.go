package tie

import (
	"github.com/avecost/sg8bv/db"
	"time"
)

type Tie struct {
	Id          int
	Login       string
	GameName    string
	BetBanker   float32
	BetPlayer   float32
	BetTie      float32
	TotalPayout float32
	GameNumber  string
	GameTime    time.Time
	DealerCards string
	PlayerCards string
}

func FindRowsByLoginGameTime(db *db.DB, login string, gameTime string) ([]Tie, error) {
	rows, err := db.Query("SELECT id, login, game_name, bet_banker, bet_player, "+
		"bet_tie, total_payout, game_number, game_time, dealer_cards, player_cards "+
		"FROM ties "+
		"WHERE bet_tie != 0.00 AND game_time LIKE ? AND login = ?", gameTime, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ts []Tie
	for rows.Next() {
		t := Tie{}
		err = rows.Scan(
			&t.Id, &t.Login, &t.GameName, &t.BetBanker, &t.BetPlayer, &t.BetTie, &t.TotalPayout, &t.GameNumber,
			&t.GameTime, &t.DealerCards, &t.PlayerCards)
		if err != nil {
			return nil, err
		}

		ts = append(ts, t)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return ts, nil
}
