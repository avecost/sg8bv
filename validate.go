package sg8bv

import (
	"fmt"

	"github.com/avecost/sg8bv/db"
	"github.com/avecost/sg8bv/models/nine"
	"github.com/avecost/sg8bv/models/raffle"
	"github.com/avecost/sg8bv/models/tie"
)

type Valid struct {
	AppDb *db.DB
}

func Init(conn string) (*Valid, error) {
	appDb, err := db.Open(conn)
	if err != nil {
		return nil, err
	}

	return &Valid{AppDb: appDb}, nil
}

func (v *Valid) Run(dateTo string) {
	rtRows, err := raffle.GetAllPendingBaccaratResultsByGameId(v.AppDb, 1, 366, dateTo)
	if err != nil {
		println(err.Error())
	}
	println("Processing SG8 Baccarat Raffle Entries (Tie) : ", len(rtRows))
	for _, rtRow := range rtRows {
		df := rtRow.JackpotAt.Format("2006-01-02") + "%"
		fmt.Println(rtRow.Id, rtRow.TerminalAcct, df, rtRow.JackpotAmt)

		tRows, err := tie.FindRowsByLoginGameTime(v.AppDb, rtRow.TerminalAcct, df)
		if err != nil {
			println(err.Error())
		}
		if len(tRows) > 0 { // valid
			raffle.UpdateRaffleStatus(v.AppDb, rtRow.Id, 2)
		} else { // invalid
			raffle.UpdateRaffleStatus(v.AppDb, rtRow.Id, 0)
		}
	}

	r9Rows, err := raffle.GetAllPendingBaccaratResultsByGameId(v.AppDb, 1, 367, dateTo)
	if err != nil {
		println(err.Error())
	}
	println("Processing SG8 Baccarat Raffle Entries (Nine) : ", len(r9Rows))
	for _, r9Row := range r9Rows {
		df := r9Row.JackpotAt.Format("2006-01-02") + "%"
		fmt.Println(r9Row.Id, r9Row.TerminalAcct, df, r9Row.JackpotAmt)

		nRows, err := nine.FindRowsByLoginGameTime(v.AppDb, r9Row.TerminalAcct, df)
		if err != nil {
			println(err.Error())
		}
		if len(nRows) > 0 { // valid
			raffle.UpdateRaffleStatus(v.AppDb, r9Row.Id, 2)
		} else { // invalid
			raffle.UpdateRaffleStatus(v.AppDb, r9Row.Id, 0)
		}
	}

}
