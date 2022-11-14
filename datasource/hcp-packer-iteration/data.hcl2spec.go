// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package hcp_packer_iteration

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Bucket              *string           `mapstructure:"bucket_name" required:"true" cty:"bucket_name" hcl:"bucket_name"`
	Channel             *string           `mapstructure:"channel" required:"true" cty:"channel" hcl:"channel"`
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
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"bucket_name":                &hcldec.AttrSpec{Name: "bucket_name", Type: cty.String, Required: false},
		"channel":                    &hcldec.AttrSpec{Name: "channel", Type: cty.String, Required: false},
	}
	return s
}

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	AuthorID           *string `mapstructure:"author_id" cty:"author_id" hcl:"author_id"`
	BucketName         *string `mapstructure:"bucket_name" cty:"bucket_name" hcl:"bucket_name"`
	Complete           *bool   `mapstructure:"complete" cty:"complete" hcl:"complete"`
	CreatedAt          *string `mapstructure:"created_at" cty:"created_at" hcl:"created_at"`
	Fingerprint        *string `mapstructure:"fingerprint" cty:"fingerprint" hcl:"fingerprint"`
	ID                 *string `mapstructure:"id" cty:"id" hcl:"id"`
	IncrementalVersion *int32  `mapstructure:"incremental_version" cty:"incremental_version" hcl:"incremental_version"`
	UpdatedAt          *string `mapstructure:"updated_at" cty:"updated_at" hcl:"updated_at"`
	ChannelID          *string `mapstructure:"channel_id" cty:"channel_id" hcl:"channel_id"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"author_id":           &hcldec.AttrSpec{Name: "author_id", Type: cty.String, Required: false},
		"bucket_name":         &hcldec.AttrSpec{Name: "bucket_name", Type: cty.String, Required: false},
		"complete":            &hcldec.AttrSpec{Name: "complete", Type: cty.Bool, Required: false},
		"created_at":          &hcldec.AttrSpec{Name: "created_at", Type: cty.String, Required: false},
		"fingerprint":         &hcldec.AttrSpec{Name: "fingerprint", Type: cty.String, Required: false},
		"id":                  &hcldec.AttrSpec{Name: "id", Type: cty.String, Required: false},
		"incremental_version": &hcldec.AttrSpec{Name: "incremental_version", Type: cty.Number, Required: false},
		"updated_at":          &hcldec.AttrSpec{Name: "updated_at", Type: cty.String, Required: false},
		"channel_id":          &hcldec.AttrSpec{Name: "channel_id", Type: cty.String, Required: false},
	}
	return s
}
