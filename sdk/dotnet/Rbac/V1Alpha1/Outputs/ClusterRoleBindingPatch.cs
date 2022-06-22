// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Rbac.V1Alpha1
{

    /// <summary>
    /// ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 ClusterRoleBinding, and will no longer be served in v1.20.
    /// </summary>
    [OutputType]
    public sealed class ClusterRoleBindingPatch
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch Metadata;
        /// <summary>
        /// RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Rbac.V1Alpha1.RoleRefPatch RoleRef;
        /// <summary>
        /// Subjects holds references to the objects the role applies to.
        /// </summary>
        public readonly ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Rbac.V1Alpha1.SubjectPatch> Subjects;

        [OutputConstructor]
        private ClusterRoleBindingPatch(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch metadata,

            Pulumi.Kubernetes.Types.Outputs.Rbac.V1Alpha1.RoleRefPatch roleRef,

            ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Rbac.V1Alpha1.SubjectPatch> subjects)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            RoleRef = roleRef;
            Subjects = subjects;
        }
    }
}
