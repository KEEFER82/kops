/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/values"
)

// Convert_v1alpha2_CanalNetworkingSpec_To_kops_CanalNetworkingSpec is an autogenerated conversion function.
func Convert_v1alpha2_CanalNetworkingSpec_To_kops_CanalNetworkingSpec(in *CanalNetworkingSpec, out *kops.CanalNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_CanalNetworkingSpec_To_kops_CanalNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.FlanneldIptablesForwardRules != nil {
		out.FlanneldIptablesForwardRules = values.Bool(!*in.FlanneldIptablesForwardRules)
	}
	return nil
}

// Convert_kops_CanalNetworkingSpec_To_v1alpha2_CanalNetworkingSpec is an autogenerated conversion function.
func Convert_kops_CanalNetworkingSpec_To_v1alpha2_CanalNetworkingSpec(in *kops.CanalNetworkingSpec, out *CanalNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_kops_CanalNetworkingSpec_To_v1alpha2_CanalNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.FlanneldIptablesForwardRules != nil {
		out.FlanneldIptablesForwardRules = values.Bool(!*in.FlanneldIptablesForwardRules)
	}
	return nil
}

func Convert_v1alpha2_CiliumNetworkingSpec_To_kops_CiliumNetworkingSpec(in *CiliumNetworkingSpec, out *kops.CiliumNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_CiliumNetworkingSpec_To_kops_CiliumNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.InstallIptablesRules != nil {
		out.InstallIptablesRules = values.Bool(!*in.InstallIptablesRules)
	}
	if in.Masquerade != nil {
		out.Masquerade = values.Bool(!*in.Masquerade)
	}
	return nil
}

func Convert_kops_CiliumNetworkingSpec_To_v1alpha2_CiliumNetworkingSpec(in *kops.CiliumNetworkingSpec, out *CiliumNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_kops_CiliumNetworkingSpec_To_v1alpha2_CiliumNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.InstallIptablesRules != nil {
		out.InstallIptablesRules = values.Bool(!*in.InstallIptablesRules)
	}
	if in.Masquerade != nil {
		out.Masquerade = values.Bool(!*in.Masquerade)
	}
	return nil
}

func Convert_v1alpha2_ClusterSpec_To_kops_ClusterSpec(in *ClusterSpec, out *kops.ClusterSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_ClusterSpec_To_kops_ClusterSpec(in, out, s); err != nil {
		return err
	}
	switch kops.CloudProviderID(in.LegacyCloudProvider) {
	case kops.CloudProviderAWS:
		out.CloudProvider.AWS = &kops.AWSSpec{}
	case kops.CloudProviderAzure:
		out.CloudProvider.Azure = &kops.AzureSpec{}
		if in.CloudConfig != nil && in.CloudConfig.Azure != nil {
			if err := autoConvert_v1alpha2_AzureSpec_To_kops_AzureSpec(in.CloudConfig.Azure, out.CloudProvider.Azure, s); err != nil {
				return err
			}
		}
	case kops.CloudProviderDO:
		out.CloudProvider.DO = &kops.DOSpec{}
	case kops.CloudProviderGCE:
		out.CloudProvider.GCE = &kops.GCESpec{}
	case kops.CloudProviderOpenstack:
		out.CloudProvider.Openstack = &kops.OpenstackSpec{}
	case "":
	default:
		return field.NotSupported(field.NewPath("spec").Child("cloudProvider"), in.LegacyCloudProvider, []string{
			string(kops.CloudProviderGCE),
			string(kops.CloudProviderDO),
			string(kops.CloudProviderAzure),
			string(kops.CloudProviderAWS),
			string(kops.CloudProviderOpenstack),
		})
	}
	if in.TagSubnets != nil {
		out.TagSubnets = values.Bool(!*in.TagSubnets)
	}
	for i, hook := range in.Hooks {
		if hook.Enabled != nil {
			out.Hooks[i].Enabled = values.Bool(!*hook.Enabled)
		}
	}
	return nil
}

func Convert_kops_ClusterSpec_To_v1alpha2_ClusterSpec(in *kops.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	if err := autoConvert_kops_ClusterSpec_To_v1alpha2_ClusterSpec(in, out, s); err != nil {
		return err
	}
	out.LegacyCloudProvider = string(in.GetCloudProvider())
	switch kops.CloudProviderID(out.LegacyCloudProvider) {
	case kops.CloudProviderAzure:
		if out.CloudConfig == nil {
			out.CloudConfig = &CloudConfiguration{}
		}
		if out.CloudConfig.Azure == nil {
			out.CloudConfig.Azure = &AzureSpec{}
		}
		if err := autoConvert_kops_AzureSpec_To_v1alpha2_AzureSpec(in.CloudProvider.Azure, out.CloudConfig.Azure, s); err != nil {
			return err
		}
	}
	if in.TagSubnets != nil {
		out.TagSubnets = values.Bool(!*in.TagSubnets)
	}
	for i, hook := range in.Hooks {
		if hook.Enabled != nil {
			out.Hooks[i].Enabled = values.Bool(!*hook.Enabled)
		}
	}
	return nil
}

func Convert_v1alpha2_ExternalDNSConfig_To_kops_ExternalDNSConfig(in *ExternalDNSConfig, out *kops.ExternalDNSConfig, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_ExternalDNSConfig_To_kops_ExternalDNSConfig(in, out, s); err != nil {
		return err
	}
	if in.Disable {
		out.Provider = kops.ExternalDNSProviderNone
	}
	return nil
}

func Convert_kops_ExternalDNSConfig_To_v1alpha2_ExternalDNSConfig(in *kops.ExternalDNSConfig, out *ExternalDNSConfig, s conversion.Scope) error {
	if err := autoConvert_kops_ExternalDNSConfig_To_v1alpha2_ExternalDNSConfig(in, out, s); err != nil {
		return err
	}
	if in.Provider == kops.ExternalDNSProviderNone {
		out.Disable = true
		out.Provider = ""
	}
	return nil
}
