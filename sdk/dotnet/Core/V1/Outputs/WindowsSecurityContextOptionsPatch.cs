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
    /// WindowsSecurityContextOptions contain Windows-specific options and credentials.
    /// </summary>
    [OutputType]
    public sealed class WindowsSecurityContextOptionsPatch
    {
        /// <summary>
        /// GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
        /// </summary>
        public readonly string GmsaCredentialSpec;
        /// <summary>
        /// GMSACredentialSpecName is the name of the GMSA credential spec to use.
        /// </summary>
        public readonly string GmsaCredentialSpecName;
        /// <summary>
        /// HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.
        /// </summary>
        public readonly bool HostProcess;
        /// <summary>
        /// The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
        /// </summary>
        public readonly string RunAsUserName;

        [OutputConstructor]
        private WindowsSecurityContextOptionsPatch(
            string gmsaCredentialSpec,

            string gmsaCredentialSpecName,

            bool hostProcess,

            string runAsUserName)
        {
            GmsaCredentialSpec = gmsaCredentialSpec;
            GmsaCredentialSpecName = gmsaCredentialSpecName;
            HostProcess = hostProcess;
            RunAsUserName = runAsUserName;
        }
    }
}
