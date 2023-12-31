// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobWithAttemptsRead struct {
	Job      *JobRead      `json:"job,omitempty"`
	Attempts []AttemptRead `json:"attempts,omitempty"`
}

func (o *JobWithAttemptsRead) GetJob() *JobRead {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *JobWithAttemptsRead) GetAttempts() []AttemptRead {
	if o == nil {
		return nil
	}
	return o.Attempts
}
