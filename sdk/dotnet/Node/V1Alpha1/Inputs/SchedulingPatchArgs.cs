// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Node.V1Alpha1
{

    /// <summary>
    /// Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
    /// </summary>
    public class SchedulingPatchArgs : Pulumi.ResourceArgs
    {
        [Input("nodeSelector")]
        private InputMap<string>? _nodeSelector;

        /// <summary>
        /// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
        /// </summary>
        public InputMap<string> NodeSelector
        {
            get => _nodeSelector ?? (_nodeSelector = new InputMap<string>());
            set => _nodeSelector = value;
        }

        [Input("tolerations")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationPatchArgs>? _tolerations;

        /// <summary>
        /// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationPatchArgs> Tolerations
        {
            get => _tolerations ?? (_tolerations = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationPatchArgs>());
            set => _tolerations = value;
        }

        public SchedulingPatchArgs()
        {
        }
    }
}
