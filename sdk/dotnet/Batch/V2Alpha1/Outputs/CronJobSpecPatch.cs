// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Batch.V2Alpha1
{

    /// <summary>
    /// CronJobSpec describes how the job execution will look like and when it will actually run.
    /// </summary>
    [OutputType]
    public sealed class CronJobSpecPatch
    {
        /// <summary>
        /// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
        /// </summary>
        public readonly string ConcurrencyPolicy;
        /// <summary>
        /// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
        /// </summary>
        public readonly int FailedJobsHistoryLimit;
        /// <summary>
        /// Specifies the job that will be created when executing a CronJob.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Batch.V2Alpha1.JobTemplateSpecPatch JobTemplate;
        /// <summary>
        /// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
        /// </summary>
        public readonly int StartingDeadlineSeconds;
        /// <summary>
        /// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
        /// </summary>
        public readonly int SuccessfulJobsHistoryLimit;
        /// <summary>
        /// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
        /// </summary>
        public readonly bool Suspend;

        [OutputConstructor]
        private CronJobSpecPatch(
            string concurrencyPolicy,

            int failedJobsHistoryLimit,

            Pulumi.Kubernetes.Types.Outputs.Batch.V2Alpha1.JobTemplateSpecPatch jobTemplate,

            string schedule,

            int startingDeadlineSeconds,

            int successfulJobsHistoryLimit,

            bool suspend)
        {
            ConcurrencyPolicy = concurrencyPolicy;
            FailedJobsHistoryLimit = failedJobsHistoryLimit;
            JobTemplate = jobTemplate;
            Schedule = schedule;
            StartingDeadlineSeconds = startingDeadlineSeconds;
            SuccessfulJobsHistoryLimit = successfulJobsHistoryLimit;
            Suspend = suspend;
        }
    }
}
