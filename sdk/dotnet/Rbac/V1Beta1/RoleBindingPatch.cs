// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Rbac.V1Beta1
{
    /// <summary>
    /// RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 RoleBinding, and will no longer be served in v1.20.
    /// </summary>
    [KubernetesResourceType("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingPatch")]
    public partial class RoleBindingPatch : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        /// </summary>
        [Output("roleRef")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Rbac.V1Beta1.RoleRef> RoleRef { get; private set; } = null!;

        /// <summary>
        /// Subjects holds references to the objects the role applies to.
        /// </summary>
        [Output("subjects")]
        public Output<ImmutableArray<Pulumi.Kubernetes.Types.Outputs.Rbac.V1Beta1.Subject>> Subjects { get; private set; } = null!;


        /// <summary>
        /// Create a RoleBindingPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleBindingPatch(string name, Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.RoleBindingPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal RoleBindingPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private RoleBindingPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.RoleBindingPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.RoleBindingPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.RoleBindingPatchArgs();
            args.ApiVersion = "rbac.authorization.k8s.io/v1beta1";
            args.Kind = "RoleBinding";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:rbac.authorization.k8s.io/v1:RoleBinding"},
                    new Pulumi.Alias { Type = "kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBinding"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RoleBindingPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleBindingPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RoleBindingPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1
{

    public class RoleBindingPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object's metadata.
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
        /// </summary>
        [Input("roleRef")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.RoleRefArgs>? RoleRef { get; set; }

        [Input("subjects")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.SubjectArgs>? _subjects;

        /// <summary>
        /// Subjects holds references to the objects the role applies to.
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.SubjectArgs> Subjects
        {
            get => _subjects ?? (_subjects = new InputList<Pulumi.Kubernetes.Types.Inputs.Rbac.V1Beta1.SubjectArgs>());
            set => _subjects = value;
        }

        public RoleBindingPatchArgs()
        {
        }
    }
}
