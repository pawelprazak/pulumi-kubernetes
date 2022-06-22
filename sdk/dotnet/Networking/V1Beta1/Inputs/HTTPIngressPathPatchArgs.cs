// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1
{

    /// <summary>
    /// HTTPIngressPath associates a path regex with a backend. Incoming urls matching the path are forwarded to the backend.
    /// </summary>
    public class HTTPIngressPathPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
        /// </summary>
        [Input("backend")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Networking.V1Beta1.IngressBackendPatchArgs>? Backend { get; set; }

        /// <summary>
        /// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// PathType determines the interpretation of the Path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is
        ///   done on a path element by element basis. A path element refers is the
        ///   list of labels in the path split by the '/' separator. A request is a
        ///   match for path p if every p is an element-wise prefix of p of the
        ///   request path. Note that if the last element of the path is a substring
        ///   of the last element in request path, it is not a match (e.g. /foo/bar
        ///   matches /foo/bar/baz, but does not match /foo/barbaz).
        /// * ImplementationSpecific: Interpretation of the Path matching is up to
        ///   the IngressClass. Implementations can treat this as a separate PathType
        ///   or treat it identically to Prefix or Exact path types.
        /// Implementations are required to support all path types. Defaults to ImplementationSpecific.
        /// </summary>
        [Input("pathType")]
        public Input<string>? PathType { get; set; }

        public HTTPIngressPathPatchArgs()
        {
        }
    }
}
