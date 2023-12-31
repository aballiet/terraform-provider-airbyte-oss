// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type StreamStatusRead struct {
	AttemptNumber int    `json:"attemptNumber"`
	ConnectionID  string `json:"connectionId"`
	ID            string `json:"id"`
	JobID         int64  `json:"jobId"`
	// Values:
	//   * `FAILED` - A failure has occurred
	//   * `CANCELED` - The run has been canceled
	//
	IncompleteRunCause *StreamStatusIncompleteRunCause `json:"incompleteRunCause,omitempty"`
	JobType            StreamStatusJobType             `json:"jobType"`
	// Values:
	//   * `PENDING` - The stream operation has been selected to run
	//   * `RUNNING` - The stream operation is running
	//   * `COMPLETE` - The stream operation ran successfully to completion
	//   * `INCOMPLETE` - The stream operation has terminated in an incomplete state.
	//   See StreamStatusIncompleteRunCause for more details.
	//
	RunState        StreamStatusRunState `json:"runState"`
	StreamName      string               `json:"streamName"`
	StreamNamespace string               `json:"streamNamespace"`
	TransitionedAt  int64                `json:"transitionedAt"`
	WorkspaceID     string               `json:"workspaceId"`
}

func (o *StreamStatusRead) GetAttemptNumber() int {
	if o == nil {
		return 0
	}
	return o.AttemptNumber
}

func (o *StreamStatusRead) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *StreamStatusRead) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *StreamStatusRead) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

func (o *StreamStatusRead) GetIncompleteRunCause() *StreamStatusIncompleteRunCause {
	if o == nil {
		return nil
	}
	return o.IncompleteRunCause
}

func (o *StreamStatusRead) GetJobType() StreamStatusJobType {
	if o == nil {
		return StreamStatusJobType("")
	}
	return o.JobType
}

func (o *StreamStatusRead) GetRunState() StreamStatusRunState {
	if o == nil {
		return StreamStatusRunState("")
	}
	return o.RunState
}

func (o *StreamStatusRead) GetStreamName() string {
	if o == nil {
		return ""
	}
	return o.StreamName
}

func (o *StreamStatusRead) GetStreamNamespace() string {
	if o == nil {
		return ""
	}
	return o.StreamNamespace
}

func (o *StreamStatusRead) GetTransitionedAt() int64 {
	if o == nil {
		return 0
	}
	return o.TransitionedAt
}

func (o *StreamStatusRead) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
