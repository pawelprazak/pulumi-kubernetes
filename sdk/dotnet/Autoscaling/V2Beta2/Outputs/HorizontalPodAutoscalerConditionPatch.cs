// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2
{

    /// <summary>
    /// HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
    /// </summary>
    [OutputType]
    public sealed class HorizontalPodAutoscalerConditionPatch
    {
        /// <summary>
        /// lastTransitionTime is the last time the condition transitioned from one status to another
        /// </summary>
        public readonly string LastTransitionTime;
        /// <summary>
        /// message is a human-readable explanation containing details about the transition
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// reason is the reason for the condition's last transition.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// status is the status of the condition (True, False, Unknown)
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// type describes the current condition
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private HorizontalPodAutoscalerConditionPatch(
            string lastTransitionTime,

            string message,

            string reason,

            string status,

            string type)
        {
            LastTransitionTime = lastTransitionTime;
            Message = message;
            Reason = reason;
            Status = status;
            Type = type;
        }
    }
}
