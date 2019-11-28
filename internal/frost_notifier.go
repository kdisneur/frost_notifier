package internal

import (
	"context"
	"fmt"
	"frostnotifier/internal/localcache"
	"frostnotifier/internal/openweather"
	"frostnotifier/internal/twilio"
	"time"
)

type FrostNotifier struct {
	Twilio      *twilio.Twilio
	LocalCache  *localcache.LocalCache
	OpenWeather *openweather.OpenWeather
	Message     string
	Logger      Logger
}

type FrostReportStatus int

//go:generate stringer -type=FrostReportStatus -linecomment

const (
	FrostReportStatusFrostRisk FrostReportStatus = iota // Frost Risk
	FrostReportStatusNoRisk                             // No Risk

	SmallestHourBeforeNextDay int = 8
)

type FrostReport struct {
	TimeRange          TimeRange
	HumidityPercentage float64
	Status             FrostReportStatus
	TempCelsius        float64
}

func (f *FrostNotifier) Run(ctx context.Context, now time.Time, recipient string) error {
	timeRange := f.currentNight(now)

	logger := f.Logger.AddMetadata(Metadata{
		"from": timeRange.From().Format("2006-01-02"),
		"to":   timeRange.To().Format("2006-01-02"),
	})

	alreadySent, err := f.LocalCache.HasItem(ctx, timeRange.String())
	if err != nil {
		return err
	}

	if alreadySent {
		logger.Info("notification already sent")
		return nil
	}

	report, err := f.buildReport(ctx, logger, timeRange)
	if err != nil {
		return err
	}

	if report.Status == FrostReportStatusNoRisk {
		logger.Info("no notification to send")
		return nil
	}

	logger.Debug("send text message notification")
	if err := f.Twilio.SendMessage(ctx, recipient, f.Message); err != nil {
		return fmt.Errorf("can't send notification: %w", err)
	}

	if err := f.LocalCache.Save(ctx, report.TimeRange.String()); err != nil {
		logger.Warn(fmt.Sprintf("can't persist night in cache: %s", err.Error()))
	}

	logger.InfoWithMetadata(
		"notification sent",
		Metadata{
			"temperature": fmt.Sprintf("%.2f°C", report.TempCelsius),
			"humidity":    fmt.Sprintf("%.2f%%", report.HumidityPercentage),
		},
	)

	return nil
}

func (f *FrostNotifier) buildReport(ctx context.Context, logger Logger, timeRange TimeRange) (FrostReport, error) {
	probes, err := f.OpenWeather.Get(ctx, timeRange.From(), timeRange.To())
	if err != nil {
		return FrostReport{}, fmt.Errorf("can't fetch weather forecast: %w", err)
	}

	for _, probe := range probes {
		logger.Debug(fmt.Sprintf("probe: %#+v", probe))

		if f.hasFrostRisk(probe) {
			logger.Debug(fmt.Sprintf("probe with frost risk: %#+v", probe))
			return FrostReport{
				TimeRange:          timeRange,
				HumidityPercentage: probe.HumidityPercentage,
				TempCelsius:        probe.TempCelsius,
				Status:             FrostReportStatusFrostRisk,
			}, nil
		}
	}

	return FrostReport{TimeRange: timeRange, Status: FrostReportStatusNoRisk}, nil
}

func (f *FrostNotifier) currentNight(now time.Time) TimeRange {
	startingDay := now
	if startingDay.Hour() < 8 {
		startingDay = time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location())
	}

	from := time.Date(startingDay.Year(), startingDay.Month(), startingDay.Day(), 19, 0, 0, 0, startingDay.Location())
	to := time.Date(startingDay.Year(), startingDay.Month(), startingDay.Day()+1, 8, 0, 0, 0, startingDay.Location())

	return NewTimeRange(from, to)
}

func (f *FrostNotifier) hasFrostRisk(probe openweather.Probe) bool {
	return probe.TempCelsius <= 0 || (probe.TempCelsius <= 3 && probe.HumidityPercentage >= 80)
}
