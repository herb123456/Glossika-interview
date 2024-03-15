package services

import "log/slog"

func SendEmail(to string, subject string, body string) {
	slog.Info("Send email to " + to + " with subject: " + subject + " and body: " + body)
}
