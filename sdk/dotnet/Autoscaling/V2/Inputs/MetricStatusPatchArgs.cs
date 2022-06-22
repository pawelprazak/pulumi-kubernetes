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
    /// MetricStatus describes the last-read state of a single metric.
    /// </summary>
    public class MetricStatusPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// container resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing a single container in each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
        /// </summary>
        [Input("containerResource")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.ContainerResourceMetricStatusPatchArgs>? ContainerResource { get; set; }

        /// <summary>
        /// external refers to a global metric that is not associated with any Kubernetes object. It allows autoscaling based on information coming from components running outside of cluster (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
        /// </summary>
        [Input("external")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.ExternalMetricStatusPatchArgs>? External { get; set; }

        /// <summary>
        /// object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object).
        /// </summary>
        [Input("object")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.ObjectMetricStatusPatchArgs>? Object { get; set; }

        /// <summary>
        /// pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value.
        /// </summary>
        [Input("pods")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.PodsMetricStatusPatchArgs>? Pods { get; set; }

        /// <summary>
        /// resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
        /// </summary>
        [Input("resource")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2.ResourceMetricStatusPatchArgs>? Resource { get; set; }

        /// <summary>
        /// type is the type of metric source.  It will be one of "ContainerResource", "External", "Object", "Pods" or "Resource", each corresponds to a matching field in the object. Note: "ContainerResource" type is available on when the feature-gate HPAContainerMetrics is enabled
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MetricStatusPatchArgs()
        {
        }
    }
}
