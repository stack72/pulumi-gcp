# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RegionInstanceGroupManager(pulumi.CustomResource):
    auto_healing_policies: pulumi.Output[dict]
    """
    The autohealing policies for this managed instance
    group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    """
    base_instance_name: pulumi.Output[str]
    """
    The base instance name to use for
    instances in this group. The value must be a valid
    [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
    are lowercase letters, numbers, and hyphens (-). Instances are named by
    appending a hyphen and a random four-character string to the base instance
    name.
    """
    description: pulumi.Output[str]
    """
    An optional textual description of the instance
    group manager.
    """
    distribution_policy_zones: pulumi.Output[list]
    """
    The distribution policy for this managed instance
    group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
    - - -
    """
    fingerprint: pulumi.Output[str]
    """
    The fingerprint of the instance group manager.
    """
    instance_group: pulumi.Output[str]
    """
    The full URL of the instance group created by the manager.
    """
    name: pulumi.Output[str]
    """
    - Version name.
    """
    named_ports: pulumi.Output[list]
    """
    The named port configuration. See the section below
    for details on configuration.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    The region where the managed instance group resides.
    """
    self_link: pulumi.Output[str]
    """
    The URL of the created resource.
    """
    target_pools: pulumi.Output[list]
    """
    The full URL of all target pools to which new
    instances in the group are added. Updating the target pools attribute does
    not affect existing instances.
    """
    target_size: pulumi.Output[int]
    """
    - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
    """
    update_policy: pulumi.Output[dict]
    versions: pulumi.Output[list]
    """
    Application versions managed by this instance group. Each
    version deals with a specific instance template, allowing canary release scenarios.
    Conflicts with `instance_template`. Structure is documented below. Beware that
    exactly one version must not specify a target size. It means that versions with
    a target size will respect the setting, and the one without target size will
    be applied to all remaining Instances (top level target_size - each version target_size).
    This property is in beta, and should be used with the terraform-provider-google-beta provider.
    See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
    """
    wait_for_instances: pulumi.Output[bool]
    """
    Whether to wait for all instances to be created/updated before
    returning. Note that if this is set to true and the operation does not succeed, Terraform will
    continue trying until it times out.
    """
    def __init__(__self__, __name__, __opts__=None, auto_healing_policies=None, base_instance_name=None, description=None, distribution_policy_zones=None, name=None, named_ports=None, project=None, region=None, target_pools=None, target_size=None, update_policy=None, versions=None, wait_for_instances=None):
        """
        The Google Compute Engine Regional Instance Group Manager API creates and manages pools
        of homogeneous Compute Engine virtual machine instances from a common instance
        template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups)
        and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroupManagers)
        
        > **Note:** Use [google_compute_instance_group_manager](https://www.terraform.io/docs/providers/google/r/compute_instance_group_manager.html) to create a single-zone instance group manager.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] auto_healing_policies: The autohealing policies for this managed instance
               group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
        :param pulumi.Input[str] base_instance_name: The base instance name to use for
               instances in this group. The value must be a valid
               [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
               are lowercase letters, numbers, and hyphens (-). Instances are named by
               appending a hyphen and a random four-character string to the base instance
               name.
        :param pulumi.Input[str] description: An optional textual description of the instance
               group manager.
        :param pulumi.Input[list] distribution_policy_zones: The distribution policy for this managed instance
               group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones).
               - - -
        :param pulumi.Input[str] name: - Version name.
        :param pulumi.Input[list] named_ports: The named port configuration. See the section below
               for details on configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region where the managed instance group resides.
        :param pulumi.Input[list] target_pools: The full URL of all target pools to which new
               instances in the group are added. Updating the target pools attribute does
               not affect existing instances.
        :param pulumi.Input[int] target_size: - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
        :param pulumi.Input[dict] update_policy
        :param pulumi.Input[list] versions: Application versions managed by this instance group. Each
               version deals with a specific instance template, allowing canary release scenarios.
               Conflicts with `instance_template`. Structure is documented below. Beware that
               exactly one version must not specify a target size. It means that versions with
               a target size will respect the setting, and the one without target size will
               be applied to all remaining Instances (top level target_size - each version target_size).
               This property is in beta, and should be used with the terraform-provider-google-beta provider.
               See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
        :param pulumi.Input[bool] wait_for_instances: Whether to wait for all instances to be created/updated before
               returning. Note that if this is set to true and the operation does not succeed, Terraform will
               continue trying until it times out.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['auto_healing_policies'] = auto_healing_policies

        if not base_instance_name:
            raise TypeError('Missing required property base_instance_name')
        __props__['base_instance_name'] = base_instance_name

        __props__['description'] = description

        __props__['distribution_policy_zones'] = distribution_policy_zones

        __props__['name'] = name

        __props__['named_ports'] = named_ports

        __props__['project'] = project

        if not region:
            raise TypeError('Missing required property region')
        __props__['region'] = region

        __props__['target_pools'] = target_pools

        __props__['target_size'] = target_size

        __props__['update_policy'] = update_policy

        if not versions:
            raise TypeError('Missing required property versions')
        __props__['versions'] = versions

        __props__['wait_for_instances'] = wait_for_instances

        __props__['fingerprint'] = None
        __props__['instance_group'] = None
        __props__['self_link'] = None

        super(RegionInstanceGroupManager, __self__).__init__(
            'gcp:compute/regionInstanceGroupManager:RegionInstanceGroupManager',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

