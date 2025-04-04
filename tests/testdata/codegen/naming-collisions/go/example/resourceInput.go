// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"naming-collisions/example/internal"
)

type ResourceInputResource struct {
	pulumi.CustomResourceState

	Bar pulumi.StringPtrOutput `pulumi:"bar"`
}

// NewResourceInputResource registers a new resource with the given unique name, arguments, and options.
func NewResourceInputResource(ctx *pulumi.Context,
	name string, args *ResourceInputResourceArgs, opts ...pulumi.ResourceOption) (*ResourceInputResource, error) {
	if args == nil {
		args = &ResourceInputResourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceInputResource
	err := ctx.RegisterResource("example::ResourceInput", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceInputResource gets an existing ResourceInputResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceInputResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceInputResourceState, opts ...pulumi.ResourceOption) (*ResourceInputResource, error) {
	var resource ResourceInputResource
	err := ctx.ReadResource("example::ResourceInput", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceInputResource resources.
type resourceInputResourceState struct {
}

type ResourceInputResourceState struct {
}

func (ResourceInputResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInputResourceState)(nil)).Elem()
}

type resourceInputResourceArgs struct {
}

// The set of arguments for constructing a ResourceInputResource resource.
type ResourceInputResourceArgs struct {
}

func (ResourceInputResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceInputResourceArgs)(nil)).Elem()
}

type ResourceInputResourceInput interface {
	pulumi.Input

	ToResourceInputResourceOutput() ResourceInputResourceOutput
	ToResourceInputResourceOutputWithContext(ctx context.Context) ResourceInputResourceOutput
}

func (*ResourceInputResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInputResource)(nil)).Elem()
}

func (i *ResourceInputResource) ToResourceInputResourceOutput() ResourceInputResourceOutput {
	return i.ToResourceInputResourceOutputWithContext(context.Background())
}

func (i *ResourceInputResource) ToResourceInputResourceOutputWithContext(ctx context.Context) ResourceInputResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceInputResourceOutput)
}

type ResourceInputResourceOutput struct{ *pulumi.OutputState }

func (ResourceInputResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceInputResource)(nil)).Elem()
}

func (o ResourceInputResourceOutput) ToResourceInputResourceOutput() ResourceInputResourceOutput {
	return o
}

func (o ResourceInputResourceOutput) ToResourceInputResourceOutputWithContext(ctx context.Context) ResourceInputResourceOutput {
	return o
}

func (o ResourceInputResourceOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceInputResource) pulumi.StringPtrOutput { return v.Bar }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceInputResourceInput)(nil)).Elem(), &ResourceInputResource{})
	pulumi.RegisterOutputType(ResourceInputResourceOutput{})
}
