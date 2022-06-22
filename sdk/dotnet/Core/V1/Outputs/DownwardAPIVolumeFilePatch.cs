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
    /// DownwardAPIVolumeFile represents information to create the file containing the pod field
    /// </summary>
    [OutputType]
    public sealed class DownwardAPIVolumeFilePatch
    {
        /// <summary>
        /// Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectFieldSelectorPatch FieldRef;
        /// <summary>
        /// Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
        /// </summary>
        public readonly int Mode;
        /// <summary>
        /// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceFieldSelectorPatch ResourceFieldRef;

        [OutputConstructor]
        private DownwardAPIVolumeFilePatch(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectFieldSelectorPatch fieldRef,

            int mode,

            string path,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.ResourceFieldSelectorPatch resourceFieldRef)
        {
            FieldRef = fieldRef;
            Mode = mode;
            Path = path;
            ResourceFieldRef = resourceFieldRef;
        }
    }
}
