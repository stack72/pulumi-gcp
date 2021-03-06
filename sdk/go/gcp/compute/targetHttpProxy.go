// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a TargetHttpProxy resource, which is used by one or more global
// forwarding rule to route incoming HTTP requests to a URL map.
// 
// 
// To get more information about TargetHttpProxy, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpProxies)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=target_http_proxy_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type TargetHttpProxy struct {
	s *pulumi.ResourceState
}

// NewTargetHttpProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpProxy(ctx *pulumi.Context,
	name string, args *TargetHttpProxyArgs, opts ...pulumi.ResourceOpt) (*TargetHttpProxy, error) {
	if args == nil || args.UrlMap == nil {
		return nil, errors.New("missing required argument 'UrlMap'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["urlMap"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["urlMap"] = args.UrlMap
	}
	inputs["creationTimestamp"] = nil
	inputs["proxyId"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/targetHttpProxy:TargetHttpProxy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetHttpProxy{s: s}, nil
}

// GetTargetHttpProxy gets an existing TargetHttpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpProxy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TargetHttpProxyState, opts ...pulumi.ResourceOpt) (*TargetHttpProxy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["proxyId"] = state.ProxyId
		inputs["selfLink"] = state.SelfLink
		inputs["urlMap"] = state.UrlMap
	}
	s, err := ctx.ReadResource("gcp:compute/targetHttpProxy:TargetHttpProxy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetHttpProxy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TargetHttpProxy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TargetHttpProxy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *TargetHttpProxy) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *TargetHttpProxy) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *TargetHttpProxy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *TargetHttpProxy) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *TargetHttpProxy) ProxyId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["proxyId"])
}

// The URI of the created resource.
func (r *TargetHttpProxy) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *TargetHttpProxy) UrlMap() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["urlMap"])
}

// Input properties used for looking up and filtering TargetHttpProxy resources.
type TargetHttpProxyState struct {
	CreationTimestamp interface{}
	Description interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	ProxyId interface{}
	// The URI of the created resource.
	SelfLink interface{}
	UrlMap interface{}
}

// The set of arguments for constructing a TargetHttpProxy resource.
type TargetHttpProxyArgs struct {
	Description interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	UrlMap interface{}
}
