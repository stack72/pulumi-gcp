# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Account(pulumi.CustomResource):
    account_id: pulumi.Output[str]
    """
    The account id that is used to generate the service
    account email address and a stable unique id. It is unique within a project,
    must be 6-30 characters long, and match the regular expression `a-z`
    to comply with RFC1035. Changing this forces a new service account to be created.
    """
    display_name: pulumi.Output[str]
    """
    The display name for the service account.
    Can be updated without creating a new resource.
    """
    email: pulumi.Output[str]
    """
    The e-mail address of the service account. This value
    should be referenced from any `google_iam_policy` data sources
    that would grant the service account privileges.
    """
    name: pulumi.Output[str]
    """
    The fully-qualified name of the service account.
    """
    project: pulumi.Output[str]
    """
    The ID of the project that the service account will be created in.
    Defaults to the provider project configuration.
    """
    unique_id: pulumi.Output[str]
    """
    The unique id of the service account.
    """
    def __init__(__self__, resource_name, opts=None, account_id=None, display_name=None, project=None, __name__=None, __opts__=None):
        """
        Allows management of a [Google Cloud Platform service account](https://cloud.google.com/compute/docs/access/service-accounts)
        
        > Creation of service accounts is eventually consistent, and that can lead to
        errors when you try to apply ACLs to service accounts immediately after
        creation. If using these resources in the same config, you can add a
        [`sleep` using `local-exec`](https://github.com/hashicorp/terraform/issues/17726#issuecomment-377357866).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The account id that is used to generate the service
               account email address and a stable unique id. It is unique within a project,
               must be 6-30 characters long, and match the regular expression `a-z`
               to comply with RFC1035. Changing this forces a new service account to be created.
        :param pulumi.Input[str] display_name: The display name for the service account.
               Can be updated without creating a new resource.
        :param pulumi.Input[str] project: The ID of the project that the service account will be created in.
               Defaults to the provider project configuration.
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

        if account_id is None:
            raise TypeError('Missing required property account_id')
        __props__['account_id'] = account_id

        __props__['display_name'] = display_name

        __props__['project'] = project

        __props__['email'] = None
        __props__['name'] = None
        __props__['unique_id'] = None

        super(Account, __self__).__init__(
            'gcp:serviceAccount/account:Account',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

