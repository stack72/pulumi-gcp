// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of enabled API services for an existing Google Cloud
// Platform project. Services in an existing project that are not defined
// in the config will be removed.
// 
// For a list of services available, visit the
// [API library page](https://console.cloud.google.com/apis/library) or run `gcloud services list`.
// 
// ~> **Note:** This resource attempts to be the authoritative source on *all* enabled APIs, which often
// 	leads to conflicts when certain actions enable other APIs. If you do not need to ensure that
// 	*exclusively* a particular set of APIs are enabled, you should most likely use the
// 	google_project_service resource, one resource per API.
type Services struct {
	s *pulumi.ResourceState
}

// NewServices registers a new resource with the given unique name, arguments, and options.
func NewServices(ctx *pulumi.Context,
	name string, args *ServicesArgs, opts ...pulumi.ResourceOpt) (*Services, error) {
	if args == nil || args.Services == nil {
		return nil, errors.New("missing required argument 'Services'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["disableOnDestroy"] = nil
		inputs["project"] = nil
		inputs["services"] = nil
	} else {
		inputs["disableOnDestroy"] = args.DisableOnDestroy
		inputs["project"] = args.Project
		inputs["services"] = args.Services
	}
	s, err := ctx.RegisterResource("gcp:projects/services:Services", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Services{s: s}, nil
}

// GetServices gets an existing Services resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServices(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ServicesState, opts ...pulumi.ResourceOpt) (*Services, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["disableOnDestroy"] = state.DisableOnDestroy
		inputs["project"] = state.Project
		inputs["services"] = state.Services
	}
	s, err := ctx.ReadResource("gcp:projects/services:Services", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Services{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Services) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Services) ID() *pulumi.IDOutput {
	return r.s.ID
}

func (r *Services) DisableOnDestroy() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["disableOnDestroy"])
}

// The project ID.
// Changing this forces Terraform to attempt to disable all previously managed
// API services in the previous project.
func (r *Services) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The list of services that are enabled. Supports
// update.
func (r *Services) Services() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["services"])
}

// Input properties used for looking up and filtering Services resources.
type ServicesState struct {
	DisableOnDestroy interface{}
	// The project ID.
	// Changing this forces Terraform to attempt to disable all previously managed
	// API services in the previous project.
	Project interface{}
	// The list of services that are enabled. Supports
	// update.
	Services interface{}
}

// The set of arguments for constructing a Services resource.
type ServicesArgs struct {
	DisableOnDestroy interface{}
	// The project ID.
	// Changing this forces Terraform to attempt to disable all previously managed
	// API services in the previous project.
	Project interface{}
	// The list of services that are enabled. Supports
	// update.
	Services interface{}
}
