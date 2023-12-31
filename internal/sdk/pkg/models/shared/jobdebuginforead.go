// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobDebugInfoRead struct {
	Job           JobDebugRead       `json:"job"`
	Attempts      []AttemptInfoRead  `json:"attempts"`
	WorkflowState *WorkflowStateRead `json:"workflowState,omitempty"`
}

func (o *JobDebugInfoRead) GetJob() JobDebugRead {
	if o == nil {
		return JobDebugRead{}
	}
	return o.Job
}

func (o *JobDebugInfoRead) GetAttempts() []AttemptInfoRead {
	if o == nil {
		return []AttemptInfoRead{}
	}
	return o.Attempts
}

func (o *JobDebugInfoRead) GetWorkflowState() *WorkflowStateRead {
	if o == nil {
		return nil
	}
	return o.WorkflowState
}
