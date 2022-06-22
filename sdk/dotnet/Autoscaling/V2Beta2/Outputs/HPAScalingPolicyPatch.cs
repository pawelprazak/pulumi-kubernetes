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
    /// HPAScalingPolicy is a single policy which must hold true for a specified past interval.
    /// </summary>
    [OutputType]
    public sealed class HPAScalingPolicyPatch
    {
        /// <summary>
        /// PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).
        /// </summary>
        public readonly int PeriodSeconds;
        /// <summary>
        /// Type is used to specify the scaling policy.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Value contains the amount of change which is permitted by the policy. It must be greater than zero
        /// </summary>
        public readonly int Value;

        [OutputConstructor]
        private HPAScalingPolicyPatch(
            int periodSeconds,

            string type,

            int value)
        {
            PeriodSeconds = periodSeconds;
            Type = type;
            Value = value;
        }
    }
}
