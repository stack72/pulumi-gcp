// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RouterNat extends pulumi.CustomResource {
    /**
     * Get an existing RouterNat resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouterNatState, opts?: pulumi.CustomResourceOptions): RouterNat {
        return new RouterNat(name, <any>state, { ...opts, id: id });
    }

    public readonly icmpIdleTimeoutSec: pulumi.Output<number | undefined>;
    public readonly minPortsPerVm: pulumi.Output<number | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly natIpAllocateOption: pulumi.Output<string>;
    public readonly natIps: pulumi.Output<string[] | undefined>;
    public readonly project: pulumi.Output<string>;
    public readonly region: pulumi.Output<string>;
    public readonly router: pulumi.Output<string>;
    public readonly sourceSubnetworkIpRangesToNat: pulumi.Output<string | undefined>;
    public readonly subnetworks: pulumi.Output<{ name: string, secondaryIpRangeNames?: string[], sourceIpRangesToNats: string[] }[] | undefined>;
    public readonly tcpEstablishedIdleTimeoutSec: pulumi.Output<number | undefined>;
    public readonly tcpTransitoryIdleTimeoutSec: pulumi.Output<number | undefined>;
    public readonly udpIdleTimeoutSec: pulumi.Output<number | undefined>;

    /**
     * Create a RouterNat resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouterNatArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouterNatArgs | RouterNatState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RouterNatState = argsOrState as RouterNatState | undefined;
            inputs["icmpIdleTimeoutSec"] = state ? state.icmpIdleTimeoutSec : undefined;
            inputs["minPortsPerVm"] = state ? state.minPortsPerVm : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["natIpAllocateOption"] = state ? state.natIpAllocateOption : undefined;
            inputs["natIps"] = state ? state.natIps : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["router"] = state ? state.router : undefined;
            inputs["sourceSubnetworkIpRangesToNat"] = state ? state.sourceSubnetworkIpRangesToNat : undefined;
            inputs["subnetworks"] = state ? state.subnetworks : undefined;
            inputs["tcpEstablishedIdleTimeoutSec"] = state ? state.tcpEstablishedIdleTimeoutSec : undefined;
            inputs["tcpTransitoryIdleTimeoutSec"] = state ? state.tcpTransitoryIdleTimeoutSec : undefined;
            inputs["udpIdleTimeoutSec"] = state ? state.udpIdleTimeoutSec : undefined;
        } else {
            const args = argsOrState as RouterNatArgs | undefined;
            if (!args || args.natIpAllocateOption === undefined) {
                throw new Error("Missing required property 'natIpAllocateOption'");
            }
            if (!args || args.router === undefined) {
                throw new Error("Missing required property 'router'");
            }
            inputs["icmpIdleTimeoutSec"] = args ? args.icmpIdleTimeoutSec : undefined;
            inputs["minPortsPerVm"] = args ? args.minPortsPerVm : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["natIpAllocateOption"] = args ? args.natIpAllocateOption : undefined;
            inputs["natIps"] = args ? args.natIps : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["router"] = args ? args.router : undefined;
            inputs["sourceSubnetworkIpRangesToNat"] = args ? args.sourceSubnetworkIpRangesToNat : undefined;
            inputs["subnetworks"] = args ? args.subnetworks : undefined;
            inputs["tcpEstablishedIdleTimeoutSec"] = args ? args.tcpEstablishedIdleTimeoutSec : undefined;
            inputs["tcpTransitoryIdleTimeoutSec"] = args ? args.tcpTransitoryIdleTimeoutSec : undefined;
            inputs["udpIdleTimeoutSec"] = args ? args.udpIdleTimeoutSec : undefined;
        }
        super("gcp:compute/routerNat:RouterNat", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouterNat resources.
 */
export interface RouterNatState {
    readonly icmpIdleTimeoutSec?: pulumi.Input<number>;
    readonly minPortsPerVm?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly natIpAllocateOption?: pulumi.Input<string>;
    readonly natIps?: pulumi.Input<pulumi.Input<string>[]>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly router?: pulumi.Input<string>;
    readonly sourceSubnetworkIpRangesToNat?: pulumi.Input<string>;
    readonly subnetworks?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, secondaryIpRangeNames?: pulumi.Input<pulumi.Input<string>[]>, sourceIpRangesToNats: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    readonly tcpEstablishedIdleTimeoutSec?: pulumi.Input<number>;
    readonly tcpTransitoryIdleTimeoutSec?: pulumi.Input<number>;
    readonly udpIdleTimeoutSec?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RouterNat resource.
 */
export interface RouterNatArgs {
    readonly icmpIdleTimeoutSec?: pulumi.Input<number>;
    readonly minPortsPerVm?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly natIpAllocateOption: pulumi.Input<string>;
    readonly natIps?: pulumi.Input<pulumi.Input<string>[]>;
    readonly project?: pulumi.Input<string>;
    readonly region?: pulumi.Input<string>;
    readonly router: pulumi.Input<string>;
    readonly sourceSubnetworkIpRangesToNat?: pulumi.Input<string>;
    readonly subnetworks?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, secondaryIpRangeNames?: pulumi.Input<pulumi.Input<string>[]>, sourceIpRangesToNats: pulumi.Input<pulumi.Input<string>[]> }>[]>;
    readonly tcpEstablishedIdleTimeoutSec?: pulumi.Input<number>;
    readonly tcpTransitoryIdleTimeoutSec?: pulumi.Input<number>;
    readonly udpIdleTimeoutSec?: pulumi.Input<number>;
}
