package ical

import (
  "strconv"
  "strings"
	"time"
)

type calSerializer struct {
	calendar *Calendar
	buffer   *strBuffer
}

func (this *calSerializer) serialize() string {
	this.serializeCalendar()
	return strings.TrimSpace(this.buffer.String())
}

func (this *calSerializer) serializeCalendar() {
	this.begin()
	this.version()
	this.items()
	this.end()
}

func (this *calSerializer) begin() {
	this.buffer.Write("BEGIN:VCALENDAR\n")
}

func (this *calSerializer) end() {
	this.buffer.Write("END:VCALENDAR\n")
}

func (this *calSerializer) version() {
	this.buffer.Write("VERSION:2.0\n")
}

func (this *calSerializer) items() {
	for _, item := range this.calendar.Items {
		item.serializeWithBuffer(this.buffer)
	}
}

type calEventSerializer struct {
	event  *CalendarEvent
	buffer *strBuffer
}

const (
	eventSerializerTimeFormat = "20060102T150405Z"
)

func (this *calEventSerializer) serialize() string {
	this.serializeEvent()
	return strings.TrimSpace(this.buffer.String())
}

func (this *calEventSerializer) serializeEvent() {
	this.begin()
	this.uid()
	this.created()
	this.lastModified()
	this.dtstart()
	this.dtend()
	this.organizer()
	this.attendee()
	this.summary()
	this.description()
	this.location()
	this.url()
	this.end()
}

func (this *calEventSerializer) begin() {
	this.buffer.Write("BEGIN:VEVENT\n")
}

func (this *calEventSerializer) end() {
	this.buffer.Write("END:VEVENT\n")
}

func (this *calEventSerializer) uid() {
	this.serializeStringProp("UID", this.event.Id)
}

func (this *calEventSerializer) summary() {
	this.serializeStringProp("SUMMARY", this.event.Summary)
}

func (this *calEventSerializer) description() {
	this.serializeStringProp("DESCRIPTION", this.event.Description)
}

func (this *calEventSerializer) location() {
	this.serializeStringProp("LOCATION", this.event.Location)
}

func (this *calEventSerializer) url() {
	this.serializeStringProp("URL", this.event.URL)
}

func (this *calEventSerializer) dtstart() {
	this.serializeTimeProp("DTSTART", this.event.StartAtUTC())
}

func (this *calEventSerializer) dtend() {
	this.serializeTimeProp("DTEND", this.event.EndAtUTC())
}

func (this *calEventSerializer) organizer() {
  if this.event.Organizer != nil {
    this.serializeStringProp("ORGANIZER", "CN="+this.event.Organizer.CN+":mailto:"+this.event.Organizer.Address)
  }
}

func (this *calEventSerializer) attendee() {
  if this.event.Attendees != nil {
    for _, a := range this.event.Attendees {
      this.serializeStringProp("ATTENDEE", "RSVP="+strconv.FormatBool(a.RSVP)+";ROLE="+a.ParticipationRole+";CUTYPE="+a.CalenderUserType+":mailto:"+a.Address)
    }
  }
}

func (this *calEventSerializer) created() {
	this.serializeTimeProp("CREATED", this.event.CreatedAtUTC)
}

func (this *calEventSerializer) lastModified() {
	this.serializeTimeProp("LAST-MODIFIED", this.event.ModifiedAtUTC)
}

func (this *calEventSerializer) serializeStringProp(name, value string) {
	if value != "" {
		escapedValue := escapeTextType(value)
		this.buffer.Write("%s:%s\n", name, escapedValue)
	}
}

func (this *calEventSerializer) serializeTimeProp(name string, value *time.Time) {
	if value != nil {
		this.buffer.Write("%s:%s\n", name, value.Format(eventSerializerTimeFormat))
	}
}
