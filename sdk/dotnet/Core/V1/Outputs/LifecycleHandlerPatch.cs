// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
    /// </summary>
    [OutputType]
    public sealed class LifecycleHandlerPatch
    {
        /// <summary>
        /// Exec specifies the action to take.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.ExecActionPatch Exec;
        /// <summary>
        /// HTTPGet specifies the http request to perform.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.HTTPGetActionPatch HttpGet;
        /// <summary>
        /// Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Core.V1.TCPSocketActionPatch TcpSocket;

        [OutputConstructor]
        private LifecycleHandlerPatch(
            Pulumi.Kubernetes.Types.Outputs.Core.V1.ExecActionPatch exec,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.HTTPGetActionPatch httpGet,

            Pulumi.Kubernetes.Types.Outputs.Core.V1.TCPSocketActionPatch tcpSocket)
        {
            Exec = exec;
            HttpGet = httpGet;
            TcpSocket = tcpSocket;
        }
    }
}
