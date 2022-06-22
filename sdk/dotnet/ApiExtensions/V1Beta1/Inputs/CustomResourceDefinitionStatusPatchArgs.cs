// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.ApiExtensions.V1Beta1
{

    /// <summary>
    /// CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
    /// </summary>
    public class CustomResourceDefinitionStatusPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// acceptedNames are the names that are actually being used to serve discovery. They may be different than the names in spec.
        /// </summary>
        [Input("acceptedNames")]
        public Input<Pulumi.Kubernetes.Types.Inputs.ApiExtensions.V1Beta1.CustomResourceDefinitionNamesPatchArgs>? AcceptedNames { get; set; }

        [Input("conditions")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.ApiExtensions.V1Beta1.CustomResourceDefinitionConditionPatchArgs>? _conditions;

        /// <summary>
        /// conditions indicate state for particular aspects of a CustomResourceDefinition
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.ApiExtensions.V1Beta1.CustomResourceDefinitionConditionPatchArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Pulumi.Kubernetes.Types.Inputs.ApiExtensions.V1Beta1.CustomResourceDefinitionConditionPatchArgs>());
            set => _conditions = value;
        }

        [Input("storedVersions")]
        private InputList<string>? _storedVersions;

        /// <summary>
        /// storedVersions lists all versions of CustomResources that were ever persisted. Tracking these versions allows a migration path for stored versions in etcd. The field is mutable so a migration controller can finish a migration to another version (ensuring no old objects are left in storage), and then remove the rest of the versions from this list. Versions may not be removed from `spec.versions` while they exist in this list.
        /// </summary>
        public InputList<string> StoredVersions
        {
            get => _storedVersions ?? (_storedVersions = new InputList<string>());
            set => _storedVersions = value;
        }

        public CustomResourceDefinitionStatusPatchArgs()
        {
        }
    }
}
