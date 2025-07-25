// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package v1alpha1

import (
	v1 "github.com/crossplane/crossplane/apis/pkg/meta/v1"
	v11 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GeneratedFromHubConverter struct{}

func (c *GeneratedFromHubConverter) Configuration(source *v1.Configuration) *Configuration {
	var pV1alpha1Configuration *Configuration
	if source != nil {
		var v1alpha1Configuration Configuration
		v1alpha1Configuration.TypeMeta = c.v1TypeMetaToV1TypeMeta((*source).TypeMeta)
		v1alpha1Configuration.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1alpha1Configuration.Spec = c.v1ConfigurationSpecToV1alpha1ConfigurationSpec((*source).Spec)
		pV1alpha1Configuration = &v1alpha1Configuration
	}
	return pV1alpha1Configuration
}
func (c *GeneratedFromHubConverter) Provider(source *v1.Provider) *Provider {
	var pV1alpha1Provider *Provider
	if source != nil {
		var v1alpha1Provider Provider
		v1alpha1Provider.TypeMeta = c.v1TypeMetaToV1TypeMeta((*source).TypeMeta)
		v1alpha1Provider.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1alpha1Provider.Spec = c.v1ProviderSpecToV1alpha1ProviderSpec((*source).Spec)
		pV1alpha1Provider = &v1alpha1Provider
	}
	return pV1alpha1Provider
}
func (c *GeneratedFromHubConverter) pV1CrossplaneConstraintsToPV1alpha1CrossplaneConstraints(source *v1.CrossplaneConstraints) *CrossplaneConstraints {
	var pV1alpha1CrossplaneConstraints *CrossplaneConstraints
	if source != nil {
		var v1alpha1CrossplaneConstraints CrossplaneConstraints
		v1alpha1CrossplaneConstraints.Version = (*source).Version
		pV1alpha1CrossplaneConstraints = &v1alpha1CrossplaneConstraints
	}
	return pV1alpha1CrossplaneConstraints
}
func (c *GeneratedFromHubConverter) v1ConfigurationSpecToV1alpha1ConfigurationSpec(source v1.ConfigurationSpec) ConfigurationSpec {
	var v1alpha1ConfigurationSpec ConfigurationSpec
	v1alpha1ConfigurationSpec.MetaSpec = c.v1MetaSpecToV1alpha1MetaSpec(source.MetaSpec)
	return v1alpha1ConfigurationSpec
}
func (c *GeneratedFromHubConverter) v1DependencyToV1alpha1Dependency(source v1.Dependency) Dependency {
	var v1alpha1Dependency Dependency
	if source.APIVersion != nil {
		xstring := *source.APIVersion
		v1alpha1Dependency.APIVersion = &xstring
	}
	if source.Kind != nil {
		xstring2 := *source.Kind
		v1alpha1Dependency.Kind = &xstring2
	}
	if source.Package != nil {
		xstring3 := *source.Package
		v1alpha1Dependency.Package = &xstring3
	}
	if source.Provider != nil {
		xstring4 := *source.Provider
		v1alpha1Dependency.Provider = &xstring4
	}
	if source.Configuration != nil {
		xstring5 := *source.Configuration
		v1alpha1Dependency.Configuration = &xstring5
	}
	if source.Function != nil {
		xstring6 := *source.Function
		v1alpha1Dependency.Function = &xstring6
	}
	v1alpha1Dependency.Version = source.Version
	return v1alpha1Dependency
}
func (c *GeneratedFromHubConverter) v1MetaSpecToV1alpha1MetaSpec(source v1.MetaSpec) MetaSpec {
	var v1alpha1MetaSpec MetaSpec
	v1alpha1MetaSpec.Crossplane = c.pV1CrossplaneConstraintsToPV1alpha1CrossplaneConstraints(source.Crossplane)
	if source.DependsOn != nil {
		v1alpha1MetaSpec.DependsOn = make([]Dependency, len(source.DependsOn))
		for i := 0; i < len(source.DependsOn); i++ {
			v1alpha1MetaSpec.DependsOn[i] = c.v1DependencyToV1alpha1Dependency(source.DependsOn[i])
		}
	}
	return v1alpha1MetaSpec
}
func (c *GeneratedFromHubConverter) v1ProviderSpecToV1alpha1ProviderSpec(source v1.ProviderSpec) ProviderSpec {
	var v1alpha1ProviderSpec ProviderSpec
	v1alpha1ProviderSpec.MetaSpec = c.v1MetaSpecToV1alpha1MetaSpec(source.MetaSpec)
	return v1alpha1ProviderSpec
}
func (c *GeneratedFromHubConverter) v1TypeMetaToV1TypeMeta(source v11.TypeMeta) v11.TypeMeta {
	var v1TypeMeta v11.TypeMeta
	v1TypeMeta.Kind = source.Kind
	v1TypeMeta.APIVersion = source.APIVersion
	return v1TypeMeta
}

type GeneratedToHubConverter struct{}

func (c *GeneratedToHubConverter) Configuration(source *Configuration) *v1.Configuration {
	var pV1Configuration *v1.Configuration
	if source != nil {
		var v1Configuration v1.Configuration
		v1Configuration.TypeMeta = c.v1TypeMetaToV1TypeMeta2((*source).TypeMeta)
		v1Configuration.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1Configuration.Spec = c.v1alpha1ConfigurationSpecToV1ConfigurationSpec((*source).Spec)
		pV1Configuration = &v1Configuration
	}
	return pV1Configuration
}
func (c *GeneratedToHubConverter) Provider(source *Provider) *v1.Provider {
	var pV1Provider *v1.Provider
	if source != nil {
		var v1Provider v1.Provider
		v1Provider.TypeMeta = c.v1TypeMetaToV1TypeMeta2((*source).TypeMeta)
		v1Provider.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1Provider.Spec = c.v1alpha1ProviderSpecToV1ProviderSpec((*source).Spec)
		pV1Provider = &v1Provider
	}
	return pV1Provider
}
func (c *GeneratedToHubConverter) pV1alpha1CrossplaneConstraintsToPV1CrossplaneConstraints(source *CrossplaneConstraints) *v1.CrossplaneConstraints {
	var pV1CrossplaneConstraints *v1.CrossplaneConstraints
	if source != nil {
		var v1CrossplaneConstraints v1.CrossplaneConstraints
		v1CrossplaneConstraints.Version = (*source).Version
		pV1CrossplaneConstraints = &v1CrossplaneConstraints
	}
	return pV1CrossplaneConstraints
}
func (c *GeneratedToHubConverter) v1TypeMetaToV1TypeMeta2(source v11.TypeMeta) v11.TypeMeta {
	var v1TypeMeta v11.TypeMeta
	v1TypeMeta.Kind = source.Kind
	v1TypeMeta.APIVersion = source.APIVersion
	return v1TypeMeta
}
func (c *GeneratedToHubConverter) v1alpha1ConfigurationSpecToV1ConfigurationSpec(source ConfigurationSpec) v1.ConfigurationSpec {
	var v1ConfigurationSpec v1.ConfigurationSpec
	v1ConfigurationSpec.MetaSpec = c.v1alpha1MetaSpecToV1MetaSpec(source.MetaSpec)
	return v1ConfigurationSpec
}
func (c *GeneratedToHubConverter) v1alpha1DependencyToV1Dependency(source Dependency) v1.Dependency {
	var v1Dependency v1.Dependency
	if source.APIVersion != nil {
		xstring := *source.APIVersion
		v1Dependency.APIVersion = &xstring
	}
	if source.Kind != nil {
		xstring2 := *source.Kind
		v1Dependency.Kind = &xstring2
	}
	if source.Package != nil {
		xstring3 := *source.Package
		v1Dependency.Package = &xstring3
	}
	if source.Provider != nil {
		xstring4 := *source.Provider
		v1Dependency.Provider = &xstring4
	}
	if source.Configuration != nil {
		xstring5 := *source.Configuration
		v1Dependency.Configuration = &xstring5
	}
	if source.Function != nil {
		xstring6 := *source.Function
		v1Dependency.Function = &xstring6
	}
	v1Dependency.Version = source.Version
	return v1Dependency
}
func (c *GeneratedToHubConverter) v1alpha1MetaSpecToV1MetaSpec(source MetaSpec) v1.MetaSpec {
	var v1MetaSpec v1.MetaSpec
	v1MetaSpec.Crossplane = c.pV1alpha1CrossplaneConstraintsToPV1CrossplaneConstraints(source.Crossplane)
	if source.DependsOn != nil {
		v1MetaSpec.DependsOn = make([]v1.Dependency, len(source.DependsOn))
		for i := 0; i < len(source.DependsOn); i++ {
			v1MetaSpec.DependsOn[i] = c.v1alpha1DependencyToV1Dependency(source.DependsOn[i])
		}
	}
	return v1MetaSpec
}
func (c *GeneratedToHubConverter) v1alpha1ProviderSpecToV1ProviderSpec(source ProviderSpec) v1.ProviderSpec {
	var v1ProviderSpec v1.ProviderSpec
	v1ProviderSpec.MetaSpec = c.v1alpha1MetaSpecToV1MetaSpec(source.MetaSpec)
	return v1ProviderSpec
}
