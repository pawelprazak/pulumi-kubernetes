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
    /// NamespaceStatus is information about the current status of a Namespace.
    /// </summary>
    public class NamespaceStatusPatchArgs : Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NamespaceConditionPatchArgs>? _conditions;

        /// <summary>
        /// Represents the latest available observations of a namespace's current state.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NamespaceConditionPatchArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.NamespaceConditionPatchArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
        /// </summary>
        [Input("phase")]
        public Input<string>? Phase { get; set; }

        public NamespaceStatusPatchArgs()
        {
        }
    }
}
