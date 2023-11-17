// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk"
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &AirbyteProvider{}

type AirbyteProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AirbyteProviderModel describes the provider data model.
type AirbyteProviderModel struct {
	ServerURL types.String `tfsdk:"server_url"`
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
}

func (p *AirbyteProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "airbyte"
	resp.Version = p.version
}

func (p *AirbyteProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Airbyte Configuration API: Airbyte Configuration API` + "\n" +
			`[https://airbyte.io](https://airbyte.io).` + "\n" +
			`` + "\n" +
			`The Configuration API is an internal Airbyte API that is designed for communications between different Airbyte components.` + "\n" +
			`* It's main purpose is to enable the Airbyte Engineering team to configure the internal state of [Airbyte Cloud](https://airbyte.com/airbyte-cloud)` + "\n" +
			`* It is also sometimes used by OSS users to configure their own Self-Hosted Airbyte deployment (internal state, etc)` + "\n" +
			`` + "\n" +
			`WARNING` + "\n" +
			`* Airbyte does NOT have active commitments to support this API long-term.` + "\n" +
			`* OSS users can utilize the Configuration API, but at their own risk.` + "\n" +
			`* This API is utilized internally by the Airbyte Engineering team and may be modified in the future if the need arises.` + "\n" +
			`* Modifications by the Airbyte Engineering team could create breaking changes and OSS users would need to update their code to catch up to any backwards incompatible changes in the API.` + "\n" +
			`` + "\n" +
			`This API is a collection of HTTP RPC-style methods. While it is not a REST API, those familiar with REST should find the conventions of this API recognizable.` + "\n" +
			`` + "\n" +
			`Here are some conventions that this API follows:` + "\n" +
			`* All endpoints are http POST methods.` + "\n" +
			`* All endpoints accept data via ` + "`" + `application/json` + "`" + ` request bodies. The API does not accept any data via query params.` + "\n" +
			`* The naming convention for endpoints is: localhost:8000/{VERSION}/{METHOD_FAMILY}/{METHOD_NAME} e.g. ` + "`" + `localhost:8000/v1/connections/create` + "`" + `.` + "\n" +
			`* For all ` + "`" + `update` + "`" + ` methods, the whole object must be passed in, even the fields that did not change.` + "\n" +
			`` + "\n" +
			`Authentication (OSS):` + "\n" +
			`* When authenticating to the Configuration API, you must use Basic Authentication by setting the Authentication Header to Basic and base64 encoding the username and password (which are ` + "`" + `airbyte` + "`" + ` and ` + "`" + `password` + "`" + ` by default - so base64 encoding ` + "`" + `airbyte:password` + "`" + ` results in ` + "`" + `YWlyYnl0ZTpwYXNzd29yZA==` + "`" + `). So the full header reads ` + "`" + `'Authorization': "Basic YWlyYnl0ZTpwYXNzd29yZA=="` + "`" + `` + "\n" +
			``,
		Attributes: map[string]schema.Attribute{
			"server_url": schema.StringAttribute{
				MarkdownDescription: "Server URL (defaults to http://localhost:8000/api)",
				Optional:            true,
				Required:            false,
			},
			"username": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *AirbyteProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data AirbyteProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ServerURL := data.ServerURL.ValueString()

	if ServerURL == "" {
		ServerURL = "http://localhost:8000/api"
	}

	username := data.Username.ValueString()
	password := data.Password.ValueString()
	security := shared.Security{
		Username: username,
		Password: password,
	}

	opts := []sdk.SDKOption{
		sdk.WithServerURL(ServerURL),
		sdk.WithSecurity(security),
	}
	client := sdk.New(opts...)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *AirbyteProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewConnectionResource,
		NewSourceResource,
		NewSourceDefinitionResource,
		NewSourceDefinitionManifestResource,
		NewWorkspaceResource,
	}
}

func (p *AirbyteProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSourceDataSource,
		NewSourceSchemaDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AirbyteProvider{
			version: version,
		}
	}
}
