// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1
{

    /// <summary>
    /// VolumeAttachmentStatus is the status of a VolumeAttachment request.
    /// </summary>
    public class VolumeAttachmentStatusPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        /// </summary>
        [Input("attachError")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeErrorPatchArgs>? AttachError { get; set; }

        /// <summary>
        /// Indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        /// </summary>
        [Input("attached")]
        public Input<bool>? Attached { get; set; }

        [Input("attachmentMetadata")]
        private InputMap<string>? _attachmentMetadata;

        /// <summary>
        /// Upon successful attach, this field is populated with any information returned by the attach operation that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
        /// </summary>
        public InputMap<string> AttachmentMetadata
        {
            get => _attachmentMetadata ?? (_attachmentMetadata = new InputMap<string>());
            set => _attachmentMetadata = value;
        }

        /// <summary>
        /// The last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
        /// </summary>
        [Input("detachError")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Storage.V1Alpha1.VolumeErrorPatchArgs>? DetachError { get; set; }

        public VolumeAttachmentStatusPatchArgs()
        {
        }
    }
}
