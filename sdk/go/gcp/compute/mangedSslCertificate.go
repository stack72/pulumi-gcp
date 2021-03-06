// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// An SslCertificate resource, used for HTTPS load balancing.  This resource
// represents a certificate for which the certificate secrets are created and
// managed by Google.
// 
// For a resource where you provide the key, see the
// SSL Certificate resource.
// 
// > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
// 
// To get more information about ManagedSslCertificate, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/sslCertificates)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/ssl-certificates)
// 
// > **Warning:** This resource should be used with extreme caution!  Provisioning an SSL
// certificate is complex.  Ensure that you understand the lifecycle of a
// certificate before attempting complex tasks like cert rotation automatically.
// This resource will "return" as soon as the certificate object is created,
// but post-creation the certificate object will go through a "provisioning"
// process.  The provisioning process can complete only when the domain name
// for which the certificate is created points to a target pool which, itself,
// points at the certificate.  Depending on your DNS provider, this may take
// some time, and migrating from self-managed certificates to Google-managed
// certificates may entail some downtime while the certificate provisions.
// 
// In conclusion: Be extremely cautious.
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=managed_ssl_certificate_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type MangedSslCertificate struct {
	s *pulumi.ResourceState
}

// NewMangedSslCertificate registers a new resource with the given unique name, arguments, and options.
func NewMangedSslCertificate(ctx *pulumi.Context,
	name string, args *MangedSslCertificateArgs, opts ...pulumi.ResourceOpt) (*MangedSslCertificate, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["certificateId"] = nil
		inputs["description"] = nil
		inputs["managed"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["type"] = nil
	} else {
		inputs["certificateId"] = args.CertificateId
		inputs["description"] = args.Description
		inputs["managed"] = args.Managed
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["type"] = args.Type
	}
	inputs["creationTimestamp"] = nil
	inputs["expireTime"] = nil
	inputs["selfLink"] = nil
	inputs["subjectAlternativeNames"] = nil
	s, err := ctx.RegisterResource("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MangedSslCertificate{s: s}, nil
}

// GetMangedSslCertificate gets an existing MangedSslCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMangedSslCertificate(ctx *pulumi.Context,
	name string, id pulumi.ID, state *MangedSslCertificateState, opts ...pulumi.ResourceOpt) (*MangedSslCertificate, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["certificateId"] = state.CertificateId
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["expireTime"] = state.ExpireTime
		inputs["managed"] = state.Managed
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["subjectAlternativeNames"] = state.SubjectAlternativeNames
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("gcp:compute/mangedSslCertificate:MangedSslCertificate", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &MangedSslCertificate{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *MangedSslCertificate) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *MangedSslCertificate) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *MangedSslCertificate) CertificateId() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["certificateId"])
}

func (r *MangedSslCertificate) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *MangedSslCertificate) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *MangedSslCertificate) ExpireTime() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["expireTime"])
}

func (r *MangedSslCertificate) Managed() *pulumi.Output {
	return r.s.State["managed"]
}

func (r *MangedSslCertificate) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *MangedSslCertificate) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *MangedSslCertificate) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *MangedSslCertificate) SubjectAlternativeNames() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["subjectAlternativeNames"])
}

func (r *MangedSslCertificate) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering MangedSslCertificate resources.
type MangedSslCertificateState struct {
	CertificateId interface{}
	CreationTimestamp interface{}
	Description interface{}
	ExpireTime interface{}
	Managed interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	SubjectAlternativeNames interface{}
	Type interface{}
}

// The set of arguments for constructing a MangedSslCertificate resource.
type MangedSslCertificateArgs struct {
	CertificateId interface{}
	Description interface{}
	Managed interface{}
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Type interface{}
}
