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
    /// ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
    /// </summary>
    [OutputType]
    public sealed class ExternalMetricStatusPatch
    {
        /// <summary>
        /// current contains the current value for the given metric
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricValueStatusPatch Current;
        /// <summary>
        /// metric identifies the target metric by name and selector
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricIdentifierPatch Metric;

        [OutputConstructor]
        private ExternalMetricStatusPatch(
            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricValueStatusPatch current,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2.MetricIdentifierPatch metric)
        {
            Current = current;
            Metric = metric;
        }
    }
}
