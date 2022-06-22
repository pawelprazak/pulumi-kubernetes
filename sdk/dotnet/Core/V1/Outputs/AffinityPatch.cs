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
    /// Affinity is a group of affinity scheduling rules.
    /// </summary>
    [OutputType]
    public sealed class AffinityPatch
    {
        /// <summary>
        /// Describes node affinity scheduling rules for the pod.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeAffinityPatch NodeAffinity;
        /// <summary>
        /// Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PodAffinityPatch PodAffinity;
        /// <summary>
        /// Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.PodAntiAffinityPatch PodAntiAffinity;

        [OutputConstructor]
        private AffinityPatch(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.NodeAffinityPatch nodeAffinity,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.PodAffinityPatch podAffinity,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.PodAntiAffinityPatch podAntiAffinity)
        {
            NodeAffinity = nodeAffinity;
            PodAffinity = podAffinity;
            PodAntiAffinity = podAntiAffinity;
        }
    }
}
