// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2
{

    /// <summary>
    /// HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
    /// </summary>
    public class HorizontalPodAutoscalerBehaviorPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).
        /// </summary>
        [Input("scaleDown")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HPAScalingRulesPatchArgs>? ScaleDown { get; set; }

        /// <summary>
        /// scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of:
        ///   * increase no more than 4 pods per 60 seconds
        ///   * double the number of pods per 60 seconds
        /// No stabilization is used.
        /// </summary>
        [Input("scaleUp")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Autoscaling.V2Beta2.HPAScalingRulesPatchArgs>? ScaleUp { get; set; }

        public HorizontalPodAutoscalerBehaviorPatchArgs()
        {
        }
    }
}
