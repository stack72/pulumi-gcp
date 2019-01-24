# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetAccountKeyResult(object):
    """
    A collection of values returned by getAccountKey.
    """
    def __init__(__self__, key_algorithm=None, public_key=None, id=None):
        if key_algorithm and not isinstance(key_algorithm, str):
            raise TypeError('Expected argument key_algorithm to be a str')
        __self__.key_algorithm = key_algorithm
        if public_key and not isinstance(public_key, str):
            raise TypeError('Expected argument public_key to be a str')
        __self__.public_key = public_key
        """
        The public key, base64 encoded
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_account_key(name=None, project=None, public_key_type=None):
    """
    Get service account public key. For more information, see [the official documentation](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) and [API](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys/get).
    
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['project'] = project
    __args__['publicKeyType'] = public_key_type
    __ret__ = await pulumi.runtime.invoke('gcp:serviceAccount/getAccountKey:getAccountKey', __args__)

    return GetAccountKeyResult(
        key_algorithm=__ret__.get('keyAlgorithm'),
        public_key=__ret__.get('publicKey'),
        id=__ret__.get('id'))
