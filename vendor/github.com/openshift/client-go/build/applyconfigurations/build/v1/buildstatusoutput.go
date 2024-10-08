// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// BuildStatusOutputApplyConfiguration represents a declarative configuration of the BuildStatusOutput type for use
// with apply.
type BuildStatusOutputApplyConfiguration struct {
	To *BuildStatusOutputToApplyConfiguration `json:"to,omitempty"`
}

// BuildStatusOutputApplyConfiguration constructs a declarative configuration of the BuildStatusOutput type for use with
// apply.
func BuildStatusOutput() *BuildStatusOutputApplyConfiguration {
	return &BuildStatusOutputApplyConfiguration{}
}

// WithTo sets the To field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the To field is set to the value of the last call.
func (b *BuildStatusOutputApplyConfiguration) WithTo(value *BuildStatusOutputToApplyConfiguration) *BuildStatusOutputApplyConfiguration {
	b.To = value
	return b
}
