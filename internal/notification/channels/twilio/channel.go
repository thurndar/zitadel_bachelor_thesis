package twilio

import (
	"github.com/caos/logging"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/notification/channels"
	"github.com/caos/zitadel/internal/notification/messages"
	twilio "github.com/kevinburke/twilio-go"
)

func InitTwilioChannel(config TwilioConfig) channels.NotificationChannel {
	client := twilio.NewClient(config.SID, config.Token, nil)

	logging.Log("NOTIF-KaxDZ").Debug("successfully initialized twilio sms channel")

	return channels.HandleMessageFunc(func(message channels.Message) error {
		twilioMsg, ok := message.(*messages.SMS)
		if !ok {
			return caos_errs.ThrowInternal(nil, "TWILI-s0pLc", "message is not SMS")
		}
		m, err := client.Messages.SendMessage(twilioMsg.SenderPhoneNumber, twilioMsg.RecipientPhoneNumber, twilioMsg.GetContent(), nil)
		if err != nil {
			return caos_errs.ThrowInternal(err, "TWILI-osk3S", "could not send message")
		}
		logging.LogWithFields("SMS_-f335c523", "message_sid", m.Sid, "status", m.Status).Debug("sms sent")
		return nil
	})
}
