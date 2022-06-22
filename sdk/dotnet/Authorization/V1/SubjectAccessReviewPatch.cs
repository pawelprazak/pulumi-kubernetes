// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Authorization.V1
{
    /// <summary>
    /// SubjectAccessReview checks whether or not a user or group can perform an action.
    /// </summary>
    [KubernetesResourceType("kubernetes:authorization.k8s.io/v1:SubjectAccessReviewPatch")]
    public partial class SubjectAccessReviewPatch : KubernetesResource
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
        /// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// Spec holds information about the request being evaluated
        /// </summary>
        [Output("spec")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Authorization.V1.SubjectAccessReviewSpec> Spec { get; private set; } = null!;

        /// <summary>
        /// Status is filled in by the server and indicates whether the request is allowed or not
        /// </summary>
        [Output("status")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Authorization.V1.SubjectAccessReviewStatus> Status { get; private set; } = null!;


        /// <summary>
        /// Create a SubjectAccessReviewPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubjectAccessReviewPatch(string name, Pulumi.Kubernetes.Types.Inputs.Authorization.V1.SubjectAccessReviewPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:authorization.k8s.io/v1:SubjectAccessReviewPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal SubjectAccessReviewPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:authorization.k8s.io/v1:SubjectAccessReviewPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private SubjectAccessReviewPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:authorization.k8s.io/v1:SubjectAccessReviewPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Authorization.V1.SubjectAccessReviewPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Authorization.V1.SubjectAccessReviewPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Authorization.V1.SubjectAccessReviewPatchArgs();
            args.ApiVersion = "authorization.k8s.io/v1";
            args.Kind = "SubjectAccessReview";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:authorization.k8s.io/v1beta1:SubjectAccessReview"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SubjectAccessReviewPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubjectAccessReviewPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SubjectAccessReviewPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Authorization.V1
{

    public class SubjectAccessReviewPatchArgs : Pulumi.ResourceArgs
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
        /// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// Spec holds information about the request being evaluated
        /// </summary>
        [Input("spec")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Authorization.V1.SubjectAccessReviewSpecArgs>? Spec { get; set; }

        public SubjectAccessReviewPatchArgs()
        {
        }
    }
}
