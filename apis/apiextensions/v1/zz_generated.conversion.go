// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package v1

import (
	v11 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type GeneratedRevisionSpecConverter struct{}

func (c *GeneratedRevisionSpecConverter) FromRevisionSpec(source CompositionRevisionSpec) CompositionSpec {
	var v1CompositionSpec CompositionSpec
	v1CompositionSpec.CompositeTypeRef = c.v1TypeReferenceToV1TypeReference(source.CompositeTypeRef)
	var pV1CompositionMode *CompositionMode
	if source.Mode != nil {
		v1CompositionMode := CompositionMode(*source.Mode)
		pV1CompositionMode = &v1CompositionMode
	}
	v1CompositionSpec.Mode = pV1CompositionMode
	var v1PatchSetList []PatchSet
	if source.PatchSets != nil {
		v1PatchSetList = make([]PatchSet, len(source.PatchSets))
		for i := 0; i < len(source.PatchSets); i++ {
			v1PatchSetList[i] = c.v1PatchSetToV1PatchSet(source.PatchSets[i])
		}
	}
	v1CompositionSpec.PatchSets = v1PatchSetList
	var v1ComposedTemplateList []ComposedTemplate
	if source.Resources != nil {
		v1ComposedTemplateList = make([]ComposedTemplate, len(source.Resources))
		for j := 0; j < len(source.Resources); j++ {
			v1ComposedTemplateList[j] = c.v1ComposedTemplateToV1ComposedTemplate(source.Resources[j])
		}
	}
	v1CompositionSpec.Resources = v1ComposedTemplateList
	var v1PipelineStepList []PipelineStep
	if source.Pipeline != nil {
		v1PipelineStepList = make([]PipelineStep, len(source.Pipeline))
		for k := 0; k < len(source.Pipeline); k++ {
			v1PipelineStepList[k] = c.v1PipelineStepToV1PipelineStep(source.Pipeline[k])
		}
	}
	v1CompositionSpec.Pipeline = v1PipelineStepList
	var pString *string
	if source.WriteConnectionSecretsToNamespace != nil {
		xstring := *source.WriteConnectionSecretsToNamespace
		pString = &xstring
	}
	v1CompositionSpec.WriteConnectionSecretsToNamespace = pString
	v1CompositionSpec.PublishConnectionDetailsWithStoreConfigRef = c.pV1StoreConfigReferenceToPV1StoreConfigReference(source.PublishConnectionDetailsWithStoreConfigRef)
	return v1CompositionSpec
}
func (c *GeneratedRevisionSpecConverter) ToRevisionSpec(source CompositionSpec) CompositionRevisionSpec {
	var v1CompositionRevisionSpec CompositionRevisionSpec
	v1CompositionRevisionSpec.CompositeTypeRef = c.v1TypeReferenceToV1TypeReference(source.CompositeTypeRef)
	var pV1CompositionMode *CompositionMode
	if source.Mode != nil {
		v1CompositionMode := CompositionMode(*source.Mode)
		pV1CompositionMode = &v1CompositionMode
	}
	v1CompositionRevisionSpec.Mode = pV1CompositionMode
	var v1PatchSetList []PatchSet
	if source.PatchSets != nil {
		v1PatchSetList = make([]PatchSet, len(source.PatchSets))
		for i := 0; i < len(source.PatchSets); i++ {
			v1PatchSetList[i] = c.v1PatchSetToV1PatchSet(source.PatchSets[i])
		}
	}
	v1CompositionRevisionSpec.PatchSets = v1PatchSetList
	var v1ComposedTemplateList []ComposedTemplate
	if source.Resources != nil {
		v1ComposedTemplateList = make([]ComposedTemplate, len(source.Resources))
		for j := 0; j < len(source.Resources); j++ {
			v1ComposedTemplateList[j] = c.v1ComposedTemplateToV1ComposedTemplate(source.Resources[j])
		}
	}
	v1CompositionRevisionSpec.Resources = v1ComposedTemplateList
	var v1PipelineStepList []PipelineStep
	if source.Pipeline != nil {
		v1PipelineStepList = make([]PipelineStep, len(source.Pipeline))
		for k := 0; k < len(source.Pipeline); k++ {
			v1PipelineStepList[k] = c.v1PipelineStepToV1PipelineStep(source.Pipeline[k])
		}
	}
	v1CompositionRevisionSpec.Pipeline = v1PipelineStepList
	var pString *string
	if source.WriteConnectionSecretsToNamespace != nil {
		xstring := *source.WriteConnectionSecretsToNamespace
		pString = &xstring
	}
	v1CompositionRevisionSpec.WriteConnectionSecretsToNamespace = pString
	v1CompositionRevisionSpec.PublishConnectionDetailsWithStoreConfigRef = c.pV1StoreConfigReferenceToPV1StoreConfigReference(source.PublishConnectionDetailsWithStoreConfigRef)
	return v1CompositionRevisionSpec
}
func (c *GeneratedRevisionSpecConverter) pRuntimeRawExtensionToPRuntimeRawExtension(source *runtime.RawExtension) *runtime.RawExtension {
	var pRuntimeRawExtension *runtime.RawExtension
	if source != nil {
		runtimeRawExtension := ConvertRawExtension((*source))
		pRuntimeRawExtension = &runtimeRawExtension
	}
	return pRuntimeRawExtension
}
func (c *GeneratedRevisionSpecConverter) pV1CombineToPV1Combine(source *Combine) *Combine {
	var pV1Combine *Combine
	if source != nil {
		var v1Combine Combine
		var v1CombineVariableList []CombineVariable
		if (*source).Variables != nil {
			v1CombineVariableList = make([]CombineVariable, len((*source).Variables))
			for i := 0; i < len((*source).Variables); i++ {
				v1CombineVariableList[i] = c.v1CombineVariableToV1CombineVariable((*source).Variables[i])
			}
		}
		v1Combine.Variables = v1CombineVariableList
		v1Combine.Strategy = CombineStrategy((*source).Strategy)
		v1Combine.String = c.pV1StringCombineToPV1StringCombine((*source).String)
		pV1Combine = &v1Combine
	}
	return pV1Combine
}
func (c *GeneratedRevisionSpecConverter) pV1ConvertTransformToPV1ConvertTransform(source *ConvertTransform) *ConvertTransform {
	var pV1ConvertTransform *ConvertTransform
	if source != nil {
		var v1ConvertTransform ConvertTransform
		v1ConvertTransform.ToType = TransformIOType((*source).ToType)
		var pV1ConvertTransformFormat *ConvertTransformFormat
		if (*source).Format != nil {
			v1ConvertTransformFormat := ConvertTransformFormat(*(*source).Format)
			pV1ConvertTransformFormat = &v1ConvertTransformFormat
		}
		v1ConvertTransform.Format = pV1ConvertTransformFormat
		pV1ConvertTransform = &v1ConvertTransform
	}
	return pV1ConvertTransform
}
func (c *GeneratedRevisionSpecConverter) pV1MapTransformToPV1MapTransform(source *MapTransform) *MapTransform {
	var pV1MapTransform *MapTransform
	if source != nil {
		var v1MapTransform MapTransform
		var mapStringV1JSON map[string]v1.JSON
		if (*source).Pairs != nil {
			mapStringV1JSON = make(map[string]v1.JSON, len((*source).Pairs))
			for key, value := range (*source).Pairs {
				mapStringV1JSON[key] = c.v1JSONToV1JSON(value)
			}
		}
		v1MapTransform.Pairs = mapStringV1JSON
		pV1MapTransform = &v1MapTransform
	}
	return pV1MapTransform
}
func (c *GeneratedRevisionSpecConverter) pV1MatchConditionReadinessCheckToPV1MatchConditionReadinessCheck(source *MatchConditionReadinessCheck) *MatchConditionReadinessCheck {
	var pV1MatchConditionReadinessCheck *MatchConditionReadinessCheck
	if source != nil {
		var v1MatchConditionReadinessCheck MatchConditionReadinessCheck
		v1MatchConditionReadinessCheck.Type = v11.ConditionType((*source).Type)
		v1MatchConditionReadinessCheck.Status = v12.ConditionStatus((*source).Status)
		pV1MatchConditionReadinessCheck = &v1MatchConditionReadinessCheck
	}
	return pV1MatchConditionReadinessCheck
}
func (c *GeneratedRevisionSpecConverter) pV1MatchTransformToPV1MatchTransform(source *MatchTransform) *MatchTransform {
	var pV1MatchTransform *MatchTransform
	if source != nil {
		var v1MatchTransform MatchTransform
		var v1MatchTransformPatternList []MatchTransformPattern
		if (*source).Patterns != nil {
			v1MatchTransformPatternList = make([]MatchTransformPattern, len((*source).Patterns))
			for i := 0; i < len((*source).Patterns); i++ {
				v1MatchTransformPatternList[i] = c.v1MatchTransformPatternToV1MatchTransformPattern((*source).Patterns[i])
			}
		}
		v1MatchTransform.Patterns = v1MatchTransformPatternList
		v1MatchTransform.FallbackValue = c.v1JSONToV1JSON((*source).FallbackValue)
		v1MatchTransform.FallbackTo = MatchFallbackTo((*source).FallbackTo)
		pV1MatchTransform = &v1MatchTransform
	}
	return pV1MatchTransform
}
func (c *GeneratedRevisionSpecConverter) pV1MathTransformToPV1MathTransform(source *MathTransform) *MathTransform {
	var pV1MathTransform *MathTransform
	if source != nil {
		var v1MathTransform MathTransform
		v1MathTransform.Type = MathTransformType((*source).Type)
		var pInt64 *int64
		if (*source).Multiply != nil {
			xint64 := *(*source).Multiply
			pInt64 = &xint64
		}
		v1MathTransform.Multiply = pInt64
		var pInt642 *int64
		if (*source).ClampMin != nil {
			xint642 := *(*source).ClampMin
			pInt642 = &xint642
		}
		v1MathTransform.ClampMin = pInt642
		var pInt643 *int64
		if (*source).ClampMax != nil {
			xint643 := *(*source).ClampMax
			pInt643 = &xint643
		}
		v1MathTransform.ClampMax = pInt643
		pV1MathTransform = &v1MathTransform
	}
	return pV1MathTransform
}
func (c *GeneratedRevisionSpecConverter) pV1MergeOptionsToPV1MergeOptions(source *v11.MergeOptions) *v11.MergeOptions {
	var pV1MergeOptions *v11.MergeOptions
	if source != nil {
		var v1MergeOptions v11.MergeOptions
		var pBool *bool
		if (*source).KeepMapValues != nil {
			xbool := *(*source).KeepMapValues
			pBool = &xbool
		}
		v1MergeOptions.KeepMapValues = pBool
		var pBool2 *bool
		if (*source).AppendSlice != nil {
			xbool2 := *(*source).AppendSlice
			pBool2 = &xbool2
		}
		v1MergeOptions.AppendSlice = pBool2
		pV1MergeOptions = &v1MergeOptions
	}
	return pV1MergeOptions
}
func (c *GeneratedRevisionSpecConverter) pV1PatchPolicyToPV1PatchPolicy(source *PatchPolicy) *PatchPolicy {
	var pV1PatchPolicy *PatchPolicy
	if source != nil {
		var v1PatchPolicy PatchPolicy
		var pV1FromFieldPathPolicy *FromFieldPathPolicy
		if (*source).FromFieldPath != nil {
			v1FromFieldPathPolicy := FromFieldPathPolicy(*(*source).FromFieldPath)
			pV1FromFieldPathPolicy = &v1FromFieldPathPolicy
		}
		v1PatchPolicy.FromFieldPath = pV1FromFieldPathPolicy
		v1PatchPolicy.MergeOptions = c.pV1MergeOptionsToPV1MergeOptions((*source).MergeOptions)
		pV1PatchPolicy = &v1PatchPolicy
	}
	return pV1PatchPolicy
}
func (c *GeneratedRevisionSpecConverter) pV1SecretReferenceToPV1SecretReference(source *v11.SecretReference) *v11.SecretReference {
	var pV1SecretReference *v11.SecretReference
	if source != nil {
		var v1SecretReference v11.SecretReference
		v1SecretReference.Name = (*source).Name
		v1SecretReference.Namespace = (*source).Namespace
		pV1SecretReference = &v1SecretReference
	}
	return pV1SecretReference
}
func (c *GeneratedRevisionSpecConverter) pV1StoreConfigReferenceToPV1StoreConfigReference(source *StoreConfigReference) *StoreConfigReference {
	var pV1StoreConfigReference *StoreConfigReference
	if source != nil {
		var v1StoreConfigReference StoreConfigReference
		v1StoreConfigReference.Name = (*source).Name
		pV1StoreConfigReference = &v1StoreConfigReference
	}
	return pV1StoreConfigReference
}
func (c *GeneratedRevisionSpecConverter) pV1StringCombineToPV1StringCombine(source *StringCombine) *StringCombine {
	var pV1StringCombine *StringCombine
	if source != nil {
		var v1StringCombine StringCombine
		v1StringCombine.Format = (*source).Format
		pV1StringCombine = &v1StringCombine
	}
	return pV1StringCombine
}
func (c *GeneratedRevisionSpecConverter) pV1StringTransformJoinToPV1StringTransformJoin(source *StringTransformJoin) *StringTransformJoin {
	var pV1StringTransformJoin *StringTransformJoin
	if source != nil {
		var v1StringTransformJoin StringTransformJoin
		v1StringTransformJoin.Separator = (*source).Separator
		pV1StringTransformJoin = &v1StringTransformJoin
	}
	return pV1StringTransformJoin
}
func (c *GeneratedRevisionSpecConverter) pV1StringTransformRegexpToPV1StringTransformRegexp(source *StringTransformRegexp) *StringTransformRegexp {
	var pV1StringTransformRegexp *StringTransformRegexp
	if source != nil {
		var v1StringTransformRegexp StringTransformRegexp
		v1StringTransformRegexp.Match = (*source).Match
		var pInt *int
		if (*source).Group != nil {
			xint := *(*source).Group
			pInt = &xint
		}
		v1StringTransformRegexp.Group = pInt
		pV1StringTransformRegexp = &v1StringTransformRegexp
	}
	return pV1StringTransformRegexp
}
func (c *GeneratedRevisionSpecConverter) pV1StringTransformToPV1StringTransform(source *StringTransform) *StringTransform {
	var pV1StringTransform *StringTransform
	if source != nil {
		var v1StringTransform StringTransform
		v1StringTransform.Type = StringTransformType((*source).Type)
		var pString *string
		if (*source).Format != nil {
			xstring := *(*source).Format
			pString = &xstring
		}
		v1StringTransform.Format = pString
		var pV1StringConversionType *StringConversionType
		if (*source).Convert != nil {
			v1StringConversionType := StringConversionType(*(*source).Convert)
			pV1StringConversionType = &v1StringConversionType
		}
		v1StringTransform.Convert = pV1StringConversionType
		var pString2 *string
		if (*source).Trim != nil {
			xstring2 := *(*source).Trim
			pString2 = &xstring2
		}
		v1StringTransform.Trim = pString2
		v1StringTransform.Regexp = c.pV1StringTransformRegexpToPV1StringTransformRegexp((*source).Regexp)
		v1StringTransform.Join = c.pV1StringTransformJoinToPV1StringTransformJoin((*source).Join)
		pV1StringTransform = &v1StringTransform
	}
	return pV1StringTransform
}
func (c *GeneratedRevisionSpecConverter) v1CombineVariableToV1CombineVariable(source CombineVariable) CombineVariable {
	var v1CombineVariable CombineVariable
	v1CombineVariable.FromFieldPath = source.FromFieldPath
	return v1CombineVariable
}
func (c *GeneratedRevisionSpecConverter) v1ComposedTemplateToV1ComposedTemplate(source ComposedTemplate) ComposedTemplate {
	var v1ComposedTemplate ComposedTemplate
	var pString *string
	if source.Name != nil {
		xstring := *source.Name
		pString = &xstring
	}
	v1ComposedTemplate.Name = pString
	v1ComposedTemplate.Base = ConvertRawExtension(source.Base)
	var v1PatchList []Patch
	if source.Patches != nil {
		v1PatchList = make([]Patch, len(source.Patches))
		for i := 0; i < len(source.Patches); i++ {
			v1PatchList[i] = c.v1PatchToV1Patch(source.Patches[i])
		}
	}
	v1ComposedTemplate.Patches = v1PatchList
	var v1ConnectionDetailList []ConnectionDetail
	if source.ConnectionDetails != nil {
		v1ConnectionDetailList = make([]ConnectionDetail, len(source.ConnectionDetails))
		for j := 0; j < len(source.ConnectionDetails); j++ {
			v1ConnectionDetailList[j] = c.v1ConnectionDetailToV1ConnectionDetail(source.ConnectionDetails[j])
		}
	}
	v1ComposedTemplate.ConnectionDetails = v1ConnectionDetailList
	var v1ReadinessCheckList []ReadinessCheck
	if source.ReadinessChecks != nil {
		v1ReadinessCheckList = make([]ReadinessCheck, len(source.ReadinessChecks))
		for k := 0; k < len(source.ReadinessChecks); k++ {
			v1ReadinessCheckList[k] = c.v1ReadinessCheckToV1ReadinessCheck(source.ReadinessChecks[k])
		}
	}
	v1ComposedTemplate.ReadinessChecks = v1ReadinessCheckList
	return v1ComposedTemplate
}
func (c *GeneratedRevisionSpecConverter) v1ConnectionDetailToV1ConnectionDetail(source ConnectionDetail) ConnectionDetail {
	var v1ConnectionDetail ConnectionDetail
	var pString *string
	if source.Name != nil {
		xstring := *source.Name
		pString = &xstring
	}
	v1ConnectionDetail.Name = pString
	var pV1ConnectionDetailType *ConnectionDetailType
	if source.Type != nil {
		v1ConnectionDetailType := ConnectionDetailType(*source.Type)
		pV1ConnectionDetailType = &v1ConnectionDetailType
	}
	v1ConnectionDetail.Type = pV1ConnectionDetailType
	var pString2 *string
	if source.FromConnectionSecretKey != nil {
		xstring2 := *source.FromConnectionSecretKey
		pString2 = &xstring2
	}
	v1ConnectionDetail.FromConnectionSecretKey = pString2
	var pString3 *string
	if source.FromFieldPath != nil {
		xstring3 := *source.FromFieldPath
		pString3 = &xstring3
	}
	v1ConnectionDetail.FromFieldPath = pString3
	var pString4 *string
	if source.Value != nil {
		xstring4 := *source.Value
		pString4 = &xstring4
	}
	v1ConnectionDetail.Value = pString4
	return v1ConnectionDetail
}
func (c *GeneratedRevisionSpecConverter) v1FunctionCredentialsToV1FunctionCredentials(source FunctionCredentials) FunctionCredentials {
	var v1FunctionCredentials FunctionCredentials
	v1FunctionCredentials.Name = source.Name
	v1FunctionCredentials.Source = FunctionCredentialsSource(source.Source)
	v1FunctionCredentials.SecretRef = c.pV1SecretReferenceToPV1SecretReference(source.SecretRef)
	return v1FunctionCredentials
}
func (c *GeneratedRevisionSpecConverter) v1FunctionReferenceToV1FunctionReference(source FunctionReference) FunctionReference {
	var v1FunctionReference FunctionReference
	v1FunctionReference.Name = source.Name
	return v1FunctionReference
}
func (c *GeneratedRevisionSpecConverter) v1JSONToV1JSON(source v1.JSON) v1.JSON {
	var v1JSON v1.JSON
	var byteList []uint8
	if source.Raw != nil {
		byteList = make([]uint8, len(source.Raw))
		for i := 0; i < len(source.Raw); i++ {
			byteList[i] = source.Raw[i]
		}
	}
	v1JSON.Raw = byteList
	return v1JSON
}
func (c *GeneratedRevisionSpecConverter) v1MatchTransformPatternToV1MatchTransformPattern(source MatchTransformPattern) MatchTransformPattern {
	var v1MatchTransformPattern MatchTransformPattern
	v1MatchTransformPattern.Type = MatchTransformPatternType(source.Type)
	var pString *string
	if source.Literal != nil {
		xstring := *source.Literal
		pString = &xstring
	}
	v1MatchTransformPattern.Literal = pString
	var pString2 *string
	if source.Regexp != nil {
		xstring2 := *source.Regexp
		pString2 = &xstring2
	}
	v1MatchTransformPattern.Regexp = pString2
	v1MatchTransformPattern.Result = c.v1JSONToV1JSON(source.Result)
	return v1MatchTransformPattern
}
func (c *GeneratedRevisionSpecConverter) v1PatchSetToV1PatchSet(source PatchSet) PatchSet {
	var v1PatchSet PatchSet
	v1PatchSet.Name = source.Name
	var v1PatchList []Patch
	if source.Patches != nil {
		v1PatchList = make([]Patch, len(source.Patches))
		for i := 0; i < len(source.Patches); i++ {
			v1PatchList[i] = c.v1PatchToV1Patch(source.Patches[i])
		}
	}
	v1PatchSet.Patches = v1PatchList
	return v1PatchSet
}
func (c *GeneratedRevisionSpecConverter) v1PatchToV1Patch(source Patch) Patch {
	var v1Patch Patch
	v1Patch.Type = PatchType(source.Type)
	var pString *string
	if source.FromFieldPath != nil {
		xstring := *source.FromFieldPath
		pString = &xstring
	}
	v1Patch.FromFieldPath = pString
	v1Patch.Combine = c.pV1CombineToPV1Combine(source.Combine)
	var pString2 *string
	if source.ToFieldPath != nil {
		xstring2 := *source.ToFieldPath
		pString2 = &xstring2
	}
	v1Patch.ToFieldPath = pString2
	var pString3 *string
	if source.PatchSetName != nil {
		xstring3 := *source.PatchSetName
		pString3 = &xstring3
	}
	v1Patch.PatchSetName = pString3
	var v1TransformList []Transform
	if source.Transforms != nil {
		v1TransformList = make([]Transform, len(source.Transforms))
		for i := 0; i < len(source.Transforms); i++ {
			v1TransformList[i] = c.v1TransformToV1Transform(source.Transforms[i])
		}
	}
	v1Patch.Transforms = v1TransformList
	v1Patch.Policy = c.pV1PatchPolicyToPV1PatchPolicy(source.Policy)
	return v1Patch
}
func (c *GeneratedRevisionSpecConverter) v1PipelineStepToV1PipelineStep(source PipelineStep) PipelineStep {
	var v1PipelineStep PipelineStep
	v1PipelineStep.Step = source.Step
	v1PipelineStep.FunctionRef = c.v1FunctionReferenceToV1FunctionReference(source.FunctionRef)
	v1PipelineStep.Input = c.pRuntimeRawExtensionToPRuntimeRawExtension(source.Input)
	var v1FunctionCredentialsList []FunctionCredentials
	if source.Credentials != nil {
		v1FunctionCredentialsList = make([]FunctionCredentials, len(source.Credentials))
		for i := 0; i < len(source.Credentials); i++ {
			v1FunctionCredentialsList[i] = c.v1FunctionCredentialsToV1FunctionCredentials(source.Credentials[i])
		}
	}
	v1PipelineStep.Credentials = v1FunctionCredentialsList
	return v1PipelineStep
}
func (c *GeneratedRevisionSpecConverter) v1ReadinessCheckToV1ReadinessCheck(source ReadinessCheck) ReadinessCheck {
	var v1ReadinessCheck ReadinessCheck
	v1ReadinessCheck.Type = ReadinessCheckType(source.Type)
	v1ReadinessCheck.FieldPath = source.FieldPath
	v1ReadinessCheck.MatchString = source.MatchString
	v1ReadinessCheck.MatchInteger = source.MatchInteger
	v1ReadinessCheck.MatchCondition = c.pV1MatchConditionReadinessCheckToPV1MatchConditionReadinessCheck(source.MatchCondition)
	return v1ReadinessCheck
}
func (c *GeneratedRevisionSpecConverter) v1TransformToV1Transform(source Transform) Transform {
	var v1Transform Transform
	v1Transform.Type = TransformType(source.Type)
	v1Transform.Math = c.pV1MathTransformToPV1MathTransform(source.Math)
	v1Transform.Map = c.pV1MapTransformToPV1MapTransform(source.Map)
	v1Transform.Match = c.pV1MatchTransformToPV1MatchTransform(source.Match)
	v1Transform.String = c.pV1StringTransformToPV1StringTransform(source.String)
	v1Transform.Convert = c.pV1ConvertTransformToPV1ConvertTransform(source.Convert)
	return v1Transform
}
func (c *GeneratedRevisionSpecConverter) v1TypeReferenceToV1TypeReference(source TypeReference) TypeReference {
	var v1TypeReference TypeReference
	v1TypeReference.APIVersion = source.APIVersion
	v1TypeReference.Kind = source.Kind
	return v1TypeReference
}
