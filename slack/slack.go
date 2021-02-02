package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nlopes/slack"
)

func sendSlackMsg(token, channel, result string, start time.Time) error {
	api := slack.New(token)
	// attachment := slack.Attachment{
	// 	Text: createSlackMsg(start, result),
	// }

	//channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText("gke-test", false), slack.MsgOptionAttachments(attachment))
	channelID, timestamp, err := api.PostMessage(channel, slack.MsgOptionText(createSlackMsg(start, result), false), slack.MsgOptionUsername("gke-test-Bot"), slack.MsgOptionIconEmoji(":sunny:"))
	if err != nil {
		return fmt.Errorf("failed to send message: %s", err)
	}
	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}

func createSlackMsg(start time.Time, res string) string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	finish := time.Now()
	processingTime := time.Since(start).Truncate(time.Second)

	msg := "*gke-test が正常に終了しました。*\n"
	msg += fmt.Sprintf("起動時刻: %v\n", start.In(jst).Format("2006-01-02 15:04:05"))
	msg += fmt.Sprintf("終了時刻: %v\n", finish.In(jst).Format("2006-01-02 15:04:05"))
	msg += fmt.Sprintf("所要時間: %v\n\n", processingTime)
	msg += fmt.Sprintf("%s\n", res)
	return msg
}
