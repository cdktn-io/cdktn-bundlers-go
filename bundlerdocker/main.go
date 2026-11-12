// Container-based asset bundler for CDK Terrain
package bundlerdocker

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@cdktn/bundler-docker.BundlingFileAccess",
		reflect.TypeOf((*BundlingFileAccess)(nil)).Elem(),
		map[string]interface{}{
			"BIND_MOUNT": BundlingFileAccess_BIND_MOUNT,
			"VOLUME_COPY": BundlingFileAccess_VOLUME_COPY,
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/bundler-docker.DockerBundler",
		reflect.TypeOf((*DockerBundler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bundle", GoMethod: "Bundle"},
			_jsii_.MemberProperty{JsiiProperty: "bundlerKey", GoGetter: "BundlerKey"},
			_jsii_.MemberProperty{JsiiProperty: "outputFileName", GoGetter: "OutputFileName"},
		},
		func() interface{} {
			j := jsiiProxy_DockerBundler{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnIAssetBundler)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/bundler-docker.DockerBundlerProps",
		reflect.TypeOf((*DockerBundlerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/bundler-docker.DockerImage",
		reflect.TypeOf((*DockerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberMethod{JsiiMethod: "run", GoMethod: "Run"},
		},
		func() interface{} {
			return &jsiiProxy_DockerImage{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/bundler-docker.DockerNamedVolume",
		reflect.TypeOf((*DockerNamedVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/bundler-docker.DockerRunOptions",
		reflect.TypeOf((*DockerRunOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/bundler-docker.DockerVolume",
		reflect.TypeOf((*DockerVolume)(nil)).Elem(),
	)
}
