// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.AuditRegistraion.V1Alpha1
{

    /// <summary>
    /// WebhookThrottleConfig holds the configuration for throttling events
    /// </summary>
    [OutputType]
    public sealed class WebhookThrottleConfigPatch
    {
        /// <summary>
        /// ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
        /// </summary>
        public readonly int Burst;
        /// <summary>
        /// ThrottleQPS maximum number of batches per second default 10 QPS
        /// </summary>
        public readonly int Qps;

        [OutputConstructor]
        private WebhookThrottleConfigPatch(
            int burst,

            int qps)
        {
            Burst = burst;
            Qps = qps;
        }
    }
}
