package nine

import (
	"github.com/avecost/sg8bv/db"
	"time"
)

type Nine struct {
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
	Banker9     bool
	Player9     bool
}

func FindRowsByLoginGameTime(db *db.DB, login string, gameTime string) ([]Nine, error) {
	rows, err := db.Query("SELECT id, login, game_name, bet_banker, bet_player, "+
		"bet_tie, total_payout, game_number, game_time, dealer_cards, player_cards, banker_n9, player_n9 "+
		"FROM nines "+
		"WHERE total_payout != 0.00 AND game_time LIKE ? AND login = ?", gameTime, login)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nines []Nine
	for rows.Next() {
		n := Nine{}
		err = rows.Scan(&n.Id, &n.Login, &n.GameName, &n.BetBanker, &n.BetPlayer, &n.BetTie, &n.TotalPayout, &n.GameNumber,
			&n.GameTime, &n.DealerCards, &n.PlayerCards, &n.Banker9, &n.Player9)
		if err != nil {
			return nil, err
		}

		nines = append(nines, n)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return nines, nil
}
