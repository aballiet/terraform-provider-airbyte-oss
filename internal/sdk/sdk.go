// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Your Airbyte API server
	"http://localhost:8000/api",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// SDK - Airbyte Configuration API: Airbyte Configuration API
// [https://airbyte.io](https://airbyte.io).
//
// The Configuration API is an internal Airbyte API that is designed for communications between different Airbyte components.
// * It's main purpose is to enable the Airbyte Engineering team to configure the internal state of [Airbyte Cloud](https://airbyte.com/airbyte-cloud)
// * It is also sometimes used by OSS users to configure their own Self-Hosted Airbyte deployment (internal state, etc)
//
// WARNING
// * Airbyte does NOT have active commitments to support this API long-term.
// * OSS users can utilize the Configuration API, but at their own risk.
// * This API is utilized internally by the Airbyte Engineering team and may be modified in the future if the need arises.
// * Modifications by the Airbyte Engineering team could create breaking changes and OSS users would need to update their code to catch up to any backwards incompatible changes in the API.
//
// This API is a collection of HTTP RPC-style methods. While it is not a REST API, those familiar with REST should find the conventions of this API recognizable.
//
// Here are some conventions that this API follows:
// * All endpoints are http POST methods.
// * All endpoints accept data via `application/json` request bodies. The API does not accept any data via query params.
// * The naming convention for endpoints is: localhost:8000/{VERSION}/{METHOD_FAMILY}/{METHOD_NAME} e.g. `localhost:8000/v1/connections/create`.
// * For all `update` methods, the whole object must be passed in, even the fields that did not change.
//
// Authentication (OSS):
// * When authenticating to the Configuration API, you must use Basic Authentication by setting the Authentication Header to Basic and base64 encoding the username and password (which are `airbyte` and `password` by default - so base64 encoding `airbyte:password` results in `YWlyYnl0ZTpwYXNzd29yZA==`). So the full header reads `'Authorization': "Basic YWlyYnl0ZTpwYXNzd29yZA=="`
//
// https://airbyte.io - Find out more about Airbyte
type SDK struct {
	// Workspace related resources.
	Workspace     *Workspace
	Notifications *Notifications
	// SourceDefinition related resources.
	SourceDefinition *SourceDefinition
	// Interactions with actor definition version related resources.
	ActorDefinitionVersion *ActorDefinitionVersion
	// SourceDefinition specification related resources.
	SourceDefinitionSpecification *SourceDefinitionSpecification
	DeclarativeSourceDefinitions  *DeclarativeSourceDefinitions
	ConnectorBuilderProject       *ConnectorBuilderProject
	ConnectorDocumentation        *ConnectorDocumentation
	// Source related resources.
	Source   *Source
	Internal *Internal
	// DestinationDefinition related resources.
	DestinationDefinition *DestinationDefinition
	// DestinationDefinitionSpecification related resources.
	DestinationDefinitionSpecification *DestinationDefinitionSpecification
	// Destination related resources.
	Destination *Destination
	// Connection between sources and destinations.
	Connection *Connection
	// Interactions with state related resources.
	State     *State
	Operation *Operation
	Scheduler *Scheduler
	// Source OAuth related resources to delegate access from user.
	SourceOauth *SourceOauth
	// Source OAuth related resources to delegate access from user.
	DestinationOauth *DestinationOauth
	// Endpoints for the Airbyte web application. Those APIs should not be called outside the web application implementation and are not
	// guaranteeing any backwards compatibility.
	//
	WebBackend *WebBackend
	Jobs       *Jobs
	// Interactions with attempt related resources.
	Attempt *Attempt
	// Interactions with user related resources.
	User *User
	// Interactions with permission related resources.
	Permission               *Permission
	SecretsPersistenceConfig *SecretsPersistenceConfig
	// Healthchecks
	Health                *Health
	Logs                  *Logs
	Openapi               *Openapi
	StreamStatuses        *StreamStatuses
	Streams               *Streams
	InstanceConfiguration *InstanceConfiguration
	JobRetryStates        *JobRetryStates
	// Interactions with organizations.
	Organization *Organization

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SDK)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SDK) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SDK) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.0",
			SDKVersion:        "1.1.0",
			GenVersion:        "2.194.1",
			UserAgent:         "speakeasy-sdk/go 1.1.0 2.194.1 1.0.0 airbyte",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Workspace = newWorkspace(sdk.sdkConfiguration)

	sdk.Notifications = newNotifications(sdk.sdkConfiguration)

	sdk.SourceDefinition = newSourceDefinition(sdk.sdkConfiguration)

	sdk.ActorDefinitionVersion = newActorDefinitionVersion(sdk.sdkConfiguration)

	sdk.SourceDefinitionSpecification = newSourceDefinitionSpecification(sdk.sdkConfiguration)

	sdk.DeclarativeSourceDefinitions = newDeclarativeSourceDefinitions(sdk.sdkConfiguration)

	sdk.ConnectorBuilderProject = newConnectorBuilderProject(sdk.sdkConfiguration)

	sdk.ConnectorDocumentation = newConnectorDocumentation(sdk.sdkConfiguration)

	sdk.Source = newSource(sdk.sdkConfiguration)

	sdk.Internal = newInternal(sdk.sdkConfiguration)

	sdk.DestinationDefinition = newDestinationDefinition(sdk.sdkConfiguration)

	sdk.DestinationDefinitionSpecification = newDestinationDefinitionSpecification(sdk.sdkConfiguration)

	sdk.Destination = newDestination(sdk.sdkConfiguration)

	sdk.Connection = newConnection(sdk.sdkConfiguration)

	sdk.State = newState(sdk.sdkConfiguration)

	sdk.Operation = newOperation(sdk.sdkConfiguration)

	sdk.Scheduler = newScheduler(sdk.sdkConfiguration)

	sdk.SourceOauth = newSourceOauth(sdk.sdkConfiguration)

	sdk.DestinationOauth = newDestinationOauth(sdk.sdkConfiguration)

	sdk.WebBackend = newWebBackend(sdk.sdkConfiguration)

	sdk.Jobs = newJobs(sdk.sdkConfiguration)

	sdk.Attempt = newAttempt(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	sdk.Permission = newPermission(sdk.sdkConfiguration)

	sdk.SecretsPersistenceConfig = newSecretsPersistenceConfig(sdk.sdkConfiguration)

	sdk.Health = newHealth(sdk.sdkConfiguration)

	sdk.Logs = newLogs(sdk.sdkConfiguration)

	sdk.Openapi = newOpenapi(sdk.sdkConfiguration)

	sdk.StreamStatuses = newStreamStatuses(sdk.sdkConfiguration)

	sdk.Streams = newStreams(sdk.sdkConfiguration)

	sdk.InstanceConfiguration = newInstanceConfiguration(sdk.sdkConfiguration)

	sdk.JobRetryStates = newJobRetryStates(sdk.sdkConfiguration)

	sdk.Organization = newOrganization(sdk.sdkConfiguration)

	return sdk
}
