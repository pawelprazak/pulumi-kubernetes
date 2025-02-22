// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2
{

    /// <summary>
    /// ContainerResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing a single container in each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
    /// </summary>
    public class ContainerResourceMetricStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container is the name of the container in the pods of the scaling target
        /// </summary>
        [Input("container", required: true)]
        public Input<string> Container { get; set; } = null!;

        /// <summary>
        /// current contains the current value for the given metric
        /// </summary>
        [Input("current", required: true)]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.MetricValueStatusArgs> Current { get; set; } = null!;

        /// <summary>
        /// Name is the name of the resource in question.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ContainerResourceMetricStatusArgs()
        {
        }
    }
}
