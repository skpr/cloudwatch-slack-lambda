package cloudwatch

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

const (
	// TagKeyCluster is a key to identify the cluster which an alarm relates to.
	TagKeyCluster = "skpr.io/cluster"
	// TagKeyProject is a key to identify the project which an alarm relates to.
	TagKeyProject = "skpr.io/project"
	// TagKeyEnvironment is a key to identify the environment which an alarm relates to.
	TagKeyEnvironment = "skpr.io/environment"
	// TagKeyInstance is a key to identify the instance which an alarm relates to.
	TagKeyInstance = "skpr.io/instance"
	// TagKeyLinkDashboard is a key to get the dashboard related to an alarm.
	TagKeyLinkDashboard = "skpr.io/link/dashboard"
	// TagKeyLinkDocumentation is a key to get the documentation related to an alarm.
	TagKeyLinkDocumentation = "skpr.io/link/documentation"
	// TagKeyAssetIcon is a key to get the icon related to an alarm.
	TagKeyAssetIcon = "skpr.io/asset/icon"
)

// GetValueFromTag returns the value of a tag with the given key, or false if the tag does not exist.
func GetValueFromTag(tags []types.Tag, key string) string {
	for _, tag := range tags {
		if *tag.Key == key {
			return *tag.Value
		}
	}

	return ""
}
