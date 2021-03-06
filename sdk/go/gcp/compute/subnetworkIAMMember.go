// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Warning:** These resources are in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
// 
// Three different resources help you manage your IAM policy for GCE subnetwork. Each of these resources serves a different use case:
// 
// * `google_compute_subnetwork_iam_policy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
// * `google_compute_subnetwork_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
// * `google_compute_subnetwork_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.
// 
// > **Note:** `google_compute_subnetwork_iam_policy` **cannot** be used in conjunction with `google_compute_subnetwork_iam_binding` and `google_compute_subnetwork_iam_member` or they will fight over what your policy should be.
// 
// > **Note:** `google_compute_subnetwork_iam_binding` resources **can be** used in conjunction with `google_compute_subnetwork_iam_member` resources **only if** they do not grant privilege to the same role.
type SubnetworkIAMMember struct {
	s *pulumi.ResourceState
}

// NewSubnetworkIAMMember registers a new resource with the given unique name, arguments, and options.
func NewSubnetworkIAMMember(ctx *pulumi.Context,
	name string, args *SubnetworkIAMMemberArgs, opts ...pulumi.ResourceOpt) (*SubnetworkIAMMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Subnetwork == nil {
		return nil, errors.New("missing required argument 'Subnetwork'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["member"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["role"] = nil
		inputs["subnetwork"] = nil
	} else {
		inputs["member"] = args.Member
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["role"] = args.Role
		inputs["subnetwork"] = args.Subnetwork
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:compute/subnetworkIAMMember:SubnetworkIAMMember", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetworkIAMMember{s: s}, nil
}

// GetSubnetworkIAMMember gets an existing SubnetworkIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetworkIAMMember(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetworkIAMMemberState, opts ...pulumi.ResourceOpt) (*SubnetworkIAMMember, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["member"] = state.Member
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["role"] = state.Role
		inputs["subnetwork"] = state.Subnetwork
	}
	s, err := ctx.ReadResource("gcp:compute/subnetworkIAMMember:SubnetworkIAMMember", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SubnetworkIAMMember{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SubnetworkIAMMember) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SubnetworkIAMMember) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the subnetwork's IAM policy.
func (r *SubnetworkIAMMember) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *SubnetworkIAMMember) Member() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["member"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *SubnetworkIAMMember) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The region of the subnetwork. If
// unspecified, this defaults to the region configured in the provider.
func (r *SubnetworkIAMMember) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

// The role that should be applied. Only one
// `google_compute_subnetwork_iam_binding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *SubnetworkIAMMember) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// The name of the subnetwork.
func (r *SubnetworkIAMMember) Subnetwork() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["subnetwork"])
}

// Input properties used for looking up and filtering SubnetworkIAMMember resources.
type SubnetworkIAMMemberState struct {
	// (Computed) The etag of the subnetwork's IAM policy.
	Etag interface{}
	Member interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region of the subnetwork. If
	// unspecified, this defaults to the region configured in the provider.
	Region interface{}
	// The role that should be applied. Only one
	// `google_compute_subnetwork_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The name of the subnetwork.
	Subnetwork interface{}
}

// The set of arguments for constructing a SubnetworkIAMMember resource.
type SubnetworkIAMMemberArgs struct {
	Member interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The region of the subnetwork. If
	// unspecified, this defaults to the region configured in the provider.
	Region interface{}
	// The role that should be applied. Only one
	// `google_compute_subnetwork_iam_binding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// The name of the subnetwork.
	Subnetwork interface{}
}
