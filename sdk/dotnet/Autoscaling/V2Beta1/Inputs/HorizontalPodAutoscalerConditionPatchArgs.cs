// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta1
{

    /// <summary>
    /// HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
    /// </summary>
    public class HorizontalPodAutoscalerConditionPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// lastTransitionTime is the last time the condition transitioned from one status to another
        /// </summary>
        [Input("lastTransitionTime")]
        public Input<string>? LastTransitionTime { get; set; }

        /// <summary>
        /// message is a human-readable explanation containing details about the transition
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        /// <summary>
        /// reason is the reason for the condition's last transition.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// status is the status of the condition (True, False, Unknown)
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// type describes the current condition
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public HorizontalPodAutoscalerConditionPatchArgs()
        {
        }
    }
}
