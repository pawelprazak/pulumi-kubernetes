// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1
{

    /// <summary>
    /// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
    /// </summary>
    public class VolumeAttachmentSpecPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
        /// </summary>
        [Input("attacher")]
        public Input<string>? Attacher { get; set; }

        /// <summary>
        /// The node that the volume should be attached to.
        /// </summary>
        [Input("nodeName")]
        public Input<string>? NodeName { get; set; }

        /// <summary>
        /// Source represents the volume that should be attached.
        /// </summary>
        [Input("source")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Storage.V1.VolumeAttachmentSourcePatchArgs>? Source { get; set; }

        public VolumeAttachmentSpecPatchArgs()
        {
        }
    }
}
