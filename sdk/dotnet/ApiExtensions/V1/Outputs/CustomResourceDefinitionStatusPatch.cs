// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1
{

    /// <summary>
    /// CustomResourceDefinitionStatus indicates the state of the CustomResourceDefinition
    /// </summary>
    [OutputType]
    public sealed class CustomResourceDefinitionStatusPatch
    {
        /// <summary>
        /// acceptedNames are the names that are actually being used to serve discovery. They may be different than the names in spec.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceDefinitionNamesPatch AcceptedNames;
        /// <summary>
        /// conditions indicate state for particular aspects of a CustomResourceDefinition
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceDefinitionConditionPatch> Conditions;
        /// <summary>
        /// storedVersions lists all versions of CustomResources that were ever persisted. Tracking these versions allows a migration path for stored versions in etcd. The field is mutable so a migration controller can finish a migration to another version (ensuring no old objects are left in storage), and then remove the rest of the versions from this list. Versions may not be removed from `spec.versions` while they exist in this list.
        /// </summary>
        public readonly ImmutableArray<string> StoredVersions;

        [OutputConstructor]
        private CustomResourceDefinitionStatusPatch(
            Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceDefinitionNamesPatch acceptedNames,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.ApiExtensions.V1.CustomResourceDefinitionConditionPatch> conditions,

            ImmutableArray<string> storedVersions)
        {
            AcceptedNames = acceptedNames;
            Conditions = conditions;
            StoredVersions = storedVersions;
        }
    }
}
