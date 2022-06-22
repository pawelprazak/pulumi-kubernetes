// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Events.V1
{
    /// <summary>
    /// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
    /// </summary>
    [KubernetesResourceType("kubernetes:events.k8s.io/v1:EventPatch")]
    public partial class EventPatch : KubernetesResource
    {
        /// <summary>
        /// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Output("deprecatedCount")]
        public Output<int> DeprecatedCount { get; private set; } = null!;

        /// <summary>
        /// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Output("deprecatedFirstTimestamp")]
        public Output<string> DeprecatedFirstTimestamp { get; private set; } = null!;

        /// <summary>
        /// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Output("deprecatedLastTimestamp")]
        public Output<string> DeprecatedLastTimestamp { get; private set; } = null!;

        /// <summary>
        /// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Output("deprecatedSource")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Core.V1.EventSource> DeprecatedSource { get; private set; } = null!;

        /// <summary>
        /// eventTime is the time when this Event was first observed. It is required.
        /// </summary>
        [Output("eventTime")]
        public Output<string> EventTime { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
        /// </summary>
        [Output("note")]
        public Output<string> Note { get; private set; } = null!;

        /// <summary>
        /// reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
        /// </summary>
        [Output("reason")]
        public Output<string> Reason { get; private set; } = null!;

        /// <summary>
        /// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
        /// </summary>
        [Output("regarding")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference> Regarding { get; private set; } = null!;

        /// <summary>
        /// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
        /// </summary>
        [Output("related")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Core.V1.ObjectReference> Related { get; private set; } = null!;

        /// <summary>
        /// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
        /// </summary>
        [Output("reportingController")]
        public Output<string> ReportingController { get; private set; } = null!;

        /// <summary>
        /// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
        /// </summary>
        [Output("reportingInstance")]
        public Output<string> ReportingInstance { get; private set; } = null!;

        /// <summary>
        /// series is data about the Event series this event represents or nil if it's a singleton Event.
        /// </summary>
        [Output("series")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Events.V1.EventSeries> Series { get; private set; } = null!;

        /// <summary>
        /// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EventPatch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventPatch(string name, Pulumi.Kubernetes.Types.Inputs.Events.V1.EventPatchArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:events.k8s.io/v1:EventPatch", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal EventPatch(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:events.k8s.io/v1:EventPatch", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private EventPatch(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:events.k8s.io/v1:EventPatch", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Events.V1.EventPatchArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Events.V1.EventPatchArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Events.V1.EventPatchArgs();
            args.ApiVersion = "events.k8s.io/v1";
            args.Kind = "Event";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "kubernetes:core/v1:Event"},
                    new Pulumi.Alias { Type = "kubernetes:events.k8s.io/v1beta1:Event"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventPatch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventPatch Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventPatch(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Events.V1
{

    public class EventPatchArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Input("deprecatedCount")]
        public Input<int>? DeprecatedCount { get; set; }

        /// <summary>
        /// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Input("deprecatedFirstTimestamp")]
        public Input<string>? DeprecatedFirstTimestamp { get; set; }

        /// <summary>
        /// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Input("deprecatedLastTimestamp")]
        public Input<string>? DeprecatedLastTimestamp { get; set; }

        /// <summary>
        /// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
        /// </summary>
        [Input("deprecatedSource")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.EventSourceArgs>? DeprecatedSource { get; set; }

        /// <summary>
        /// eventTime is the time when this Event was first observed. It is required.
        /// </summary>
        [Input("eventTime")]
        public Input<string>? EventTime { get; set; }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
        /// </summary>
        [Input("note")]
        public Input<string>? Note { get; set; }

        /// <summary>
        /// reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
        /// </summary>
        [Input("regarding")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ObjectReferenceArgs>? Regarding { get; set; }

        /// <summary>
        /// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
        /// </summary>
        [Input("related")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ObjectReferenceArgs>? Related { get; set; }

        /// <summary>
        /// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
        /// </summary>
        [Input("reportingController")]
        public Input<string>? ReportingController { get; set; }

        /// <summary>
        /// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
        /// </summary>
        [Input("reportingInstance")]
        public Input<string>? ReportingInstance { get; set; }

        /// <summary>
        /// series is data about the Event series this event represents or nil if it's a singleton Event.
        /// </summary>
        [Input("series")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Events.V1.EventSeriesArgs>? Series { get; set; }

        /// <summary>
        /// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EventPatchArgs()
        {
        }
    }
}
