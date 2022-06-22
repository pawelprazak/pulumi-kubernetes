// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.FlowControl.V1Beta1
{

    /// <summary>
    /// PriorityLevelConfigurationReference contains information that points to the "request-priority" being used.
    /// </summary>
    public class PriorityLevelConfigurationReferencePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// `name` is the name of the priority level configuration being referenced Required.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PriorityLevelConfigurationReferencePatchArgs()
        {
        }
    }
}
