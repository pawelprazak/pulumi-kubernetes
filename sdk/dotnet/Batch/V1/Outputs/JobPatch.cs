// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Batch.V1
{

    /// <summary>
    /// Job represents the configuration of a single job.
    /// 
    /// This resource waits until its status is ready before registering success
    /// for create/update, and populating output properties from the current state of the resource.
    /// The following conditions are used to determine whether the resource creation has
    /// succeeded or failed:
    /// 
    /// 1. The Job's '.status.startTime' is set, which indicates that the Job has started running.
    /// 2. The Job's '.status.conditions' has a status of type 'Complete', and a 'status' set
    ///    to 'True'.
    /// 3. The Job's '.status.conditions' do not have a status of type 'Failed', with a
    /// 	'status' set to 'True'. If this condition is set, we should fail the Job immediately.
    /// 
    /// If the Job has not reached a Ready state after 10 minutes, it will
    /// time out and mark the resource update as Failed. You can override the default timeout value
    /// by setting the 'customTimeouts' option on the resource.
    /// 
    /// By default, if a resource failed to become ready in a previous update, 
    /// Pulumi will continue to wait for readiness on the next update. If you would prefer
    /// to schedule a replacement for an unready resource on the next update, you can add the
    /// "pulumi.com/replaceUnready": "true" annotation to the resource definition.
    /// </summary>
    [OutputType]
    public sealed class JobPatch
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
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch Metadata;
        /// <summary>
        /// Specification of the desired behavior of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobSpecPatch Spec;
        /// <summary>
        /// Current status of a job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobStatusPatch Status;

        [OutputConstructor]
        private JobPatch(
            string apiVersion,

            string kind,

            Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMetaPatch metadata,

            Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobSpecPatch spec,

            Pulumi.Kubernetes.Types.Outputs.Batch.V1.JobStatusPatch status)
        {
            ApiVersion = apiVersion;
            Kind = kind;
            Metadata = metadata;
            Spec = spec;
            Status = status;
        }
    }
}
