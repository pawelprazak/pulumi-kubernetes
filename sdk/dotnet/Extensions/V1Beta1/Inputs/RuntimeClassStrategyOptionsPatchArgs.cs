// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Extensions.V1Beta1
{

    /// <summary>
    /// RuntimeClassStrategyOptions define the strategy that will dictate the allowable RuntimeClasses for a pod.
    /// </summary>
    public class RuntimeClassStrategyOptionsPatchArgs : Pulumi.ResourceArgs
    {
        [Input("allowedRuntimeClassNames")]
        private InputList<string>? _allowedRuntimeClassNames;

        /// <summary>
        /// allowedRuntimeClassNames is a whitelist of RuntimeClass names that may be specified on a pod. A value of "*" means that any RuntimeClass name is allowed, and must be the only item in the list. An empty list requires the RuntimeClassName field to be unset.
        /// </summary>
        public InputList<string> AllowedRuntimeClassNames
        {
            get => _allowedRuntimeClassNames ?? (_allowedRuntimeClassNames = new InputList<string>());
            set => _allowedRuntimeClassNames = value;
        }

        /// <summary>
        /// defaultRuntimeClassName is the default RuntimeClassName to set on the pod. The default MUST be allowed by the allowedRuntimeClassNames list. A value of nil does not mutate the Pod.
        /// </summary>
        [Input("defaultRuntimeClassName")]
        public Input<string>? DefaultRuntimeClassName { get; set; }

        public RuntimeClassStrategyOptionsPatchArgs()
        {
        }
    }
}
