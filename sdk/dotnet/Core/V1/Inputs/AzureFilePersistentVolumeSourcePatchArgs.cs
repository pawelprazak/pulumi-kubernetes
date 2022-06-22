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
    /// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
    /// </summary>
    public class AzureFilePersistentVolumeSourcePatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// secretName is the name of secret that contains Azure Storage Account Name and Key
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        /// <summary>
        /// secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
        /// </summary>
        [Input("secretNamespace")]
        public Input<string>? SecretNamespace { get; set; }

        /// <summary>
        /// shareName is the azure Share Name
        /// </summary>
        [Input("shareName")]
        public Input<string>? ShareName { get; set; }

        public AzureFilePersistentVolumeSourcePatchArgs()
        {
        }
    }
}
