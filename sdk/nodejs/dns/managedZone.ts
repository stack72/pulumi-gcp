// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a zone within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/zones/) and
 * [API](https://cloud.google.com/dns/api/v1/managedZones).
 */
export class ManagedZone extends pulumi.CustomResource {
    /**
     * Get an existing ManagedZone resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedZoneState): ManagedZone {
        return new ManagedZone(name, <any>state, { id });
    }

    /**
     * A textual description field. Defaults to 'Managed by Terraform'.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The fully qualified DNS name of this zone, e.g. `terraform.io.`.
     */
    public readonly dnsName: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the instance.
     */
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The list of nameservers that will be authoritative for this
     * domain. Use NS records to redirect from your DNS provider to these names,
     * thus making Google Cloud DNS authoritative for this zone.
     */
    public /*out*/ readonly nameServers: pulumi.Output<string[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;

    /**
     * Create a ManagedZone resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedZoneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedZoneArgs | ManagedZoneState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ManagedZoneState = argsOrState as ManagedZoneState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nameServers"] = state ? state.nameServers : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ManagedZoneArgs | undefined;
            if (!args || args.dnsName === undefined) {
                throw new Error("Missing required property 'dnsName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["dnsName"] = args ? args.dnsName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["nameServers"] = undefined /*out*/;
        }
        super("gcp:dns/managedZone:ManagedZone", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedZone resources.
 */
export interface ManagedZoneState {
    /**
     * A textual description field. Defaults to 'Managed by Terraform'.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The fully qualified DNS name of this zone, e.g. `terraform.io.`.
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the instance.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The list of nameservers that will be authoritative for this
     * domain. Use NS records to redirect from your DNS provider to these names,
     * thus making Google Cloud DNS authoritative for this zone.
     */
    readonly nameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedZone resource.
 */
export interface ManagedZoneArgs {
    /**
     * A textual description field. Defaults to 'Managed by Terraform'.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The fully qualified DNS name of this zone, e.g. `terraform.io.`.
     */
    readonly dnsName: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the instance.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
}
