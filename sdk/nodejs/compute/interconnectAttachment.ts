// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents an InterconnectAttachment (VLAN attachment) resource. For more
 * information, see Creating VLAN Attachments.
 * 
 */
export class InterconnectAttachment extends pulumi.CustomResource {
    /**
     * Get an existing InterconnectAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InterconnectAttachmentState, opts?: pulumi.CustomResourceOptions): InterconnectAttachment {
        return new InterconnectAttachment(name, <any>state, { ...opts, id: id });
    }

    public readonly candidateSubnets: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly cloudRouterIpAddress: pulumi.Output<string>;
    public /*out*/ readonly creationTimestamp: pulumi.Output<string>;
    public /*out*/ readonly customerRouterIpAddress: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public /*out*/ readonly googleReferenceId: pulumi.Output<string>;
    public readonly interconnect: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public /*out*/ readonly privateInterconnectInfo: pulumi.Output<{ tag8021q: number }>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    public readonly router: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public readonly vlanTag8021q: pulumi.Output<number | undefined>;

    /**
     * Create a InterconnectAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InterconnectAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InterconnectAttachmentArgs | InterconnectAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InterconnectAttachmentState = argsOrState as InterconnectAttachmentState | undefined;
            inputs["candidateSubnets"] = state ? state.candidateSubnets : undefined;
            inputs["cloudRouterIpAddress"] = state ? state.cloudRouterIpAddress : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["customerRouterIpAddress"] = state ? state.customerRouterIpAddress : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["googleReferenceId"] = state ? state.googleReferenceId : undefined;
            inputs["interconnect"] = state ? state.interconnect : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["privateInterconnectInfo"] = state ? state.privateInterconnectInfo : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["vlanTag8021q"] = state ? state.vlanTag8021q : undefined;
        } else {
            const args = argsOrState as InterconnectAttachmentArgs | undefined;
            if (!args || args.interconnect === undefined) {
                throw new Error("Missing required property 'interconnect'");
            }
            if (!args || args.router === undefined) {
                throw new Error("Missing required property 'router'");
            }
            inputs["candidateSubnets"] = args ? args.candidateSubnets : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["interconnect"] = args ? args.interconnect : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["vlanTag8021q"] = args ? args.vlanTag8021q : undefined;
            inputs["cloudRouterIpAddress"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["customerRouterIpAddress"] = undefined /*out*/;
            inputs["googleReferenceId"] = undefined /*out*/;
            inputs["privateInterconnectInfo"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        super("gcp:compute/interconnectAttachment:InterconnectAttachment", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InterconnectAttachment resources.
 */
export interface InterconnectAttachmentState {
    readonly candidateSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    readonly cloudRouterIpAddress?: pulumi.Input<string>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly customerRouterIpAddress?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly googleReferenceId?: pulumi.Input<string>;
    readonly interconnect?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly privateInterconnectInfo?: pulumi.Input<{ tag8021q?: pulumi.Input<number> }>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly router?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly vlanTag8021q?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a InterconnectAttachment resource.
 */
export interface InterconnectAttachmentArgs {
    readonly candidateSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    readonly description?: pulumi.Input<string>;
    readonly interconnect: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly router: pulumi.Input<string>;
    readonly vlanTag8021q?: pulumi.Input<number>;
}
