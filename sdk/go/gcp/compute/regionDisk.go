// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Persistent disks are durable storage devices that function similarly to
// the physical disks in a desktop or a server. Compute Engine manages the
// hardware behind these devices to ensure data redundancy and optimize
// performance for you. Persistent disks are available as either standard
// hard disk drives (HDD) or solid-state drives (SSD).
// 
// Persistent disks are located independently from your virtual machine
// instances, so you can detach or move persistent disks to keep your data
// even after you delete your instances. Persistent disk performance scales
// automatically with size, so you can resize your existing persistent disks
// or add more persistent disks to an instance to meet your performance and
// storage space requirements.
// 
// Add a persistent disk to your instance when you need reliable and
// affordable storage with consistent performance characteristics.
// 
// 
// To get more information about RegionDisk, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/regionDisks)
// * How-to Guides
//     * [Adding or Resizing Regional Persistent Disks](https://cloud.google.com/compute/docs/disks/regional-persistent-disk)
// 
// > **Warning:** All arguments including the disk encryption key will be stored in the raw
// state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=region_disk_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type RegionDisk struct {
	s *pulumi.ResourceState
}

// NewRegionDisk registers a new resource with the given unique name, arguments, and options.
func NewRegionDisk(ctx *pulumi.Context,
	name string, args *RegionDiskArgs, opts ...pulumi.ResourceOpt) (*RegionDisk, error) {
	if args == nil || args.ReplicaZones == nil {
		return nil, errors.New("missing required argument 'ReplicaZones'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["diskEncryptionKey"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["physicalBlockSizeBytes"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["replicaZones"] = nil
		inputs["size"] = nil
		inputs["snapshot"] = nil
		inputs["sourceSnapshotEncryptionKey"] = nil
		inputs["type"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["diskEncryptionKey"] = args.DiskEncryptionKey
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["physicalBlockSizeBytes"] = args.PhysicalBlockSizeBytes
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["replicaZones"] = args.ReplicaZones
		inputs["size"] = args.Size
		inputs["snapshot"] = args.Snapshot
		inputs["sourceSnapshotEncryptionKey"] = args.SourceSnapshotEncryptionKey
		inputs["type"] = args.Type
	}
	inputs["creationTimestamp"] = nil
	inputs["labelFingerprint"] = nil
	inputs["lastAttachTimestamp"] = nil
	inputs["lastDetachTimestamp"] = nil
	inputs["selfLink"] = nil
	inputs["sourceSnapshotId"] = nil
	inputs["users"] = nil
	s, err := ctx.RegisterResource("gcp:compute/regionDisk:RegionDisk", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionDisk{s: s}, nil
}

// GetRegionDisk gets an existing RegionDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDisk(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegionDiskState, opts ...pulumi.ResourceOpt) (*RegionDisk, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["diskEncryptionKey"] = state.DiskEncryptionKey
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["lastAttachTimestamp"] = state.LastAttachTimestamp
		inputs["lastDetachTimestamp"] = state.LastDetachTimestamp
		inputs["name"] = state.Name
		inputs["physicalBlockSizeBytes"] = state.PhysicalBlockSizeBytes
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["replicaZones"] = state.ReplicaZones
		inputs["selfLink"] = state.SelfLink
		inputs["size"] = state.Size
		inputs["snapshot"] = state.Snapshot
		inputs["sourceSnapshotEncryptionKey"] = state.SourceSnapshotEncryptionKey
		inputs["sourceSnapshotId"] = state.SourceSnapshotId
		inputs["type"] = state.Type
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("gcp:compute/regionDisk:RegionDisk", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionDisk{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegionDisk) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegionDisk) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *RegionDisk) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *RegionDisk) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *RegionDisk) DiskEncryptionKey() *pulumi.Output {
	return r.s.State["diskEncryptionKey"]
}

func (r *RegionDisk) LabelFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

func (r *RegionDisk) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *RegionDisk) LastAttachTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastAttachTimestamp"])
}

func (r *RegionDisk) LastDetachTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastDetachTimestamp"])
}

func (r *RegionDisk) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *RegionDisk) PhysicalBlockSizeBytes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["physicalBlockSizeBytes"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *RegionDisk) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *RegionDisk) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

func (r *RegionDisk) ReplicaZones() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["replicaZones"])
}

// The URI of the created resource.
func (r *RegionDisk) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *RegionDisk) Size() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["size"])
}

func (r *RegionDisk) Snapshot() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["snapshot"])
}

func (r *RegionDisk) SourceSnapshotEncryptionKey() *pulumi.Output {
	return r.s.State["sourceSnapshotEncryptionKey"]
}

func (r *RegionDisk) SourceSnapshotId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceSnapshotId"])
}

func (r *RegionDisk) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

func (r *RegionDisk) Users() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering RegionDisk resources.
type RegionDiskState struct {
	CreationTimestamp interface{}
	Description interface{}
	DiskEncryptionKey interface{}
	LabelFingerprint interface{}
	Labels interface{}
	LastAttachTimestamp interface{}
	LastDetachTimestamp interface{}
	Name interface{}
	PhysicalBlockSizeBytes interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	ReplicaZones interface{}
	// The URI of the created resource.
	SelfLink interface{}
	Size interface{}
	Snapshot interface{}
	SourceSnapshotEncryptionKey interface{}
	SourceSnapshotId interface{}
	Type interface{}
	Users interface{}
}

// The set of arguments for constructing a RegionDisk resource.
type RegionDiskArgs struct {
	Description interface{}
	DiskEncryptionKey interface{}
	Labels interface{}
	Name interface{}
	PhysicalBlockSizeBytes interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	ReplicaZones interface{}
	Size interface{}
	Snapshot interface{}
	SourceSnapshotEncryptionKey interface{}
	Type interface{}
}
