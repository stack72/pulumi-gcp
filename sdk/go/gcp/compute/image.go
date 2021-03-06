// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents an Image resource.
// 
// Google Compute Engine uses operating system images to create the root
// persistent disks for your instances. You specify an image when you create
// an instance. Images contain a boot loader, an operating system, and a
// root file system. Linux operating system images are also capable of
// running containers on Compute Engine.
// 
// Images can be either public or custom.
// 
// Public images are provided and maintained by Google, open-source
// communities, and third-party vendors. By default, all projects have
// access to these images and can use them to create instances.  Custom
// images are available only to your project. You can create a custom image
// from root persistent disks and other images. Then, use the custom image
// to create an instance.
// 
// 
// To get more information about Image, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/images)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/compute/docs/images)
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=image_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type Image struct {
	s *pulumi.ResourceState
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOpt) (*Image, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["diskSizeGb"] = nil
		inputs["family"] = nil
		inputs["labels"] = nil
		inputs["licenses"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["rawDisk"] = nil
		inputs["sourceDisk"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["diskSizeGb"] = args.DiskSizeGb
		inputs["family"] = args.Family
		inputs["labels"] = args.Labels
		inputs["licenses"] = args.Licenses
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["rawDisk"] = args.RawDisk
		inputs["sourceDisk"] = args.SourceDisk
	}
	inputs["archiveSizeBytes"] = nil
	inputs["creationTimestamp"] = nil
	inputs["labelFingerprint"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/image:Image", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Image{s: s}, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ImageState, opts ...pulumi.ResourceOpt) (*Image, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["archiveSizeBytes"] = state.ArchiveSizeBytes
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["diskSizeGb"] = state.DiskSizeGb
		inputs["family"] = state.Family
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["licenses"] = state.Licenses
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["rawDisk"] = state.RawDisk
		inputs["selfLink"] = state.SelfLink
		inputs["sourceDisk"] = state.SourceDisk
	}
	s, err := ctx.ReadResource("gcp:compute/image:Image", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Image{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Image) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Image) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Image) ArchiveSizeBytes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["archiveSizeBytes"])
}

func (r *Image) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *Image) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Image) DiskSizeGb() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["diskSizeGb"])
}

func (r *Image) Family() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["family"])
}

func (r *Image) LabelFingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

func (r *Image) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *Image) Licenses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["licenses"])
}

func (r *Image) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Image) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Image) RawDisk() *pulumi.Output {
	return r.s.State["rawDisk"]
}

// The URI of the created resource.
func (r *Image) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *Image) SourceDisk() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceDisk"])
}

// Input properties used for looking up and filtering Image resources.
type ImageState struct {
	ArchiveSizeBytes interface{}
	CreationTimestamp interface{}
	Description interface{}
	DiskSizeGb interface{}
	Family interface{}
	LabelFingerprint interface{}
	Labels interface{}
	Licenses interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	RawDisk interface{}
	// The URI of the created resource.
	SelfLink interface{}
	SourceDisk interface{}
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	Description interface{}
	DiskSizeGb interface{}
	Family interface{}
	Labels interface{}
	Licenses interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	RawDisk interface{}
	SourceDisk interface{}
}
