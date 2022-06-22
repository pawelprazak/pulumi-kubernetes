// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Discovery.V1
{

    /// <summary>
    /// EndpointHints provides hints describing how an endpoint should be consumed.
    /// </summary>
    [OutputType]
    public sealed class EndpointHintsPatch
    {
        /// <summary>
        /// forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Discovery.V1.ForZonePatch> ForZones;

        [OutputConstructor]
        private EndpointHintsPatch(ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Discovery.V1.ForZonePatch> forZones)
        {
            ForZones = forZones;
        }
    }
}
