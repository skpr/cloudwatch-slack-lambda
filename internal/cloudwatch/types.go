package cloudwatch

// Event used to parse the CloudWatch Alarm event.
type Event struct {
	AlarmData AlarmData `json:"AlarmData"`
}

// AlarmData used to check the previous and current state of the CloudWatch Alarm.
type AlarmData struct {
	// AlarmName used to determine the CloudWatch Alarm name.
	AlarmName string `json:"alarmName"`
	// Previous state of this alarm.
	PreviousState AlarmDataState `json:"previousState"`
	// Current status of this alarm.
	State AlarmDataState `json:"state"`
	// Configuration of this alarm.
	Configuration AlarmDataConfiguration `json:"configuration"`
}

// AlarmDataStateValue used to determine the CloudWatch Alarm state.
type AlarmDataStateValue string

const (
	// AlarmDataStateValueAlarm used to determine if the CloudWatch Alarm is currently triggered.
	AlarmDataStateValueAlarm AlarmDataStateValue = "ALARM"
	// AlarmDataStateValueInsufficientData used to determine if the CloudWatch Alarm is currently in an insufficient data state.
	AlarmDataStateValueInsufficientData AlarmDataStateValue = "INSUFFICIENT_DATA"
	// AlarmDataStateValueOK used to determine if the CloudWatch Alarm is currently OK.
	AlarmDataStateValueOK AlarmDataStateValue = "OK"
)

// AlarmDataState used to check the previous and current state of the CloudWatch Alarm.
type AlarmDataState struct {
	// Value used to determine the CloudWatch Alarm state.
	Value AlarmDataStateValue `json:"value"`
	// Reason for the CloudWatch Alarm state.
	Reason string `json:"reason"`
}

// AlarmDataConfiguration used to review the configuration of the CloudWatch Alarm.
type AlarmDataConfiguration struct {
	// Description configuration on the CloudWatch Alarm.
	Description string `json:"description"`
}
