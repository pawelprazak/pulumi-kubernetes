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
    /// ServiceStatus represents the current status of a service.
    /// </summary>
    public class ServiceStatusPatchArgs : Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ConditionPatchArgs>? _conditions;

        /// <summary>
        /// Current service state
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ConditionPatchArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ConditionPatchArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// LoadBalancer contains the current status of the load-balancer, if one is present.
        /// </summary>
        [Input("loadBalancer")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.LoadBalancerStatusPatchArgs>? LoadBalancer { get; set; }

        public ServiceStatusPatchArgs()
        {
        }
    }
}
