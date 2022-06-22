// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta2
{

    /// <summary>
    /// DaemonSetSpec is the specification of a daemon set.
    /// </summary>
    public class DaemonSetSpecPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
        /// </summary>
        [Input("minReadySeconds")]
        public Input<int>? MinReadySeconds { get; set; }

        /// <summary>
        /// The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
        /// </summary>
        [Input("revisionHistoryLimit")]
        public Input<int>? RevisionHistoryLimit { get; set; }

        /// <summary>
        /// A label query over pods that are managed by the daemon set. Must match in order to be controlled. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
        /// </summary>
        [Input("selector")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.LabelSelectorPatchArgs>? Selector { get; set; }

        /// <summary>
        /// An object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template
        /// </summary>
        [Input("template")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecPatchArgs>? Template { get; set; }

        /// <summary>
        /// An update strategy to replace existing DaemonSet pods with new pods.
        /// </summary>
        [Input("updateStrategy")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Apps.V1Beta2.DaemonSetUpdateStrategyPatchArgs>? UpdateStrategy { get; set; }

        public DaemonSetSpecPatchArgs()
        {
        }
    }
}
