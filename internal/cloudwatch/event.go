package cloudwatch

// HasBecomeAlarm will unmarshal the event to determine if it became alarm.
func HasBecomeAlarm(detail EventDetail) bool {
	// Check if the alarm has already been triggered.
	if detail.PreviousState.Value == EventDetailStateValueAlarm {
		return false
	}

	// Check if the alarm is NOW in the triggered state.
	if detail.State.Value != EventDetailStateValueAlarm {
		return false
	}

	return true
}
