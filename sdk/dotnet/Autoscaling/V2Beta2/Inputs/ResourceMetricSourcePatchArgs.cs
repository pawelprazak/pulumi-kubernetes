// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2
{

    /// <summary>
    /// ResourceMetricSource indicates how to scale on a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  The values will be averaged together before being compared to the target.  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.  Only one "target" type should be set.
    /// </summary>
    public class ResourceMetricSourcePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// name is the name of the resource in question.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// target specifies the target value for the given metric
        /// </summary>
        [Input("target")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.MetricTargetPatchArgs>? Target { get; set; }

        public ResourceMetricSourcePatchArgs()
        {
        }
    }
}
