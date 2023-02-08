package cloudwatch

// EventDetail used to check the previous and current state of the CloudWatch Alarm.
type EventDetail struct {
	AlarmName     string           `json:"alarmName"`
	PreviousState EventDetailState `json:"previousState"`
	State         EventDetailState `json:"state"`
}

// EventDetailStateValue used to determine the CloudWatch Alarm state.
type EventDetailStateValue string

const (
	// EventDetailStateValueAlarm used to determine if the CloudWatch Alarm is currently triggered.
	EventDetailStateValueAlarm EventDetailStateValue = "ALARM"
)

// EventDetailState used to check the previous and current state of the CloudWatch Alarm.
type EventDetailState struct {
	Value  EventDetailStateValue `json:"value"`
	Reason string                `json:"reason"`
}
