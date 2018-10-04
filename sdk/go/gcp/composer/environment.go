// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package composer

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Environment struct {
	s *pulumi.ResourceState
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOpt) (*Environment, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["config"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
	} else {
		inputs["config"] = args.Config
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	s, err := ctx.RegisterResource("gcp:composer/environment:Environment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Environment{s: s}, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *EnvironmentState, opts ...pulumi.ResourceOpt) (*Environment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["config"] = state.Config
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("gcp:composer/environment:Environment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Environment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Environment) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Environment) ID() *pulumi.IDOutput {
	return r.s.ID
}

func (r *Environment) Config() *pulumi.Output {
	return r.s.State["config"]
}

func (r *Environment) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *Environment) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Environment) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Environment) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering Environment resources.
type EnvironmentState struct {
	Config interface{}
	Labels interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	Config interface{}
	Labels interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
}
