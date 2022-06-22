// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Certificates.V1Beta1
{

    /// <summary>
    /// This information is immutable after the request is created. Only the Request and Usages fields can be set on creation, other fields are derived by Kubernetes and cannot be modified by users.
    /// </summary>
    public class CertificateSigningRequestSpecPatchArgs : Pulumi.ResourceArgs
    {
        [Input("extra")]
        private InputMap<ImmutableArray<string>>? _extra;

        /// <summary>
        /// Extra information about the requesting user. See user.Info interface for details.
        /// </summary>
        public InputMap<ImmutableArray<string>> Extra
        {
            get => _extra ?? (_extra = new InputMap<ImmutableArray<string>>());
            set => _extra = value;
        }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Group information about the requesting user. See user.Info interface for details.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// Base64-encoded PKCS#10 CSR data
        /// </summary>
        [Input("request")]
        public Input<string>? Request { get; set; }

        /// <summary>
        /// Requested signer for the request. It is a qualified name in the form: `scope-hostname.io/name`. If empty, it will be defaulted:
        ///  1. If it's a kubelet client certificate, it is assigned
        ///     "kubernetes.io/kube-apiserver-client-kubelet".
        ///  2. If it's a kubelet serving certificate, it is assigned
        ///     "kubernetes.io/kubelet-serving".
        ///  3. Otherwise, it is assigned "kubernetes.io/legacy-unknown".
        /// Distribution of trust for signers happens out of band. You can select on this field using `spec.signerName`.
        /// </summary>
        [Input("signerName")]
        public Input<string>? SignerName { get; set; }

        /// <summary>
        /// UID information about the requesting user. See user.Info interface for details.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        [Input("usages")]
        private InputList<string>? _usages;

        /// <summary>
        /// allowedUsages specifies a set of usage contexts the key will be valid for. See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
        ///      https://tools.ietf.org/html/rfc5280#section-4.2.1.12
        /// </summary>
        public InputList<string> Usages
        {
            get => _usages ?? (_usages = new InputList<string>());
            set => _usages = value;
        }

        /// <summary>
        /// Information about the requesting user. See user.Info interface for details.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public CertificateSigningRequestSpecPatchArgs()
        {
        }
    }
}
