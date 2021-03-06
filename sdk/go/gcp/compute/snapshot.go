// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a Persistent Disk Snapshot resource.
// 
// Use snapshots to back up data from your persistent disks. Snapshots are
// different from public images and custom images, which are used primarily
// to create instances or configure instance templates. Snapshots are useful
// for periodic backup of the data on your persistent disks. You can create
// snapshots from persistent disks even while they are attached to running
// instances.
// 
// Snapshots are incremental, so you can create regular snapshots on a
// persistent disk faster and at a much lower cost than if you regularly
// created a full image of the disk.
// 
// 
// To get more information about Snapshot, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/snapshots)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=snapshot_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type Snapshot struct {
	s *pulumi.ResourceState
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOpt) (*Snapshot, error) {
	if args == nil || args.SourceDisk == nil {
		return nil, errors.New("missing required argument 'SourceDisk'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["snapshotEncryptionKey"] = nil
		inputs["sourceDisk"] = nil
		inputs["sourceDiskEncryptionKey"] = nil
		inputs["zone"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["snapshotEncryptionKey"] = args.SnapshotEncryptionKey
		inputs["sourceDisk"] = args.SourceDisk
		inputs["sourceDiskEncryptionKey"] = args.SourceDiskEncryptionKey
		inputs["zone"] = args.Zone
	}
	inputs["creationTimestamp"] = nil
	inputs["diskSizeGb"] = nil
	inputs["labelFingerprint"] = nil
	inputs["licenses"] = nil
	inputs["selfLink"] = nil
	inputs["snapshotId"] = nil
	inputs["sourceDiskLink"] = nil
	inputs["storageBytes"] = nil
	s, err := ctx.RegisterResource("gcp:compute/snapshot:Snapshot", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Snapshot{s: s}, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SnapshotState, opts ...pulumi.ResourceOpt) (*Snapshot, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["diskSizeGb"] = state.DiskSizeGb
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["licenses"] = state.Licenses
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["snapshotEncryptionKey"] = state.SnapshotEncryptionKey
		inputs["snapshotId"] = state.SnapshotId
		inputs["sourceDisk"] = state.SourceDisk
		inputs["sourceDiskEncryptionKey"] = state.SourceDiskEncryptionKey
		inputs["sourceDiskLink"] = state.SourceDiskLink
		inputs["storageBytes"] = state.StorageBytes
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/snapshot:Snapshot", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Snapshot{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Snapshot) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Snapshot) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Snapshot) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *Snapshot) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Snapshot) DiskSizeGb() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["diskSizeGb"])
}

func (r *Snapshot) LabelFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

func (r *Snapshot) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *Snapshot) Licenses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["licenses"])
}

func (r *Snapshot) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Snapshot) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *Snapshot) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *Snapshot) SnapshotEncryptionKey() *pulumi.Output {
	return r.s.State["snapshotEncryptionKey"]
}

func (r *Snapshot) SnapshotId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["snapshotId"])
}

func (r *Snapshot) SourceDisk() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceDisk"])
}

func (r *Snapshot) SourceDiskEncryptionKey() *pulumi.Output {
	return r.s.State["sourceDiskEncryptionKey"]
}

func (r *Snapshot) SourceDiskLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceDiskLink"])
}

func (r *Snapshot) StorageBytes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["storageBytes"])
}

func (r *Snapshot) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering Snapshot resources.
type SnapshotState struct {
	CreationTimestamp interface{}
	Description interface{}
	DiskSizeGb interface{}
	LabelFingerprint interface{}
	Labels interface{}
	Licenses interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	SnapshotEncryptionKey interface{}
	SnapshotId interface{}
	SourceDisk interface{}
	SourceDiskEncryptionKey interface{}
	SourceDiskLink interface{}
	StorageBytes interface{}
	Zone interface{}
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	Description interface{}
	Labels interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	SnapshotEncryptionKey interface{}
	SourceDisk interface{}
	SourceDiskEncryptionKey interface{}
	Zone interface{}
}
