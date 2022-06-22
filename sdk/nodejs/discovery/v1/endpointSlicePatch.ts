// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * EndpointSlice represents a subset of the endpoints that implement a service. For a given service there may be multiple EndpointSlice objects, selected by labels, which must be joined to produce the full set of endpoints.
 */
export class EndpointSlicePatch extends pulumi.CustomResource {
    /**
     * Get an existing EndpointSlicePatch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EndpointSlicePatch {
        return new EndpointSlicePatch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'kubernetes:discovery.k8s.io/v1:EndpointSlicePatch';

    /**
     * Returns true if the given object is an instance of EndpointSlicePatch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointSlicePatch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointSlicePatch.__pulumiType;
    }

    /**
     * addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
     */
    public readonly addressType!: pulumi.Output<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    public readonly apiVersion!: pulumi.Output<"discovery.k8s.io/v1">;
    /**
     * endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
     */
    public readonly endpoints!: pulumi.Output<outputs.discovery.v1.EndpointPatch[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    public readonly kind!: pulumi.Output<"EndpointSlice">;
    /**
     * Standard object's metadata.
     */
    public readonly metadata!: pulumi.Output<outputs.meta.v1.ObjectMetaPatch>;
    /**
     * ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
     */
    public readonly ports!: pulumi.Output<outputs.discovery.v1.EndpointPortPatch[]>;

    /**
     * Create a EndpointSlicePatch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointSlicePatchArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["addressType"] = args ? args.addressType : undefined;
            resourceInputs["apiVersion"] = "discovery.k8s.io/v1";
            resourceInputs["endpoints"] = args ? args.endpoints : undefined;
            resourceInputs["kind"] = "EndpointSlice";
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
        } else {
            resourceInputs["addressType"] = undefined /*out*/;
            resourceInputs["apiVersion"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["ports"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "kubernetes:discovery.k8s.io/v1beta1:EndpointSlicePatch" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(EndpointSlicePatch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EndpointSlicePatch resource.
 */
export interface EndpointSlicePatchArgs {
    /**
     * addressType specifies the type of address carried by this EndpointSlice. All addresses in this slice must be the same type. This field is immutable after creation. The following address types are currently supported: * IPv4: Represents an IPv4 Address. * IPv6: Represents an IPv6 Address. * FQDN: Represents a Fully Qualified Domain Name.
     */
    addressType?: pulumi.Input<string>;
    /**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: pulumi.Input<"discovery.k8s.io/v1">;
    /**
     * endpoints is a list of unique endpoints in this slice. Each slice may include a maximum of 1000 endpoints.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.discovery.v1.EndpointPatch>[]>;
    /**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: pulumi.Input<"EndpointSlice">;
    /**
     * Standard object's metadata.
     */
    metadata?: pulumi.Input<inputs.meta.v1.ObjectMetaPatch>;
    /**
     * ports specifies the list of network ports exposed by each endpoint in this slice. Each port must have a unique name. When ports is empty, it indicates that there are no defined ports. When a port is defined with a nil port value, it indicates "all ports". Each slice may include a maximum of 100 ports.
     */
    ports?: pulumi.Input<pulumi.Input<inputs.discovery.v1.EndpointPortPatch>[]>;
}
