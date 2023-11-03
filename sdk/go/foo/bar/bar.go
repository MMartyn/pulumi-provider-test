// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bar

import (
	"context"
	"reflect"

	"errors"
	"github.com/mmartyn/pulumi-foo/sdk/go/foo/internal"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Bar struct {
	pulumi.ResourceState

	Bucket     s3.BucketOutput     `pulumi:"bucket"`
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
}

// NewBar registers a new resource with the given unique name, arguments, and options.
func NewBar(ctx *pulumi.Context,
	name string, args *BarArgs, opts ...pulumi.ResourceOption) (*Bar, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bar
	err := ctx.RegisterRemoteComponentResource("foo:bar:Bar", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type barArgs struct {
	BucketName string `pulumi:"bucketName"`
}

// The set of arguments for constructing a Bar resource.
type BarArgs struct {
	BucketName pulumi.StringInput
}

func (BarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*barArgs)(nil)).Elem()
}

type BarInput interface {
	pulumi.Input

	ToBarOutput() BarOutput
	ToBarOutputWithContext(ctx context.Context) BarOutput
}

func (*Bar) ElementType() reflect.Type {
	return reflect.TypeOf((**Bar)(nil)).Elem()
}

func (i *Bar) ToBarOutput() BarOutput {
	return i.ToBarOutputWithContext(context.Background())
}

func (i *Bar) ToBarOutputWithContext(ctx context.Context) BarOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BarOutput)
}

func (i *Bar) ToOutput(ctx context.Context) pulumix.Output[*Bar] {
	return pulumix.Output[*Bar]{
		OutputState: i.ToBarOutputWithContext(ctx).OutputState,
	}
}

// BarArrayInput is an input type that accepts BarArray and BarArrayOutput values.
// You can construct a concrete instance of `BarArrayInput` via:
//
//	BarArray{ BarArgs{...} }
type BarArrayInput interface {
	pulumi.Input

	ToBarArrayOutput() BarArrayOutput
	ToBarArrayOutputWithContext(context.Context) BarArrayOutput
}

type BarArray []BarInput

func (BarArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bar)(nil)).Elem()
}

func (i BarArray) ToBarArrayOutput() BarArrayOutput {
	return i.ToBarArrayOutputWithContext(context.Background())
}

func (i BarArray) ToBarArrayOutputWithContext(ctx context.Context) BarArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BarArrayOutput)
}

func (i BarArray) ToOutput(ctx context.Context) pulumix.Output[[]*Bar] {
	return pulumix.Output[[]*Bar]{
		OutputState: i.ToBarArrayOutputWithContext(ctx).OutputState,
	}
}

// BarMapInput is an input type that accepts BarMap and BarMapOutput values.
// You can construct a concrete instance of `BarMapInput` via:
//
//	BarMap{ "key": BarArgs{...} }
type BarMapInput interface {
	pulumi.Input

	ToBarMapOutput() BarMapOutput
	ToBarMapOutputWithContext(context.Context) BarMapOutput
}

type BarMap map[string]BarInput

func (BarMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bar)(nil)).Elem()
}

func (i BarMap) ToBarMapOutput() BarMapOutput {
	return i.ToBarMapOutputWithContext(context.Background())
}

func (i BarMap) ToBarMapOutputWithContext(ctx context.Context) BarMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BarMapOutput)
}

func (i BarMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Bar] {
	return pulumix.Output[map[string]*Bar]{
		OutputState: i.ToBarMapOutputWithContext(ctx).OutputState,
	}
}

type BarOutput struct{ *pulumi.OutputState }

func (BarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bar)(nil)).Elem()
}

func (o BarOutput) ToBarOutput() BarOutput {
	return o
}

func (o BarOutput) ToBarOutputWithContext(ctx context.Context) BarOutput {
	return o
}

func (o BarOutput) ToOutput(ctx context.Context) pulumix.Output[*Bar] {
	return pulumix.Output[*Bar]{
		OutputState: o.OutputState,
	}
}

func (o BarOutput) Bucket() s3.BucketOutput {
	return o.ApplyT(func(v *Bar) s3.BucketOutput { return v.Bucket }).(s3.BucketOutput)
}

func (o BarOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *Bar) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

type BarArrayOutput struct{ *pulumi.OutputState }

func (BarArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bar)(nil)).Elem()
}

func (o BarArrayOutput) ToBarArrayOutput() BarArrayOutput {
	return o
}

func (o BarArrayOutput) ToBarArrayOutputWithContext(ctx context.Context) BarArrayOutput {
	return o
}

func (o BarArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Bar] {
	return pulumix.Output[[]*Bar]{
		OutputState: o.OutputState,
	}
}

func (o BarArrayOutput) Index(i pulumi.IntInput) BarOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bar {
		return vs[0].([]*Bar)[vs[1].(int)]
	}).(BarOutput)
}

type BarMapOutput struct{ *pulumi.OutputState }

func (BarMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bar)(nil)).Elem()
}

func (o BarMapOutput) ToBarMapOutput() BarMapOutput {
	return o
}

func (o BarMapOutput) ToBarMapOutputWithContext(ctx context.Context) BarMapOutput {
	return o
}

func (o BarMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Bar] {
	return pulumix.Output[map[string]*Bar]{
		OutputState: o.OutputState,
	}
}

func (o BarMapOutput) MapIndex(k pulumi.StringInput) BarOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bar {
		return vs[0].(map[string]*Bar)[vs[1].(string)]
	}).(BarOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BarInput)(nil)).Elem(), &Bar{})
	pulumi.RegisterInputType(reflect.TypeOf((*BarArrayInput)(nil)).Elem(), BarArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BarMapInput)(nil)).Elem(), BarMap{})
	pulumi.RegisterOutputType(BarOutput{})
	pulumi.RegisterOutputType(BarArrayOutput{})
	pulumi.RegisterOutputType(BarMapOutput{})
}
