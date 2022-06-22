// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Core.V1
{

    /// <summary>
    /// A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates the scope name and values.
    /// </summary>
    public class ScopedResourceSelectorRequirementPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// The name of the scope that the selector applies to.
        /// </summary>
        [Input("scopeName")]
        public Input<string>? ScopeName { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ScopedResourceSelectorRequirementPatchArgs()
        {
        }
    }
}
