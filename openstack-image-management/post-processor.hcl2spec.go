// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package openstackimagemanagement

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName             *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType           *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug                 *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce                 *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError               *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars              map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars         []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Username                    *string           `mapstructure:"username" required:"true" cty:"username" hcl:"username"`
	UserID                      *string           `mapstructure:"user_id" cty:"user_id" hcl:"user_id"`
	Password                    *string           `mapstructure:"password" required:"true" cty:"password" hcl:"password"`
	IdentityEndpoint            *string           `mapstructure:"identity_endpoint" required:"true" cty:"identity_endpoint" hcl:"identity_endpoint"`
	TenantID                    *string           `mapstructure:"tenant_id" required:"false" cty:"tenant_id" hcl:"tenant_id"`
	TenantName                  *string           `mapstructure:"tenant_name" cty:"tenant_name" hcl:"tenant_name"`
	DomainID                    *string           `mapstructure:"domain_id" cty:"domain_id" hcl:"domain_id"`
	DomainName                  *string           `mapstructure:"domain_name" required:"false" cty:"domain_name" hcl:"domain_name"`
	Insecure                    *bool             `mapstructure:"insecure" required:"false" cty:"insecure" hcl:"insecure"`
	Region                      *string           `mapstructure:"region" required:"false" cty:"region" hcl:"region"`
	EndpointType                *string           `mapstructure:"endpoint_type" required:"false" cty:"endpoint_type" hcl:"endpoint_type"`
	CACertFile                  *string           `mapstructure:"cacert" required:"false" cty:"cacert" hcl:"cacert"`
	ClientCertFile              *string           `mapstructure:"cert" required:"false" cty:"cert" hcl:"cert"`
	ClientKeyFile               *string           `mapstructure:"key" required:"false" cty:"key" hcl:"key"`
	Token                       *string           `mapstructure:"token" required:"false" cty:"token" hcl:"token"`
	ApplicationCredentialName   *string           `mapstructure:"application_credential_name" required:"false" cty:"application_credential_name" hcl:"application_credential_name"`
	ApplicationCredentialID     *string           `mapstructure:"application_credential_id" required:"false" cty:"application_credential_id" hcl:"application_credential_id"`
	ApplicationCredentialSecret *string           `mapstructure:"application_credential_secret" required:"false" cty:"application_credential_secret" hcl:"application_credential_secret"`
	Cloud                       *string           `mapstructure:"cloud" required:"false" cty:"cloud" hcl:"cloud"`
	Identifier                  *string           `mapstructure:"identifier" cty:"identifier" hcl:"identifier"`
	KeepReleases                *int              `mapstructure:"keep_releases" cty:"keep_releases" hcl:"keep_releases"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":             &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":           &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                  &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                  &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":               &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":         &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":    &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"username":                      &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"user_id":                       &hcldec.AttrSpec{Name: "user_id", Type: cty.String, Required: false},
		"password":                      &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"identity_endpoint":             &hcldec.AttrSpec{Name: "identity_endpoint", Type: cty.String, Required: false},
		"tenant_id":                     &hcldec.AttrSpec{Name: "tenant_id", Type: cty.String, Required: false},
		"tenant_name":                   &hcldec.AttrSpec{Name: "tenant_name", Type: cty.String, Required: false},
		"domain_id":                     &hcldec.AttrSpec{Name: "domain_id", Type: cty.String, Required: false},
		"domain_name":                   &hcldec.AttrSpec{Name: "domain_name", Type: cty.String, Required: false},
		"insecure":                      &hcldec.AttrSpec{Name: "insecure", Type: cty.Bool, Required: false},
		"region":                        &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"endpoint_type":                 &hcldec.AttrSpec{Name: "endpoint_type", Type: cty.String, Required: false},
		"cacert":                        &hcldec.AttrSpec{Name: "cacert", Type: cty.String, Required: false},
		"cert":                          &hcldec.AttrSpec{Name: "cert", Type: cty.String, Required: false},
		"key":                           &hcldec.AttrSpec{Name: "key", Type: cty.String, Required: false},
		"token":                         &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"application_credential_name":   &hcldec.AttrSpec{Name: "application_credential_name", Type: cty.String, Required: false},
		"application_credential_id":     &hcldec.AttrSpec{Name: "application_credential_id", Type: cty.String, Required: false},
		"application_credential_secret": &hcldec.AttrSpec{Name: "application_credential_secret", Type: cty.String, Required: false},
		"cloud":                         &hcldec.AttrSpec{Name: "cloud", Type: cty.String, Required: false},
		"identifier":                    &hcldec.AttrSpec{Name: "identifier", Type: cty.String, Required: false},
		"keep_releases":                 &hcldec.AttrSpec{Name: "keep_releases", Type: cty.Number, Required: false},
	}
	return s
}