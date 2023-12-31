// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobInfoRead struct {
	Job      JobRead           `json:"job"`
	Attempts []AttemptInfoRead `json:"attempts"`
}

func (o *JobInfoRead) GetJob() JobRead {
	if o == nil {
		return JobRead{}
	}
	return o.Job
}

func (o *JobInfoRead) GetAttempts() []AttemptInfoRead {
	if o == nil {
		return []AttemptInfoRead{}
	}
	return o.Attempts
}
