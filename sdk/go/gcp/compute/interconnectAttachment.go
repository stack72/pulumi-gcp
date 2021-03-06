// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents an InterconnectAttachment (VLAN attachment) resource. For more
// information, see Creating VLAN Attachments.
type InterconnectAttachment struct {
	s *pulumi.ResourceState
}

// NewInterconnectAttachment registers a new resource with the given unique name, arguments, and options.
func NewInterconnectAttachment(ctx *pulumi.Context,
	name string, args *InterconnectAttachmentArgs, opts ...pulumi.ResourceOpt) (*InterconnectAttachment, error) {
	if args == nil || args.Router == nil {
		return nil, errors.New("missing required argument 'Router'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["candidateSubnets"] = nil
		inputs["description"] = nil
		inputs["edgeAvailabilityDomain"] = nil
		inputs["interconnect"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["router"] = nil
		inputs["type"] = nil
		inputs["vlanTag8021q"] = nil
	} else {
		inputs["candidateSubnets"] = args.CandidateSubnets
		inputs["description"] = args.Description
		inputs["edgeAvailabilityDomain"] = args.EdgeAvailabilityDomain
		inputs["interconnect"] = args.Interconnect
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["router"] = args.Router
		inputs["type"] = args.Type
		inputs["vlanTag8021q"] = args.VlanTag8021q
	}
	inputs["cloudRouterIpAddress"] = nil
	inputs["creationTimestamp"] = nil
	inputs["customerRouterIpAddress"] = nil
	inputs["googleReferenceId"] = nil
	inputs["pairingKey"] = nil
	inputs["partnerAsn"] = nil
	inputs["privateInterconnectInfo"] = nil
	inputs["selfLink"] = nil
	inputs["state"] = nil
	s, err := ctx.RegisterResource("gcp:compute/interconnectAttachment:InterconnectAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InterconnectAttachment{s: s}, nil
}

// GetInterconnectAttachment gets an existing InterconnectAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterconnectAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InterconnectAttachmentState, opts ...pulumi.ResourceOpt) (*InterconnectAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["candidateSubnets"] = state.CandidateSubnets
		inputs["cloudRouterIpAddress"] = state.CloudRouterIpAddress
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["customerRouterIpAddress"] = state.CustomerRouterIpAddress
		inputs["description"] = state.Description
		inputs["edgeAvailabilityDomain"] = state.EdgeAvailabilityDomain
		inputs["googleReferenceId"] = state.GoogleReferenceId
		inputs["interconnect"] = state.Interconnect
		inputs["name"] = state.Name
		inputs["pairingKey"] = state.PairingKey
		inputs["partnerAsn"] = state.PartnerAsn
		inputs["privateInterconnectInfo"] = state.PrivateInterconnectInfo
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["router"] = state.Router
		inputs["selfLink"] = state.SelfLink
		inputs["state"] = state.State
		inputs["type"] = state.Type
		inputs["vlanTag8021q"] = state.VlanTag8021q
	}
	s, err := ctx.ReadResource("gcp:compute/interconnectAttachment:InterconnectAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InterconnectAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *InterconnectAttachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *InterconnectAttachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *InterconnectAttachment) CandidateSubnets() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["candidateSubnets"])
}

func (r *InterconnectAttachment) CloudRouterIpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["cloudRouterIpAddress"])
}

func (r *InterconnectAttachment) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *InterconnectAttachment) CustomerRouterIpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["customerRouterIpAddress"])
}

func (r *InterconnectAttachment) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *InterconnectAttachment) EdgeAvailabilityDomain() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["edgeAvailabilityDomain"])
}

func (r *InterconnectAttachment) GoogleReferenceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["googleReferenceId"])
}

func (r *InterconnectAttachment) Interconnect() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["interconnect"])
}

func (r *InterconnectAttachment) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *InterconnectAttachment) PairingKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["pairingKey"])
}

func (r *InterconnectAttachment) PartnerAsn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["partnerAsn"])
}

func (r *InterconnectAttachment) PrivateInterconnectInfo() *pulumi.Output {
	return r.s.State["privateInterconnectInfo"]
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *InterconnectAttachment) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *InterconnectAttachment) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

func (r *InterconnectAttachment) Router() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["router"])
}

// The URI of the created resource.
func (r *InterconnectAttachment) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *InterconnectAttachment) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

func (r *InterconnectAttachment) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

func (r *InterconnectAttachment) VlanTag8021q() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["vlanTag8021q"])
}

// Input properties used for looking up and filtering InterconnectAttachment resources.
type InterconnectAttachmentState struct {
	CandidateSubnets interface{}
	CloudRouterIpAddress interface{}
	CreationTimestamp interface{}
	CustomerRouterIpAddress interface{}
	Description interface{}
	EdgeAvailabilityDomain interface{}
	GoogleReferenceId interface{}
	Interconnect interface{}
	Name interface{}
	PairingKey interface{}
	PartnerAsn interface{}
	PrivateInterconnectInfo interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	Router interface{}
	// The URI of the created resource.
	SelfLink interface{}
	State interface{}
	Type interface{}
	VlanTag8021q interface{}
}

// The set of arguments for constructing a InterconnectAttachment resource.
type InterconnectAttachmentArgs struct {
	CandidateSubnets interface{}
	Description interface{}
	EdgeAvailabilityDomain interface{}
	Interconnect interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	Router interface{}
	Type interface{}
	VlanTag8021q interface{}
}
