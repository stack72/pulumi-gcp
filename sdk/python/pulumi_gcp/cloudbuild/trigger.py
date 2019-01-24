# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Trigger(pulumi.CustomResource):
    build: pulumi.Output[dict]
    """
    A build resource in the Container Builder API.
    Structure is documented below. At a high
    level, a `build` describes where to find source code, how to build it (for
    example, the builder image to run on the source), and where to store
    the built artifacts. Fields can include the following variables, which
    will be expanded when the build is created:
    * `$PROJECT_ID`: the project ID of the build.
    * `$BUILD_ID`: the autogenerated ID of the build.
    * `$REPO_NAME`: the source repository name specified by RepoSource.
    * `$BRANCH_NAME`: the branch name specified by RepoSource.
    * `$TAG_NAME`: the tag name specified by RepoSource.
    * `$REVISION_ID` or `$COMMIT_SHA`: the commit SHA specified by RepoSource
    or resolved from the specified branch or tag.
    * `$SHORT_SHA`: first 7 characters of `$REVISION_ID` or `$COMMIT_SHA`.
    """
    create_time: pulumi.Output[str]
    description: pulumi.Output[str]
    """
    A brief description of this resource.
    """
    disabled: pulumi.Output[str]
    filename: pulumi.Output[str]
    """
    Specify the path to a Cloud Build configuration file
    in the Git repo. This is mutually exclusive with `build`. This is typically
    `cloudbuild.yaml` however it can be specified by the user.
    """
    ignored_files: pulumi.Output[list]
    included_files: pulumi.Output[list]
    project: pulumi.Output[str]
    """
    The ID of the project that the trigger will be created in.
    Defaults to the provider project configuration.
    """
    substitutions: pulumi.Output[dict]
    trigger_id: pulumi.Output[str]
    trigger_template: pulumi.Output[dict]
    """
    Location of the source in a Google
    Cloud Source Repository. Structure is documented below.
    """
    def __init__(__self__, __name__, __opts__=None, build=None, description=None, disabled=None, filename=None, ignored_files=None, included_files=None, project=None, substitutions=None, trigger_template=None):
        """
        Creates a new build trigger within GCR. For more information, see
        [the official documentation](https://cloud.google.com/container-builder/docs/running-builds/automate-builds)
        and
        [API](https://godoc.org/google.golang.org/api/cloudbuild/v1#BuildTrigger).
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] build: A build resource in the Container Builder API.
               Structure is documented below. At a high
               level, a `build` describes where to find source code, how to build it (for
               example, the builder image to run on the source), and where to store
               the built artifacts. Fields can include the following variables, which
               will be expanded when the build is created:
               * `$PROJECT_ID`: the project ID of the build.
               * `$BUILD_ID`: the autogenerated ID of the build.
               * `$REPO_NAME`: the source repository name specified by RepoSource.
               * `$BRANCH_NAME`: the branch name specified by RepoSource.
               * `$TAG_NAME`: the tag name specified by RepoSource.
               * `$REVISION_ID` or `$COMMIT_SHA`: the commit SHA specified by RepoSource
               or resolved from the specified branch or tag.
               * `$SHORT_SHA`: first 7 characters of `$REVISION_ID` or `$COMMIT_SHA`.
        :param pulumi.Input[str] description: A brief description of this resource.
        :param pulumi.Input[str] disabled
        :param pulumi.Input[str] filename: Specify the path to a Cloud Build configuration file
               in the Git repo. This is mutually exclusive with `build`. This is typically
               `cloudbuild.yaml` however it can be specified by the user.
        :param pulumi.Input[list] ignored_files
        :param pulumi.Input[list] included_files
        :param pulumi.Input[str] project: The ID of the project that the trigger will be created in.
               Defaults to the provider project configuration.
        :param pulumi.Input[dict] substitutions
        :param pulumi.Input[dict] trigger_template: Location of the source in a Google
               Cloud Source Repository. Structure is documented below.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['build'] = build

        __props__['description'] = description

        __props__['disabled'] = disabled

        __props__['filename'] = filename

        __props__['ignored_files'] = ignored_files

        __props__['included_files'] = included_files

        __props__['project'] = project

        __props__['substitutions'] = substitutions

        __props__['trigger_template'] = trigger_template

        __props__['create_time'] = None
        __props__['trigger_id'] = None

        super(Trigger, __self__).__init__(
            'gcp:cloudbuild/trigger:Trigger',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

