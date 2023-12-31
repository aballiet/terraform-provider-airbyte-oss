// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Configuration struct {
}

// SecretPersistenceConfig - Object representing a secret persistence configuration
type SecretPersistenceConfig struct {
	SecretPersistenceType SecretPersistenceType `json:"secretPersistenceType"`
	Configuration         Configuration         `json:"configuration"`
	ScopeType             ScopeType             `json:"scopeType"`
	ScopeID               string                `json:"scopeId"`
}

func (o *SecretPersistenceConfig) GetSecretPersistenceType() SecretPersistenceType {
	if o == nil {
		return SecretPersistenceType("")
	}
	return o.SecretPersistenceType
}

func (o *SecretPersistenceConfig) GetConfiguration() Configuration {
	if o == nil {
		return Configuration{}
	}
	return o.Configuration
}

func (o *SecretPersistenceConfig) GetScopeType() ScopeType {
	if o == nil {
		return ScopeType("")
	}
	return o.ScopeType
}

func (o *SecretPersistenceConfig) GetScopeID() string {
	if o == nil {
		return ""
	}
	return o.ScopeID
}
