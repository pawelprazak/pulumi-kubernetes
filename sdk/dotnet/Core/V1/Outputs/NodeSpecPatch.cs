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
    /// NodeSpec describes the attributes that a node is created with.
    /// </summary>
    [OutputType]
    public sealed class NodeSpecPatch
    {
        /// <summary>
        /// Deprecated: Previously used to specify the source of the node's configuration for the DynamicKubeletConfig feature. This feature is removed from Kubelets as of 1.24 and will be fully removed in 1.26.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeConfigSourcePatch ConfigSource;
        /// <summary>
        /// Deprecated. Not all kubelets will set this field. Remove field after 1.13. see: https://issues.k8s.io/61966
        /// </summary>
        public readonly string ExternalID;
        /// <summary>
        /// PodCIDR represents the pod IP range assigned to the node.
        /// </summary>
        public readonly string PodCIDR;
        /// <summary>
        /// podCIDRs represents the IP ranges assigned to the node for usage by Pods on that node. If this field is specified, the 0th entry must match the podCIDR field. It may contain at most 1 value for each of IPv4 and IPv6.
        /// </summary>
        public readonly ImmutableArray<string> PodCIDRs;
        /// <summary>
        /// ID of the node assigned by the cloud provider in the format: &lt;ProviderName&gt;://&lt;ProviderSpecificNodeID&gt;
        /// </summary>
        public readonly string ProviderID;
        /// <summary>
        /// If specified, the node's taints.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.TaintPatch> Taints;
        /// <summary>
        /// Unschedulable controls node schedulability of new pods. By default, node is schedulable. More info: https://kubernetes.io/docs/concepts/nodes/node/#manual-node-administration
        /// </summary>
        public readonly bool Unschedulable;

        [OutputConstructor]
        private NodeSpecPatch(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeConfigSourcePatch configSource,

            string externalID,

            string podCIDR,

            ImmutableArray<string> podCIDRs,

            string providerID,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.TaintPatch> taints,

            bool unschedulable)
        {
            ConfigSource = configSource;
            ExternalID = externalID;
            PodCIDR = podCIDR;
            PodCIDRs = podCIDRs;
            ProviderID = providerID;
            Taints = taints;
            Unschedulable = unschedulable;
        }
    }
}
