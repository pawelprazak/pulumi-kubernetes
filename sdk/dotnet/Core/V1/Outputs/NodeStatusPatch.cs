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
    /// NodeStatus is information about the current status of a node.
    /// </summary>
    [OutputType]
    public sealed class NodeStatusPatch
    {
        /// <summary>
        /// List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See http://pr.k8s.io/79391 for an example.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeAddressPatch> Addresses;
        /// <summary>
        /// Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Allocatable;
        /// <summary>
        /// Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
        /// </summary>
        public readonly ImmutableDictionary<string, string> Capacity;
        /// <summary>
        /// Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeConditionPatch> Conditions;
        /// <summary>
        /// Status of the config assigned to the node via the dynamic Kubelet config feature.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeConfigStatusPatch Config;
        /// <summary>
        /// Endpoints of daemons running on the Node.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeDaemonEndpointsPatch DaemonEndpoints;
        /// <summary>
        /// List of container images on this node
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.ContainerImagePatch> Images;
        /// <summary>
        /// Set of ids/uuids to uniquely identify the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#info
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSystemInfoPatch NodeInfo;
        /// <summary>
        /// NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
        /// </summary>
        public readonly string Phase;
        /// <summary>
        /// List of volumes that are attached to the node.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.AttachedVolumePatch> VolumesAttached;
        /// <summary>
        /// List of attachable volumes in use (mounted) by the node.
        /// </summary>
        public readonly ImmutableArray<string> VolumesInUse;

        [OutputConstructor]
        private NodeStatusPatch(
            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeAddressPatch> addresses,

            ImmutableDictionary<string, string> allocatable,

            ImmutableDictionary<string, string> capacity,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeConditionPatch> conditions,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeConfigStatusPatch config,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeDaemonEndpointsPatch daemonEndpoints,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.ContainerImagePatch> images,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeSystemInfoPatch nodeInfo,

            string phase,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Core.V1.AttachedVolumePatch> volumesAttached,

            ImmutableArray<string> volumesInUse)
        {
            Addresses = addresses;
            Allocatable = allocatable;
            Capacity = capacity;
            Conditions = conditions;
            Config = config;
            DaemonEndpoints = daemonEndpoints;
            Images = images;
            NodeInfo = nodeInfo;
            Phase = phase;
            VolumesAttached = volumesAttached;
            VolumesInUse = volumesInUse;
        }
    }
}
