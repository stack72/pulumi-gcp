// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A Backend Service defines a group of virtual machines that will serve traffic for load balancing. For more information
// see [the official documentation](https://cloud.google.com/compute/docs/load-balancing/http/backend-service)
// and the [API](https://cloud.google.com/compute/docs/reference/latest/backendServices).
// 
// For internal load balancing, use a [google_compute_region_backend_service](https://www.terraform.io/docs/providers/google/r/compute_region_backend_service.html).
type BackendService struct {
	s *pulumi.ResourceState
}

// NewBackendService registers a new resource with the given unique name, arguments, and options.
func NewBackendService(ctx *pulumi.Context,
	name string, args *BackendServiceArgs, opts ...pulumi.ResourceOpt) (*BackendService, error) {
	if args == nil || args.HealthChecks == nil {
		return nil, errors.New("missing required argument 'HealthChecks'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["affinityCookieTtlSec"] = nil
		inputs["backends"] = nil
		inputs["cdnPolicy"] = nil
		inputs["connectionDrainingTimeoutSec"] = nil
		inputs["customRequestHeaders"] = nil
		inputs["description"] = nil
		inputs["enableCdn"] = nil
		inputs["healthChecks"] = nil
		inputs["iap"] = nil
		inputs["name"] = nil
		inputs["portName"] = nil
		inputs["project"] = nil
		inputs["protocol"] = nil
		inputs["securityPolicy"] = nil
		inputs["sessionAffinity"] = nil
		inputs["timeoutSec"] = nil
	} else {
		inputs["affinityCookieTtlSec"] = args.AffinityCookieTtlSec
		inputs["backends"] = args.Backends
		inputs["cdnPolicy"] = args.CdnPolicy
		inputs["connectionDrainingTimeoutSec"] = args.ConnectionDrainingTimeoutSec
		inputs["customRequestHeaders"] = args.CustomRequestHeaders
		inputs["description"] = args.Description
		inputs["enableCdn"] = args.EnableCdn
		inputs["healthChecks"] = args.HealthChecks
		inputs["iap"] = args.Iap
		inputs["name"] = args.Name
		inputs["portName"] = args.PortName
		inputs["project"] = args.Project
		inputs["protocol"] = args.Protocol
		inputs["securityPolicy"] = args.SecurityPolicy
		inputs["sessionAffinity"] = args.SessionAffinity
		inputs["timeoutSec"] = args.TimeoutSec
	}
	inputs["fingerprint"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/backendService:BackendService", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BackendService{s: s}, nil
}

// GetBackendService gets an existing BackendService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackendService(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BackendServiceState, opts ...pulumi.ResourceOpt) (*BackendService, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["affinityCookieTtlSec"] = state.AffinityCookieTtlSec
		inputs["backends"] = state.Backends
		inputs["cdnPolicy"] = state.CdnPolicy
		inputs["connectionDrainingTimeoutSec"] = state.ConnectionDrainingTimeoutSec
		inputs["customRequestHeaders"] = state.CustomRequestHeaders
		inputs["description"] = state.Description
		inputs["enableCdn"] = state.EnableCdn
		inputs["fingerprint"] = state.Fingerprint
		inputs["healthChecks"] = state.HealthChecks
		inputs["iap"] = state.Iap
		inputs["name"] = state.Name
		inputs["portName"] = state.PortName
		inputs["project"] = state.Project
		inputs["protocol"] = state.Protocol
		inputs["securityPolicy"] = state.SecurityPolicy
		inputs["selfLink"] = state.SelfLink
		inputs["sessionAffinity"] = state.SessionAffinity
		inputs["timeoutSec"] = state.TimeoutSec
	}
	s, err := ctx.ReadResource("gcp:compute/backendService:BackendService", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &BackendService{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *BackendService) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *BackendService) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *BackendService) AffinityCookieTtlSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["affinityCookieTtlSec"])
}

// The list of backends that serve this BackendService. Structure is documented below.
func (r *BackendService) Backends() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["backends"])
}

// Cloud CDN configuration for this BackendService. Structure is documented below.
func (r *BackendService) CdnPolicy() *pulumi.Output {
	return r.s.State["cdnPolicy"]
}

// Time for which instance will be drained (not accept new connections,
// but still work to finish started ones). Defaults to `300`.
func (r *BackendService) ConnectionDrainingTimeoutSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["connectionDrainingTimeoutSec"])
}

// Headers that the
// HTTP/S load balancer should add to proxied requests. See [guide](https://cloud.google.com/compute/docs/load-balancing/http/backend-service#user-defined-request-headers) for details.
// This property is in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
func (r *BackendService) CustomRequestHeaders() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["customRequestHeaders"])
}

// The textual description for the backend service.
func (r *BackendService) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Whether or not to enable the Cloud CDN on the backend service.
func (r *BackendService) EnableCdn() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableCdn"])
}

// The fingerprint of the backend service.
func (r *BackendService) Fingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fingerprint"])
}

// Specifies a list of HTTP/HTTPS health checks
// for checking the health of the backend service. Currently at most one health
// check can be specified, and a health check is required.
func (r *BackendService) HealthChecks() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["healthChecks"])
}

// Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
func (r *BackendService) Iap() *pulumi.Output {
	return r.s.State["iap"]
}

// The name of the backend service.
func (r *BackendService) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of a service that has been added to an
// instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
func (r *BackendService) PortName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["portName"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *BackendService) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The protocol for incoming requests. Defaults to
// `HTTP`.
func (r *BackendService) Protocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocol"])
}

// Name or URI of a
// [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
func (r *BackendService) SecurityPolicy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["securityPolicy"])
}

// The URI of the created resource.
func (r *BackendService) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// How to distribute load. Options are `NONE` (no
// affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
// `GENERATED_COOKIE` (distribute load using a generated session cookie).
func (r *BackendService) SessionAffinity() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sessionAffinity"])
}

// The number of secs to wait for a backend to respond
// to a request before considering the request failed. Defaults to `30`.
func (r *BackendService) TimeoutSec() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["timeoutSec"])
}

// Input properties used for looking up and filtering BackendService resources.
type BackendServiceState struct {
	AffinityCookieTtlSec interface{}
	// The list of backends that serve this BackendService. Structure is documented below.
	Backends interface{}
	// Cloud CDN configuration for this BackendService. Structure is documented below.
	CdnPolicy interface{}
	// Time for which instance will be drained (not accept new connections,
	// but still work to finish started ones). Defaults to `300`.
	ConnectionDrainingTimeoutSec interface{}
	// Headers that the
	// HTTP/S load balancer should add to proxied requests. See [guide](https://cloud.google.com/compute/docs/load-balancing/http/backend-service#user-defined-request-headers) for details.
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	CustomRequestHeaders interface{}
	// The textual description for the backend service.
	Description interface{}
	// Whether or not to enable the Cloud CDN on the backend service.
	EnableCdn interface{}
	// The fingerprint of the backend service.
	Fingerprint interface{}
	// Specifies a list of HTTP/HTTPS health checks
	// for checking the health of the backend service. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks interface{}
	// Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
	Iap interface{}
	// The name of the backend service.
	Name interface{}
	// The name of a service that has been added to an
	// instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
	PortName interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The protocol for incoming requests. Defaults to
	// `HTTP`.
	Protocol interface{}
	// Name or URI of a
	// [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
	SecurityPolicy interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// How to distribute load. Options are `NONE` (no
	// affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
	// `GENERATED_COOKIE` (distribute load using a generated session cookie).
	SessionAffinity interface{}
	// The number of secs to wait for a backend to respond
	// to a request before considering the request failed. Defaults to `30`.
	TimeoutSec interface{}
}

// The set of arguments for constructing a BackendService resource.
type BackendServiceArgs struct {
	AffinityCookieTtlSec interface{}
	// The list of backends that serve this BackendService. Structure is documented below.
	Backends interface{}
	// Cloud CDN configuration for this BackendService. Structure is documented below.
	CdnPolicy interface{}
	// Time for which instance will be drained (not accept new connections,
	// but still work to finish started ones). Defaults to `300`.
	ConnectionDrainingTimeoutSec interface{}
	// Headers that the
	// HTTP/S load balancer should add to proxied requests. See [guide](https://cloud.google.com/compute/docs/load-balancing/http/backend-service#user-defined-request-headers) for details.
	// This property is in beta, and should be used with the terraform-provider-google-beta provider.
	// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta fields.
	CustomRequestHeaders interface{}
	// The textual description for the backend service.
	Description interface{}
	// Whether or not to enable the Cloud CDN on the backend service.
	EnableCdn interface{}
	// Specifies a list of HTTP/HTTPS health checks
	// for checking the health of the backend service. Currently at most one health
	// check can be specified, and a health check is required.
	HealthChecks interface{}
	// Specification for the Identity-Aware proxy. Disabled if not specified. Structure is documented below.
	Iap interface{}
	// The name of the backend service.
	Name interface{}
	// The name of a service that has been added to an
	// instance group in this backend. See [related docs](https://cloud.google.com/compute/docs/instance-groups/#specifying_service_endpoints) for details. Defaults to http.
	PortName interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The protocol for incoming requests. Defaults to
	// `HTTP`.
	Protocol interface{}
	// Name or URI of a
	// [security policy](https://cloud.google.com/armor/docs/security-policy-concepts) to add to the backend service.
	SecurityPolicy interface{}
	// How to distribute load. Options are `NONE` (no
	// affinity), `CLIENT_IP` (hash of the source/dest addresses / ports), and
	// `GENERATED_COOKIE` (distribute load using a generated session cookie).
	SessionAffinity interface{}
	// The number of secs to wait for a backend to respond
	// to a request before considering the request failed. Defaults to `30`.
	TimeoutSec interface{}
}
