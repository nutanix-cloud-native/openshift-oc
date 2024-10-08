// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "github.com/openshift/api/config/v1alpha1"
	configv1alpha1 "github.com/openshift/client-go/config/applyconfigurations/config/v1alpha1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ImagePoliciesGetter has a method to return a ImagePolicyInterface.
// A group's client should implement this interface.
type ImagePoliciesGetter interface {
	ImagePolicies(namespace string) ImagePolicyInterface
}

// ImagePolicyInterface has methods to work with ImagePolicy resources.
type ImagePolicyInterface interface {
	Create(ctx context.Context, imagePolicy *v1alpha1.ImagePolicy, opts v1.CreateOptions) (*v1alpha1.ImagePolicy, error)
	Update(ctx context.Context, imagePolicy *v1alpha1.ImagePolicy, opts v1.UpdateOptions) (*v1alpha1.ImagePolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, imagePolicy *v1alpha1.ImagePolicy, opts v1.UpdateOptions) (*v1alpha1.ImagePolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ImagePolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ImagePolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImagePolicy, err error)
	Apply(ctx context.Context, imagePolicy *configv1alpha1.ImagePolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ImagePolicy, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, imagePolicy *configv1alpha1.ImagePolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ImagePolicy, err error)
	ImagePolicyExpansion
}

// imagePolicies implements ImagePolicyInterface
type imagePolicies struct {
	*gentype.ClientWithListAndApply[*v1alpha1.ImagePolicy, *v1alpha1.ImagePolicyList, *configv1alpha1.ImagePolicyApplyConfiguration]
}

// newImagePolicies returns a ImagePolicies
func newImagePolicies(c *ConfigV1alpha1Client, namespace string) *imagePolicies {
	return &imagePolicies{
		gentype.NewClientWithListAndApply[*v1alpha1.ImagePolicy, *v1alpha1.ImagePolicyList, *configv1alpha1.ImagePolicyApplyConfiguration](
			"imagepolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.ImagePolicy { return &v1alpha1.ImagePolicy{} },
			func() *v1alpha1.ImagePolicyList { return &v1alpha1.ImagePolicyList{} }),
	}
}
