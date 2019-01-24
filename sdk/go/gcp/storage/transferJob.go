// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TransferJob struct {
	s *pulumi.ResourceState
}

// NewTransferJob registers a new resource with the given unique name, arguments, and options.
func NewTransferJob(ctx *pulumi.Context,
	name string, args *TransferJobArgs, opts ...pulumi.ResourceOpt) (*TransferJob, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.Schedule == nil {
		return nil, errors.New("missing required argument 'Schedule'")
	}
	if args == nil || args.TransferSpec == nil {
		return nil, errors.New("missing required argument 'TransferSpec'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["project"] = nil
		inputs["schedule"] = nil
		inputs["status"] = nil
		inputs["transferSpec"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["project"] = args.Project
		inputs["schedule"] = args.Schedule
		inputs["status"] = args.Status
		inputs["transferSpec"] = args.TransferSpec
	}
	inputs["creationTime"] = nil
	inputs["deletionTime"] = nil
	inputs["lastModificationTime"] = nil
	inputs["name"] = nil
	s, err := ctx.RegisterResource("gcp:storage/transferJob:TransferJob", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TransferJob{s: s}, nil
}

// GetTransferJob gets an existing TransferJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferJob(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TransferJobState, opts ...pulumi.ResourceOpt) (*TransferJob, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTime"] = state.CreationTime
		inputs["deletionTime"] = state.DeletionTime
		inputs["description"] = state.Description
		inputs["lastModificationTime"] = state.LastModificationTime
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["schedule"] = state.Schedule
		inputs["status"] = state.Status
		inputs["transferSpec"] = state.TransferSpec
	}
	s, err := ctx.ReadResource("gcp:storage/transferJob:TransferJob", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TransferJob{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TransferJob) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TransferJob) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *TransferJob) CreationTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTime"])
}

func (r *TransferJob) DeletionTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["deletionTime"])
}

func (r *TransferJob) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *TransferJob) LastModificationTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastModificationTime"])
}

func (r *TransferJob) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *TransferJob) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *TransferJob) Schedule() *pulumi.Output {
	return r.s.State["schedule"]
}

func (r *TransferJob) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

func (r *TransferJob) TransferSpec() *pulumi.Output {
	return r.s.State["transferSpec"]
}

// Input properties used for looking up and filtering TransferJob resources.
type TransferJobState struct {
	CreationTime interface{}
	DeletionTime interface{}
	Description interface{}
	LastModificationTime interface{}
	Name interface{}
	Project interface{}
	Schedule interface{}
	Status interface{}
	TransferSpec interface{}
}

// The set of arguments for constructing a TransferJob resource.
type TransferJobArgs struct {
	Description interface{}
	Project interface{}
	Schedule interface{}
	Status interface{}
	TransferSpec interface{}
}