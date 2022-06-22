// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// NodeStatus is information about the current status of a node.
    /// </summary>
    public class NodeStatusPatchArgs : Pulumi.ResourceArgs
    {
        [Input("addresses")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeAddressPatchArgs>? _addresses;

        /// <summary>
        /// List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See http://pr.k8s.io/79391 for an example.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeAddressPatchArgs> Addresses
        {
            get => _addresses ?? (_addresses = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeAddressPatchArgs>());
            set => _addresses = value;
        }

        [Input("allocatable")]
        private InputMap<string>? _allocatable;

        /// <summary>
        /// Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
        /// </summary>
        public InputMap<string> Allocatable
        {
            get => _allocatable ?? (_allocatable = new InputMap<string>());
            set => _allocatable = value;
        }

        [Input("capacity")]
        private InputMap<string>? _capacity;

        /// <summary>
        /// Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
        /// </summary>
        public InputMap<string> Capacity
        {
            get => _capacity ?? (_capacity = new InputMap<string>());
            set => _capacity = value;
        }

        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeConditionPatchArgs>? _conditions;

        /// <summary>
        /// Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/concepts/nodes/node/#condition
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeConditionPatchArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeConditionPatchArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Status of the config assigned to the node via the dynamic Kubelet config feature.
        /// </summary>
        [Input("config")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeConfigStatusPatchArgs>? Config { get; set; }

        /// <summary>
        /// Endpoints of daemons running on the Node.
        /// </summary>
        [Input("daemonEndpoints")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeDaemonEndpointsPatchArgs>? DaemonEndpoints { get; set; }

        [Input("images")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerImagePatchArgs>? _images;

        /// <summary>
        /// List of container images on this node
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerImagePatchArgs> Images
        {
            get => _images ?? (_images = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.ContainerImagePatchArgs>());
            set => _images = value;
        }

        /// <summary>
        /// Set of ids/uuids to uniquely identify the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#info
        /// </summary>
        [Input("nodeInfo")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSystemInfoPatchArgs>? NodeInfo { get; set; }

        /// <summary>
        /// NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
        /// </summary>
        [Input("phase")]
        public Input<string>? Phase { get; set; }

        [Input("volumesAttached")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.AttachedVolumePatchArgs>? _volumesAttached;

        /// <summary>
        /// List of volumes that are attached to the node.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.AttachedVolumePatchArgs> VolumesAttached
        {
            get => _volumesAttached ?? (_volumesAttached = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.AttachedVolumePatchArgs>());
            set => _volumesAttached = value;
        }

        [Input("volumesInUse")]
        private InputList<string>? _volumesInUse;

        /// <summary>
        /// List of attachable volumes in use (mounted) by the node.
        /// </summary>
        public InputList<string> VolumesInUse
        {
            get => _volumesInUse ?? (_volumesInUse = new InputList<string>());
            set => _volumesInUse = value;
        }

        public NodeStatusPatchArgs()
        {
        }
    }
}
