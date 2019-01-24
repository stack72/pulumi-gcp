// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_endpoints_service_grpc_service = new gcp.endpoints.Service("grpc_service", {
 *     grpcConfig: fs.readFileSync("service_spec.yml", "utf-8"),
 *     project: "project-id",
 *     protocOutput: fs.readFileSync("compiled_descriptor_file.pb", "utf-8"),
 *     serviceName: "api-name.endpoints.project-id.cloud.goog",
 * });
 * const google_endpoints_service_openapi_service = new gcp.endpoints.Service("openapi_service", {
 *     openapiConfig: fs.readFileSync("openapi_spec.yml", "utf-8"),
 *     project: "project-id",
 *     serviceName: "api-name.endpoints.project-id.cloud.goog",
 * });
 * ```
 * The example in `examples/endpoints_on_compute_engine` shows the API from the quickstart running on a Compute Engine VM and reachable through Cloud Endpoints, which may also be useful.
 * 
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly apis: pulumi.Output<{ methods: { name: string, requestType: string, responseType: string, syntax: string }[], name: string, syntax: string, version: string }[]>;
    public /*out*/ readonly configId: pulumi.Output<string>;
    public /*out*/ readonly dnsAddress: pulumi.Output<string>;
    public /*out*/ readonly endpoints: pulumi.Output<{ address: string, name: string }[]>;
    public readonly grpcConfig: pulumi.Output<string | undefined>;
    public readonly openapiConfig: pulumi.Output<string | undefined>;
    public readonly project: pulumi.Output<string>;
    public readonly protocOutputBase64: pulumi.Output<string | undefined>;
    public readonly serviceName: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ServiceState = argsOrState as ServiceState | undefined;
            inputs["apis"] = state ? state.apis : undefined;
            inputs["configId"] = state ? state.configId : undefined;
            inputs["dnsAddress"] = state ? state.dnsAddress : undefined;
            inputs["endpoints"] = state ? state.endpoints : undefined;
            inputs["grpcConfig"] = state ? state.grpcConfig : undefined;
            inputs["openapiConfig"] = state ? state.openapiConfig : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["protocOutputBase64"] = state ? state.protocOutputBase64 : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["grpcConfig"] = args ? args.grpcConfig : undefined;
            inputs["openapiConfig"] = args ? args.openapiConfig : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["protocOutputBase64"] = args ? args.protocOutputBase64 : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["apis"] = undefined /*out*/;
            inputs["configId"] = undefined /*out*/;
            inputs["dnsAddress"] = undefined /*out*/;
            inputs["endpoints"] = undefined /*out*/;
        }
        super("gcp:endpoints/service:Service", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    readonly apis?: pulumi.Input<pulumi.Input<{ methods?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string>, requestType?: pulumi.Input<string>, responseType?: pulumi.Input<string>, syntax?: pulumi.Input<string> }>[]>, name?: pulumi.Input<string>, syntax?: pulumi.Input<string>, version?: pulumi.Input<string> }>[]>;
    readonly configId?: pulumi.Input<string>;
    readonly dnsAddress?: pulumi.Input<string>;
    readonly endpoints?: pulumi.Input<pulumi.Input<{ address?: pulumi.Input<string>, name?: pulumi.Input<string> }>[]>;
    readonly grpcConfig?: pulumi.Input<string>;
    readonly openapiConfig?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly protocOutputBase64?: pulumi.Input<string>;
    readonly serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    readonly grpcConfig?: pulumi.Input<string>;
    readonly openapiConfig?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
    readonly protocOutputBase64?: pulumi.Input<string>;
    readonly serviceName: pulumi.Input<string>;
}
