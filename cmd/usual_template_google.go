package cmd

const usualGoogleAnalytics = `package google

import (
    "{ms-name}/settings"
    "net/http"
    "net/url"
    "strings"
)

type GMessage struct {
    Version          string // Api version +
    TrackingId       string // Tracking ID +
    AnonymizeIp      string // Anonymize IP
    DataSource       string // Data Source
    ClientId         string // Client ID
    UserId           string // User ID
    SessionControl   string // Session control
    IpOverride       string // IP override
    UserAgenOverride string // UA override
    GeoOverride      string // Geographical Override
    UserLanguage     string // User Language
    HitType          string // Hit type
    DocumentPath     string // Document Path
    DocumentHostName string // Document Host Name
    EventCategory    string // Event category
    EventAction      string // Event action
    EventValue       string // Event value
}

func NewGMessage() *GMessage {
    return &GMessage{
        Version:    "1",
        HitType:    "event",
        TrackingId: settings.GoogleTrackingId,
    }
}

func (m *GMessage) Send() error {
    v := url.Values{}

    if m.Version != "" {
        v.Add("v", m.Version)
    }
    if m.TrackingId != "" {
        v.Add("tid", m.TrackingId)
    }
    if m.AnonymizeIp != "" {
        v.Add("aip", m.AnonymizeIp)
    }
    if m.DataSource != "" {
        v.Add("ds", m.DataSource)
    }
    if m.ClientId != "" {
        v.Add("cid", m.ClientId)
    }
    if m.UserId != "" {
        v.Add("uid", m.UserId)
    }
    if m.SessionControl != "" {
        v.Add("sc", m.SessionControl)
    }
    if m.IpOverride != "" {
        v.Add("uip", m.IpOverride)
    }
    if m.UserAgenOverride != "" {
        v.Add("ua", m.UserAgenOverride)
    }
    if m.GeoOverride != "" {
        v.Add("geoid", m.GeoOverride)
    }
    if m.UserLanguage != "" {
        v.Add("ul", m.UserLanguage)
    }
    if m.HitType != "" {
        v.Add("t", m.HitType)
    }
    if m.DocumentPath != "" {
        v.Add("dp", m.DocumentPath)
    }
    if m.DocumentHostName != "" {
        v.Add("dh", m.DocumentHostName)
    }
    if m.EventCategory != "" {
        v.Add("ec", m.EventCategory)
    }
    if m.EventAction != "" {
        v.Add("ea", m.EventAction)
    }
    if m.EventValue != "" {
        v.Add("ev", m.EventValue)
    }

    req, _ := http.NewRequest("POST", settings.GoogleAnalyticsUrl, strings.NewReader(v.Encode()))

    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("User-Agent", "Gosha{ms-name}Server/1.0.0")

    _, err := http.DefaultClient.Do(req)

    return err
}

`

var usualTemplateGoogleAnalytics = template{
	Path:    "./google/analytics.go",
	Content: assignCurrentDateTime(assignMsName(usualGoogleAnalytics)),
}
