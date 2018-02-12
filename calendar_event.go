package ical

import (
  "net/mail"
  "time"
)

type Address struct {
  CN string
  Address string
}

type Attendee struct {
  RSVP              bool
  ParticipationRole string
  CalenderUserType  string
  Address           string
}

type CalendarEvent struct {
	Id            string
	Summary       string
	Description   string
	Location      string
	URL           string
	CreatedAtUTC  *time.Time
	ModifiedAtUTC *time.Time
	StartAt       *time.Time
	EndAt         *time.Time
	Organizer     *Address
	Attendees     []Attendee
}

func (this *CalendarEvent) StartAtUTC() *time.Time {
	return inUTC(this.StartAt)
}

func (this *CalendarEvent) EndAtUTC() *time.Time {
	return inUTC(this.EndAt)
}

func (this *CalendarEvent) Serialize() string {
	buffer := new(strBuffer)
	return this.serializeWithBuffer(buffer)
}

func (this *CalendarEvent) ToICS() string {
	return this.Serialize()
}

func (this *CalendarEvent) serializeWithBuffer(buffer *strBuffer) string {
	serializer := calEventSerializer{
		event:  this,
		buffer: buffer,
	}
	return serializer.serialize()
}

func inUTC(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}

	tUTC := t.UTC()
	return &tUTC
}
