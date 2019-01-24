# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CryptoKey(pulumi.CustomResource):
    key_ring: pulumi.Output[str]
    """
    The id of the Google Cloud Platform KeyRing to which the key shall belong.
    """
    name: pulumi.Output[str]
    """
    The CryptoKey's name.
    A CryptoKey’s name must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
    """
    rotation_period: pulumi.Output[str]
    """
    Every time this period passes, generate a new CryptoKeyVersion and set it as
    the primary. The first rotation will take place after the specified period. The rotation period has the format
    of a decimal number with up to 9 fractional digits, followed by the letter s (seconds). It must be greater than
    a day (ie, 86400).
    """
    self_link: pulumi.Output[str]
    """
    The self link of the created CryptoKey. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}/cryptoKeys/{cryptoKeyName}`.
    """
    version_template: pulumi.Output[dict]
    def __init__(__self__, __name__, __opts__=None, key_ring=None, name=None, rotation_period=None, version_template=None):
        """
        Allows creation of a Google Cloud Platform KMS CryptoKey. For more information see
        [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#cryptokey)
        and
        [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys).
        
        A CryptoKey is an interface to key material which can be used to encrypt and decrypt data. A CryptoKey belongs to a
        Google Cloud KMS KeyRing.
        
        > Note: CryptoKeys cannot be deleted from Google Cloud Platform. Destroying a
        Terraform-managed CryptoKey will remove it from state and delete all
        CryptoKeyVersions, rendering the key unusable, but **will not delete the
        resource on the server**. When Terraform destroys these keys, any data
        previously encrypted with these keys will be irrecoverable. For this reason, it
        is strongly recommended that you add lifecycle hooks to the resource to prevent
        accidental destruction.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] key_ring: The id of the Google Cloud Platform KeyRing to which the key shall belong.
        :param pulumi.Input[str] name: The CryptoKey's name.
               A CryptoKey’s name must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
        :param pulumi.Input[str] rotation_period: Every time this period passes, generate a new CryptoKeyVersion and set it as
               the primary. The first rotation will take place after the specified period. The rotation period has the format
               of a decimal number with up to 9 fractional digits, followed by the letter s (seconds). It must be greater than
               a day (ie, 86400).
        :param pulumi.Input[dict] version_template
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not key_ring:
            raise TypeError('Missing required property key_ring')
        __props__['key_ring'] = key_ring

        __props__['name'] = name

        __props__['rotation_period'] = rotation_period

        __props__['version_template'] = version_template

        __props__['self_link'] = None

        super(CryptoKey, __self__).__init__(
            'gcp:kms/cryptoKey:CryptoKey',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

