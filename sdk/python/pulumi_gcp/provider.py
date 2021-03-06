# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Provider(pulumi.ProviderResource):
    def __init__(__self__, resource_name, opts=None, access_token=None, credentials=None, project=None, region=None, scopes=None, zone=None, __name__=None, __opts__=None):
        """
        The provider type for the google-beta package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://pulumi.io/reference/programming-model.html#providers) for more information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['access_token'] = access_token

        if credentials is None:
            credentials = utilities.get_env('GOOGLE_CREDENTIALS', 'GOOGLE_CLOUD_KEYFILE_JSON', 'GCLOUD_KEYFILE_JSON')
        __props__['credentials'] = credentials

        if project is None:
            project = utilities.get_env('GOOGLE_PROJECT', 'GOOGLE_CLOUD_PROJECT', 'GCLOUD_PROJECT', 'CLOUDSDK_CORE_PROJECT')
        __props__['project'] = project

        if region is None:
            region = utilities.get_env('GOOGLE_REGION', 'GCLOUD_REGION', 'CLOUDSDK_COMPUTE_REGION')
        __props__['region'] = region

        __props__['scopes'] = pulumi.Output.from_input(scopes).apply(json.dumps) if scopes is not None else None

        if zone is None:
            zone = utilities.get_env('GOOGLE_ZONE', 'GCLOUD_ZONE', 'CLOUDSDK_COMPUTE_ZONE')
        __props__['zone'] = zone

        super(Provider, __self__).__init__(
            'gcp',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

