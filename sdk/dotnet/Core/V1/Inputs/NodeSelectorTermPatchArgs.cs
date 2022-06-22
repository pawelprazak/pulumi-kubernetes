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
    /// A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
    /// </summary>
    public class NodeSelectorTermPatchArgs : Pulumi.ResourceArgs
    {
        [Input("matchExpressions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorRequirementPatchArgs>? _matchExpressions;

        /// <summary>
        /// A list of node selector requirements by node's labels.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorRequirementPatchArgs> MatchExpressions
        {
            get => _matchExpressions ?? (_matchExpressions = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorRequirementPatchArgs>());
            set => _matchExpressions = value;
        }

        [Input("matchFields")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorRequirementPatchArgs>? _matchFields;

        /// <summary>
        /// A list of node selector requirements by node's fields.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorRequirementPatchArgs> MatchFields
        {
            get => _matchFields ?? (_matchFields = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NodeSelectorRequirementPatchArgs>());
            set => _matchFields = value;
        }

        public NodeSelectorTermPatchArgs()
        {
        }
    }
}
