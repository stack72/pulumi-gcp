# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMBinding(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    (Computed) The etag of the organization's IAM policy.
    """
    members: pulumi.Output[list]
    """
    A list of users that the role should apply to.
    """
    org_id: pulumi.Output[str]
    """
    The numeric ID of the organization in which you want to create a custom role.
    """
    role: pulumi.Output[str]
    """
    The role that should be applied. Only one
    `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
    `[projects|organizations]/{parent-name}/roles/{role-name}`.
    """
    def __init__(__self__, __name__, __opts__=None, members=None, org_id=None, role=None):
        """
        Allows creation and management of a single binding within IAM policy for
        an existing Google Cloud Platform Organization.
        
        > **Note:** This resource __must not__ be used in conjunction with
           `google_organization_iam_member` for the __same role__ or they will fight over
           what your policy should be.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] members: A list of users that the role should apply to.
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `google_organization_iam_binding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not members:
            raise TypeError('Missing required property members')
        __props__['members'] = members

        if not org_id:
            raise TypeError('Missing required property org_id')
        __props__['org_id'] = org_id

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        __props__['etag'] = None

        super(IAMBinding, __self__).__init__(
            'gcp:organizations/iAMBinding:IAMBinding',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

