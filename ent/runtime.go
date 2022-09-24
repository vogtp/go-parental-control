// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/vogtp/go-win-session/ent/activity"
	"github.com/vogtp/go-win-session/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	activityFields := schema.Activity{}.Fields()
	_ = activityFields
	// activityDescActivity is the schema descriptor for activity field.
	activityDescActivity := activityFields[1].Descriptor()
	// activity.DefaultActivity holds the default value on creation for the activity field.
	activity.DefaultActivity = activityDescActivity.Default.(int64)
}