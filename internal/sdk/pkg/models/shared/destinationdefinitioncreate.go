// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationDefinitionCreate struct {
	Name             string  `json:"name"`
	DockerRepository string  `json:"dockerRepository"`
	DockerImageTag   string  `json:"dockerImageTag"`
	DocumentationURL string  `json:"documentationUrl"`
	Icon             *string `json:"icon,omitempty"`
	// actor definition specific resource requirements. if default is set, these are the requirements that should be set for ALL jobs run for this actor definition. it is overriden by the job type specific configurations. if not set, the platform will use defaults. these values will be overriden by configuration at the connection level.
	ResourceRequirements *ActorDefinitionResourceRequirements `json:"resourceRequirements,omitempty"`
}

func (o *DestinationDefinitionCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationDefinitionCreate) GetDockerRepository() string {
	if o == nil {
		return ""
	}
	return o.DockerRepository
}

func (o *DestinationDefinitionCreate) GetDockerImageTag() string {
	if o == nil {
		return ""
	}
	return o.DockerImageTag
}

func (o *DestinationDefinitionCreate) GetDocumentationURL() string {
	if o == nil {
		return ""
	}
	return o.DocumentationURL
}

func (o *DestinationDefinitionCreate) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *DestinationDefinitionCreate) GetResourceRequirements() *ActorDefinitionResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}
