// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1
{

    /// <summary>
    /// Subject matches the originator of a request, as identified by the request authentication system. There are three ways of matching an originator; by user, group, or service account.
    /// </summary>
    [OutputType]
    public sealed class SubjectPatch
    {
        /// <summary>
        /// `group` matches based on user group name.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.GroupSubjectPatch Group;
        /// <summary>
        /// `kind` indicates which one of the other fields is non-empty. Required
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// `serviceAccount` matches ServiceAccounts.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.ServiceAccountSubjectPatch ServiceAccount;
        /// <summary>
        /// `user` matches based on username.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.UserSubjectPatch User;

        [OutputConstructor]
        private SubjectPatch(
            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.GroupSubjectPatch group,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.ServiceAccountSubjectPatch serviceAccount,

            Pulumi.Kubernetes.Types.Outputs.FlowControl.V1Beta1.UserSubjectPatch user)
        {
            Group = group;
            Kind = kind;
            ServiceAccount = serviceAccount;
            User = user;
        }
    }
}
