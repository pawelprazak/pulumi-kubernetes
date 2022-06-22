// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2
{

    /// <summary>
    /// MetricValueStatus holds the current value for a metric
    /// </summary>
    [OutputType]
    public sealed class MetricValueStatusPatch
    {
        /// <summary>
        /// currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
        /// </summary>
        public readonly int AverageUtilization;
        /// <summary>
        /// averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
        /// </summary>
        public readonly string AverageValue;
        /// <summary>
        /// value is the current value of the metric (as a quantity).
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private MetricValueStatusPatch(
            int averageUtilization,

            string averageValue,

            string value)
        {
            AverageUtilization = averageUtilization;
            AverageValue = averageValue;
            Value = value;
        }
    }
}
