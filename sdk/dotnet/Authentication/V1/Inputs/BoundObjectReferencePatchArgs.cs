// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Authentication.V1
{

    /// <summary>
    /// BoundObjectReference is a reference to an object that a token is bound to.
    /// </summary>
    public class BoundObjectReferencePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// API version of the referent.
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Name of the referent.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// UID of the referent.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        public BoundObjectReferencePatchArgs()
        {
        }
    }
}
