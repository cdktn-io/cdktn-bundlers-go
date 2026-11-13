// Host-process asset bundler for CDK Terrain
package bundlerlocal

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/bundler-local.LocalBundler",
		reflect.TypeOf((*LocalBundler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bundle", GoMethod: "Bundle"},
			_jsii_.MemberProperty{JsiiProperty: "bundlerKey", GoGetter: "BundlerKey"},
			_jsii_.MemberProperty{JsiiProperty: "outputFileName", GoGetter: "OutputFileName"},
		},
		func() interface{} {
			j := jsiiProxy_LocalBundler{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnIAssetBundler)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/bundler-local.LocalBundlerProps",
		reflect.TypeOf((*LocalBundlerProps)(nil)).Elem(),
	)
}
