// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Text.Json;
using Pulumi;
using Pulumi.Kubernetes.Yaml;
using Pulumi.Utilities;

namespace Pulumi.Kubernetes.Helm
{
    public abstract class ChartBase : CollectionComponentResource
    {
        /// <summary>
        /// Create an instance of the specified Helm chart.
        /// </summary>
        /// <param name="releaseName">Name of the Chart (e.g., nginx-ingress).</param>
        /// <param name="args">Configuration options for the Chart.</param>
        /// <param name="options">A bag of options that control this resource's behavior.</param>
        protected ChartBase(string releaseName, Union<ChartArgs, LocalChartArgs> args, ComponentResourceOptions? options = null)
            : base("kubernetes:helm.sh/v2:Chart", GetName(args, releaseName), options)
        {
            releaseName = GetName(args, releaseName);
            var config = args.Unwrap();

            var configDeps = Output.Create(OutputUtilities.GetDependenciesAsync(config));
            OutputUtilities.GetIsKnownAsync(config).ContinueWith(isKnown =>
            {
                if (!isKnown.Result)
                {
                    // Note that this can only happen during a preview.
                    Log.Info("[Can't preview] all chart values must be known ahead of time to generate an accurate preview.", this);
                }
            });

            var resources = Output.Tuple(config, configDeps).Apply(values =>
            {
                var chartArgs = values.Item1;
                var dependencies = values.Item2;

                // Create temporary directories and files to hold chart data and override values.
                var overrides = Path.GetTempFileName();
                var chartDirectoryName = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
                var chartDirectory = Directory.CreateDirectory(chartDirectoryName);

                try
                {
                    string chart;
                    string defaultValues;
                    BaseChartArgsUnwrap cfgBase;
                    if (chartArgs.IsT0)
                    {
                        var cfg = chartArgs.AsT0;
                        // Fetch chart.
                        if (cfg.Repo != null && cfg.Repo.Contains("http"))
                        {
                            throw new Exception(
                                $"`{nameof(cfg.Repo)}` specifies the name of the Helm chart repo. Use `{nameof(ChartArgs)}.{nameof(cfg.Repo)}` to specify a URL.");
                        }

                        var chartToFetch = !string.IsNullOrEmpty(cfg.Repo) ? $"{cfg.Repo}/{cfg.Chart}" : cfg.Chart;
                        var fetchOptions = cfg.FetchOptions ?? new ChartFetchArgsUnwrap();
                        fetchOptions.Destination = chartDirectoryName;
                        fetchOptions.Version = cfg.Version;
                        Fetch(chartToFetch, fetchOptions);
                        // Sort the directories into alphabetical order, and choose the first
                        var fetchedChart = chartDirectory.GetDirectories().OrderBy(x => x.Name).First();
                        var fetchedChartName = fetchedChart.Name;
                        chart = fetchedChart.FullName;
                        defaultValues = Path.Join(chartDirectoryName, fetchedChartName, "values.yaml");
                        cfgBase = cfg;
                    }
                    else
                    {
                        var cfg = chartArgs.AsT1;
                        chart = cfg.Path;
                        defaultValues = Path.Join(chart, "values.yaml");
                        cfgBase = cfg;
                    }

                    // Write overrides file.
                    var data = JsonSerializer.Serialize(cfgBase.Values);
                    File.WriteAllText(overrides, data);

                    // Does not require Tiller. From the `helm template` documentation:
                    //
                    // >  Render chart templates locally and display the output.
                    // >
                    // > This does not require Tiller. However, any values that would normally be
                    // > looked up or retrieved in-cluster will be faked locally. Additionally, none
                    // > of the server-side testing of chart validity (e.g. whether an API is supported)
                    // > is done.
                    var flags = new List<string>
                    {
                        "template", chart,
                        "--name-template", releaseName,
                        "--values", defaultValues,
                        "--values", overrides
                    };
                    if (cfgBase.ApiVersions.Length > 0)
                    {
                        foreach (string version in cfgBase.ApiVersions)
                        {
                            flags.Add($"--api-versions={version}");
                        }
                    }

                    if (!string.IsNullOrEmpty(cfgBase.Namespace))
                    {
                        flags.Add("--namespace");
                        flags.Add(cfgBase.Namespace);
                    }

                    if (IsHelmV3())
                    {
                        flags.Add("--include-crds");
                    }

                    var yaml = ExecuteCommand("helm", flags);
                    return ParseTemplate(
                        yaml, cfgBase.Transformations, cfgBase.ResourcePrefix, dependencies, cfgBase.Namespace, options?.Provider);
                }
                catch (Exception e)
                {
                    // Shed stack trace, only emit the error.
                    throw new ResourceException(e.Message, this);
                }
                finally
                {
                    chartDirectory.Delete(true);
                }
            });
            RegisterResources(resources);
        }

        private static string GetName(Union<ChartArgs, LocalChartArgs> config, string releaseName)
        {
            var prefix = config.Match(v => v.ResourcePrefix, v => v.ResourcePrefix);
            return string.IsNullOrEmpty(prefix) ? releaseName : $"{prefix}-{releaseName}";
        }

        private static bool IsHelmV3()
        {
            string[] flags = { "version", "--short" };

            // Helm v2 returns version like this:
            // Client: v2.16.7+g5f2584f
            // Helm v3 returns a version like this:
            // v3.1.2+gd878d4d
            // --include-crds is available in helm v3.1+ so check for a regex matching that version
            var version = ExecuteCommand("helm", flags);
            Regex r = new Regex(@"^v3\.[1-9]");
            return r.IsMatch(version);
        }

        private void Fetch(string chart, ChartFetchArgsUnwrap opts)
        {
            var flags = new List<string> { "fetch", chart };

            // Untar by default.
            if (opts.Untar != false)
            {
                flags.Add("--untar");
            }

            Dictionary<string, string>? env = null;

            // Helm v3 removed the `--home` flag, so we must use an env var instead.
            if (!string.IsNullOrEmpty(opts.Home))
            {
                env = new Dictionary<string, string>();
                env["HELM_HOME"] = opts.Home;
            }

            if (!string.IsNullOrEmpty(opts.Version))
            {
                flags.Add("--version");
                flags.Add(opts.Version);
            }
            if (!string.IsNullOrEmpty(opts.CAFile))
            {
                flags.Add("--ca-file");
                flags.Add(opts.CAFile);
            }
            if (!string.IsNullOrEmpty(opts.CertFile))
            {
                flags.Add("--cert-file");
                flags.Add(opts.CertFile);
            }
            if (!string.IsNullOrEmpty(opts.KeyFile))
            {
                flags.Add("--key-file");
                flags.Add(opts.KeyFile);
            }
            if (!string.IsNullOrEmpty(opts.Destination))
            {
                flags.Add("--destination");
                flags.Add(opts.Destination);
            }
            if (!string.IsNullOrEmpty(opts.Keyring))
            {
                flags.Add("--keyring");
                flags.Add(opts.Keyring);
            }
            if (!string.IsNullOrEmpty(opts.Password))
            {
                flags.Add("--password");
                flags.Add(opts.Password);
            }
            if (!string.IsNullOrEmpty(opts.Repo))
            {
                flags.Add("--repo");
                flags.Add(opts.Repo);
            }
            if (!string.IsNullOrEmpty(opts.UntarDir))
            {
                flags.Add("--untardir");
                flags.Add(opts.UntarDir);
            }
            if (!string.IsNullOrEmpty(opts.Username))
            {
                flags.Add("--username");
                flags.Add(opts.Username);
            }
            if (opts.Devel == true)
            {
                flags.Add("--devel");
            }
            if (opts.Prov == true)
            {
                flags.Add("--prov");
            }
            if (opts.Verify == true)
            {
                flags.Add("--verify");
            }

            ExecuteCommand("helm", flags, env);
        }

        private Output<ImmutableDictionary<string, KubernetesResource>> ParseTemplate(string text,
            List<TransformationAction> transformations, string? resourcePrefix, ImmutableHashSet<Resource> dependsOn,
            string? defaultNamespace, Pulumi.ProviderResource provider)
        {
            return Yaml.Invokes
                .YamlDecode(new YamlDecodeArgs { Text = text, DefaultNamespace = defaultNamespace }, new InvokeOptions { Provider = provider })
                .Apply(objs =>
                {
                    var args = new ConfigGroupArgs
                    {
                        ResourcePrefix = resourcePrefix,
                        Objs = objs,
                        Transformations = transformations
                    };
                    var opts = new ComponentResourceOptions { Parent = this, DependsOn = dependsOn.ToArray(), Provider = provider };
                    return Parser.Parse(args, opts);
                });
        }

        private static string ExecuteCommand(string command, IEnumerable<string> flags, Dictionary<string, string>? env = null)
        {
            using var process = new Process
            {
                StartInfo =
                {
                    FileName = command,
                    Arguments = EscapeArguments(flags),
                    RedirectStandardOutput = true,
                    RedirectStandardError = true
                }
            };

            if (env != null)
            {
                foreach (KeyValuePair<string, string> value in env)
                {
                    process.StartInfo.EnvironmentVariables[value.Key] = value.Value;
                }
            }

            process.Start();
            string output = process.StandardOutput.ReadToEnd();
            process.WaitForExit();
            if (process.ExitCode > 0)
            {
                string error = process.StandardError.ReadToEnd();
                throw new Exception(error);
            }
            return output;
        }

        /// <summary>
        /// Convert an argument array to an argument string for using with Process.StartInfo.Arguments.
        /// </summary>
        private static string EscapeArguments(IEnumerable<string> args)
            => string.Join(" ", args.Select(a => EscapeArguments(a)));

        /// <summary>
        /// Convert argument to an argument string for using with Process.StartInfo.Arguments.
        /// </summary>
        private static string EscapeArguments(string argument)
        {
            var escapedArgument = new StringBuilder();
            var backslashCount = 0;
            var needsQuotes = false;

            foreach (var character in argument)
            {
                switch (character)
                {
                    case '\\':
                        // Backslashes are simply passed through, except when they need
                        // to be escaped when followed by a \", e.g. the argument string
                        // \", which would be encoded to \\\"
                        backslashCount++;
                        escapedArgument.Append('\\');
                        break;

                    case '\"':
                        // Escape any preceding backslashes
                        escapedArgument.Append('\\', backslashCount);

                        // Append an escaped double quote.
                        escapedArgument.Append("\\\"");

                        // Reset the backslash counter.
                        backslashCount = 0;
                        break;

                    case ' ':
                    case '\t':
                        // White spaces are escaped by surrounding the entire string with
                        // double quotes, which should be done at the end to prevent
                        // multiple wrappings.
                        needsQuotes = true;

                        // Append the whitespace
                        escapedArgument.Append(character);

                        // Reset the backslash counter.
                        backslashCount = 0;
                        break;

                    default:
                        // Reset the backslash counter.
                        backslashCount = 0;

                        // Append the current character
                        escapedArgument.Append(character);
                        break;
                }
            }

            // No need to wrap in quotes
            if (!needsQuotes)
            {
                return escapedArgument.ToString();
            }

            // Prepend the "
            escapedArgument.Insert(0, '"');

            // Escape any preceding backslashes before appending the "
            escapedArgument.Append('\\', backslashCount);

            // Append the final "
            escapedArgument.Append('\"');

            return escapedArgument.ToString();
        }
    }

    public class BaseChartArgs : ResourceArgs
    {
        private InputList<string>? _apiVersions;

        /// <summary>
        /// The optional kubernetes api versions used for Capabilities.APIVersions.
        /// </summary>
        public InputList<string> ApiVersions
        {
            get => _apiVersions ??= new InputList<string>();
            set => _apiVersions = value;
        }

        /// <summary>
        /// By default, Helm resources with the 'test', 'test-success', and 'test-failure' hooks are not installed. Set
        /// this flag to true to include these resources.
        /// </summary>
        public Input<bool>? IncludeTestHookResources { get; set; }

        // <summary>
        // By default, CRDs are rendered along with Helm chart templates. Setting this to true will skip CRD rendering.
        // </summary>
        public Input<bool>? SkipCRDRendering { get; set; }

        // <summary>
        // Skip await logic for all resources in this Chart. Resources will be marked ready as soon as they are created.
        // Warning: This option should not be used if you have resources depending on Outputs from the Chart.
        // </summary>
        public Input<bool>? SkipAwait { get; set; }

        /// <summary>
        /// The optional namespace to install chart resources into.
        /// </summary>
        public Input<string>? Namespace { get; set; }

        private InputMap<object>? _values;

        /// <summary>
        /// Overrides for chart values.
        /// </summary>
        public InputMap<object> Values
        {
            get => _values ??= new InputMap<object>();
            set => _values = value;
        }

        private List<TransformationAction>? _transformations;

        /// <summary>
        /// Optional array of transformations to apply to resources that will be created by this chart prior to
        /// creation. Allows customization of the chart behaviour without directly modifying the chart itself.
        /// </summary>
        public List<TransformationAction> Transformations
        {
            get => _transformations ??= new List<TransformationAction>();
            set => _transformations = value;
        }

        /// <summary>
        /// An optional prefix for the auto-generated resource names.
        /// Example: A resource created with resourcePrefix="foo" would produce a resource named "foo-resourceName".
        /// </summary>
        public string? ResourcePrefix { get; set; }
    }

    public class ChartArgs : BaseChartArgs
    {
        /// <summary>
        /// The repository name of the chart to deploy.
        /// Example: "stable"
        /// </summary>
        public Input<string>? Repo { get; set; }

        /// <summary>
        /// The name of the chart to deploy.  If <see cref="Repo" /> is provided, this chart
        /// name will be prefixed by the repo name.
        /// Example: Repo: "stable", Chart: "nginx-ingress" -> "stable/nginx-ingress"
        /// Example: Chart: "stable/nginx-ingress" -> "stable/nginx-ingress"
        /// </summary>
        public Input<string> Chart { get; set; } = null!;

        /// <summary>
        /// The version of the chart to deploy. If not provided, the latest version will be deployed.
        /// </summary>
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Additional options to customize the fetching of the Helm chart.
        /// </summary>
        public Input<ChartFetchArgs>? FetchOptions { get; set; }
    }

    public class LocalChartArgs : BaseChartArgs
    {
        /// <summary>
        /// The path to the chart directory which contains the `Chart.yaml` file.
        /// </summary>
        public string Path { get; set; } = null!;
    }

    /// <summary>
    /// Additional options to customize the fetching of the Helm chart.
    /// </summary>
    public class ChartFetchArgs
    {
        /// <summary>
        /// Specific version of a chart. Without this, the latest version is fetched.
        /// </summary>
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Verify certificates of HTTPS-enabled servers using this CA bundle.
        /// </summary>
        public Input<string>? CAFile { get; set; }

        /// <summary>
        /// Identify HTTPS client using this SSL certificate file.
        /// </summary>
        public Input<string>? CertFile { get; set; }

        /// <summary>
        /// Identify HTTPS client using this SSL key file.
        /// </summary>
        public Input<string>? KeyFile { get; set; }

        /// <summary>
        /// Location to write the chart. If this and tardir are specified, tardir is appended to this
        /// (default ".").
        /// </summary>
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Keyring containing public keys (default "/Users/alex/.gnupg/pubring.gpg").
        /// </summary>
        public Input<string>? Keyring { get; set; }

        /// <summary>
        /// Chart repository password.
        /// </summary>
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Chart repository url where to locate the requested chart.
        /// </summary>
        public Input<string>? Repo { get; set; }

        /// <summary>
        /// If untar is specified, this flag specifies the name of the directory into which the chart is
        /// expanded (default ".").
        /// </summary>
        public Input<string>? UntarDir { get; set; }

        /// <summary>
        /// Chart repository username.
        /// </summary>
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Location of your Helm config. Overrides $HELM_HOME (default "/Users/alex/.helm").
        /// </summary>
        public Input<string>? Home { get; set; }

        /// <summary>
        /// Use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is
        /// ignored.
        /// </summary>
        public Input<bool>? Devel { get; set; }

        /// <summary>
        /// Fetch the provenance file, but don't perform verification.
        /// </summary>
        public Input<bool>? Prov { get; set; }

        /// <summary>
        /// If set to false, will leave the chart as a tarball after downloading.
        /// </summary>
        public Input<bool>? Untar { get; set; }

        /// <summary>
        /// Verify the package against its signature.
        /// </summary>
        public Input<bool>? Verify { get; set; }
    }
}
