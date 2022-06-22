// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRole, and will no longer be served in v1.20.
 */
export class ClusterRolePatch extends pulumi.CustomResource {
    /**
     * Get an existing ClusterRolePatch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ClusterRolePatch {
        return new ClusterRolePatch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRolePatch';

    /**
     * Returns true if the given object is an instance of ClusterRolePatch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterRolePatch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterRolePatch.__pulumiType;
    }

    /**
     * AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
     */
    public readonly aggregationRule!: pulumi.Output<outputs.rbac.v1beta1.AggregationRule>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"rbac.authorization.k8s.io/v1beta1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"ClusterRole">;
    /**
     * Standard object's metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMeta>;
    /**
     * Rules holds all the PolicyRules for this ClusterRole
     */
    public readonly rules!: pulumi.Output<outputs.rbac.v1beta1.PolicyRule[]>;

    /**
     * Create a ClusterRolePatch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterRolePatchArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["aggregationRule"] = args ? args.aggregationRule : undefined;
            resourceInputs["apiVersion"] = "rbac.authorization.k8s.io/v1beta1";
            resourceInputs["kind"] = "ClusterRole";
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
        } else {
            resourceInputs["aggregationRule"] = undefined /*out*/;
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["rules"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "kubernetes:rbac.authorization.k8s.io/v1:ClusterRole" }, { type: "kubernetes:rbac.authorization.k8s.io/v1alpha1:ClusterRole" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ClusterRolePatch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ClusterRolePatch resource.
 */
export interface ClusterRolePatchArgs {
    /**
     * AggregationRule is an optional field that describes how to build the Rules for this ClusterRole. If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be stomped by the controller.
     */
    aggregationRule?: pulumi.Input<inputs.rbac.v1beta1.AggregationRule>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"rbac.authorization.k8s.io/v1beta1">;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"ClusterRole">;
    /**
     * Standard object's metadata.
     */
    metadata?: pulumi.Input<inputs.meta.v1.ObjectMeta>;
    /**
     * Rules holds all the PolicyRules for this ClusterRole
     */
    rules?: pulumi.Input<pulumi.Input<inputs.rbac.v1beta1.PolicyRule>[]>;
}
