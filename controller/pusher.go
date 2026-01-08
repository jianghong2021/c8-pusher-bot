package controller

import (
	"cg-tg-bot/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func PushayCashOut(c *gin.Context) {

	var data CashOut
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数解析错误",
		})
		log.Println("参数错误:", err.Error())
		return
	}

	msg := "```💸用户提现💸\n"
	th := []string{"编号", "金额", "时间", "用户编号", "用户名"}
	tr := []string{data.ID, fmt.Sprintf("%.2f", data.Amount), data.Time, data.Uid, data.Username}
	tb := [][]string{tr}
	msg += utils.BuildMarkdownV2List(th, tb, "")
	msg += "```"
	err = utils.SendMessageWithMarkdown(msg)

	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "推送失败",
		})
		log.Println("推送失败:", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送成功",
	})
}

func PayRecharge(c *gin.Context) {

	var data PayRechargeData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数解析错误",
		})
		log.Println("参数错误:", err.Error())
		return
	}

	msg := "```💸用户充值成功💸\n"
	th := []string{"编号", "金额", "时间", "用户编号", "用户名"}
	tr := []string{data.ID, fmt.Sprintf("%.2f", data.Amount), data.Time, data.Uid, data.Username}
	tb := [][]string{tr}
	msg += utils.BuildMarkdownV2List(th, tb, "")

	msg += "```"
	err = utils.SendMessageWithMarkdown(msg)

	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "推送失败",
		})
		log.Println("推送失败:", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送成功",
	})
}

func BilliardEntertained(c *gin.Context) {

	var data Entertained
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数解析错误",
		})
		log.Println("参数错误:", err.Error())
		return
	}

	msg := "```"
	msg += "🎱台球封盘🎱\n"
	sth := []string{"ID", "玩家ID", "玩家用户名", "投注项", "投注项2", "投注时间", "投注金额"}
	stb := [][]string{}
	for _, x := range data.BetList {
		str := []string{fmt.Sprintf("%d", x.ID), fmt.Sprintf("%d", x.PlayerID), x.PlayerName, x.BetSeat, x.BetSeat2, x.CreateTime, fmt.Sprintf("%0.2f", x.BetAmount)}
		stb = append(stb, str)
	}
	msg += utils.BuildMarkdownV2List(sth, stb, " \\- ")

	th := []string{"赛事名称", "场次名称", "投注总金额", "投注总人数"}
	tr := []string{data.TournamentName, data.PeriodName, fmt.Sprintf("%.2f", data.BetTotalAmount), fmt.Sprintf("%d", data.BetTotalPeople)}
	tb := [][]string{tr}
	msg += utils.BuildMarkdownV2List(th, tb, "")

	msg += "```"

	err = utils.SendMessageWithMarkdown(msg)

	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "推送失败",
		})
		log.Println("推送失败:", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送成功",
	})
}

func DailyFunds(c *gin.Context) {
	var data DailyFundsData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数解析错误",
		})
		log.Println("参数错误:", err.Error())
		return
	}

	msg := "```💹昨日总流水💹\n"
	th := []string{"总提现", "总盈亏", "总充值", "时间"}
	tr := []string{fmt.Sprintf("%.2f", *data.CashOut), fmt.Sprintf("%.2f", *data.Profit), fmt.Sprintf("%.2f", *data.Recharge), data.Time}
	tb := [][]string{tr}
	msg += utils.BuildMarkdownV2List(th, tb, "")
	msg += "```"
	err = utils.SendMessageWithMarkdown(msg)

	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "推送失败",
		})
		log.Println("推送失败:", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送成功",
	})
}

func Lottery28Entertained(c *gin.Context) {
	var data Lottery28EntertainedData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "参数解析错误",
		})
		log.Println("参数错误:", err.Error())
		return
	}

	msg := "```"

	msg += "🎰彩票28封盘🎰\n"
	sth := []string{"ID", "玩家ID", "玩家用户名", "投注项", "投注项2", "投注时间", "投注金额"}
	stb := [][]string{}
	for _, x := range data.BetList {
		str := []string{fmt.Sprintf("%d", x.ID), fmt.Sprintf("%d", x.PlayerID), *x.PlayerName, x.BetSeat, x.BetSeat2, x.CreateTime, fmt.Sprintf("%0.2f", x.BetAmount)}
		stb = append(stb, str)
	}
	msg += utils.BuildMarkdownV2List(sth, stb, " \\- ")

	th := []string{"赛事名称", "场次名称", "投注总金额", "投注总人数"}
	tr := []string{data.TournamentName, data.PeriodName, fmt.Sprintf("%.2f", data.BetTotalAmount), fmt.Sprintf("%d", data.BetTotalPeople)}
	tb := [][]string{tr}
	msg += utils.BuildMarkdownV2List(th, tb, "")

	msg += "```"

	err = utils.SendMessageWithMarkdown(msg)

	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "推送失败",
		})
		log.Println("推送失败:", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "推送成功",
	})
}
