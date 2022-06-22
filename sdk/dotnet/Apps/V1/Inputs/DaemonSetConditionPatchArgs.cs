// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Apps.V1
{

    /// <summary>
    /// DaemonSetCondition describes the state of a DaemonSet at a certain point.
    /// </summary>
    public class DaemonSetConditionPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Last time the condition transitioned from one status to another.
        /// </summary>
        [Input("lastTransitionTime")]
        public Input<string>? LastTransitionTime { get; set; }

        /// <summary>
        /// A human readable message indicating details about the transition.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// The reason for the condition's last transition.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Status of the condition, one of True, False, Unknown.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Type of DaemonSet condition.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DaemonSetConditionPatchArgs()
        {
        }
    }
}
