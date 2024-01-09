package cloudwatch

// Event used to parse the CloudWatch Alarm event.
type Event struct {
	AlarmData AlarmData `json:"AlarmData"`
}

// AlarmData used to check the previous and current state of the CloudWatch Alarm.
type AlarmData struct {
	// AlarmName used to determine the CloudWatch Alarm name.
	AlarmName string `json:"alarmName"`
	// Current status of this alarm.
	State AlarmDataState `json:"state"`
	// Configuration of this alarm.
	Configuration AlarmDataConfiguration `json:"configuration"`
}

// AlarmDataState used to check the previous and current state of the CloudWatch Alarm.
type AlarmDataState struct {
	// Reason for the CloudWatch Alarm state.
	Reason string `json:"reason"`
}

// AlarmDataConfiguration used to review the configuration of the CloudWatch Alarm.
type AlarmDataConfiguration struct {
	// Description configuration on the CloudWatch Alarm.
	Description string `json:"description"`
}
