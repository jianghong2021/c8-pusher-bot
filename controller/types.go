package controller

// 用户提现成功事件 Data
type CashOut struct {
	// 金额
	Amount float64 `json:"amount"`
	// 编号
	ID string `json:"id"`
	// 时间
	Time string `json:"time"`
	// 用户编号
	Uid string `json:"uid"`
	// 用户名
	Username string `json:"username"`
}

// Data
type Entertained struct {
	// 投注明细
	BetList []EntertainedBetItem `json:"betList,omitempty"`
	// 投注总金额
	BetTotalAmount float64 `json:"betTotalAmount"`
	// 投注总人数
	BetTotalPeople int64 `json:"betTotalPeople"`
	// 场次编号
	PeriodID int64 `json:"periodId"`
	// 场次名称
	PeriodName string `json:"periodName"`
	// 场次(值为0表示初盘，其他为盘中)
	PeriodNumber *int64 `json:"periodNumber,omitempty"`
	// 分局(0为正常场次，其他为分局)
	PeriodSubNumber *int64 `json:"periodSubNumber,omitempty"`
	// 赛事编号
	TournamentID int64 `json:"tournamentId"`
	// 赛事名称
	TournamentName string `json:"tournamentName"`
}

// 台球封盘事件 BetItem
type EntertainedBetItem struct {
	// 投注金额
	BetAmount float64 `json:"betAmount"`
	// 投注项
	BetSeat string `json:"betSeat"`
	// 投注项2
	BetSeat2 string `json:"betSeat2"`
	// 投注时间
	CreateTime string `json:"createTime"`
	// 编号
	ID int64 `json:"id"`
	// 玩家ID
	PlayerID int64 `json:"playerId"`
	// 玩家用户名
	PlayerName string `json:"playerName,omitempty"`
}

// 每日总流水事件 Data
type DailyFundsData struct {
	// 总提现
	CashOut *float64 `json:"cashOut,omitempty"`
	// 总盈亏
	Profit *float64 `json:"profit,omitempty"`
	// 总充值
	Recharge *float64 `json:"recharge,omitempty"`
	// 时间
	Time string `json:"time,omitempty"`
}

// 彩票28封盘
type Lottery28EntertainedData struct {
	// 投注明细
	BetList []Lottery28EntertainedDataBetItem `json:"betList,omitempty"`
	// 投注总金额
	BetTotalAmount float64 `json:"betTotalAmount"`
	// 投注总人数
	BetTotalPeople int64 `json:"betTotalPeople"`
	// 场次编号
	PeriodID int64 `json:"periodId"`
	// 场次名称
	PeriodName string `json:"periodName"`
	// 场次(值为0表示初盘，其他为盘中)
	PeriodNumber *int64 `json:"periodNumber,omitempty"`
	// 分局(0为正常场次，其他为分局)
	PeriodSubNumber *int64 `json:"periodSubNumber,omitempty"`
	// 赛事编号
	TournamentID int64 `json:"tournamentId"`
	// 赛事名称
	TournamentName string `json:"tournamentName"`
}

// 台球封盘事件 BetItem
type Lottery28EntertainedDataBetItem struct {
	// 投注金额
	BetAmount float64 `json:"betAmount"`
	// 投注项
	BetSeat string `json:"betSeat"`
	// 投注项2
	BetSeat2 string `json:"betSeat2"`
	// 投注时间
	CreateTime string `json:"createTime"`
	// 编号
	ID int64 `json:"id"`
	// 玩家ID
	PlayerID int64 `json:"playerId"`
	// 玩家用户名
	PlayerName *string `json:"playerName,omitempty"`
}

// 用户充值成功事件 Data
type PayRechargeData struct {
	// 金额
	Amount float64 `json:"amount"`
	// 编号
	ID string `json:"id"`
	// 时间
	Time string `json:"time"`
	// 用户编号
	Uid string `json:"uid"`
	// 用户名
	Username string `json:"username"`
}
