// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// ReplicationControllerCondition describes the state of a replication controller at a certain point.
    /// </summary>
    [OutputType]
    public sealed class ReplicationControllerConditionPatch
    {
        /// <summary>
        /// The last time the condition transitioned from one status to another.
        /// </summary>
        public readonly string LastTransitionTime;
        /// <summary>
        /// A human readable message indicating details about the transition.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The reason for the condition's last transition.
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// Status of the condition, one of True, False, Unknown.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Type of replication controller condition.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ReplicationControllerConditionPatch(
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
