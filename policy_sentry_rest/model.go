package policy_sentry_rest


type ActionsForServicesAtAccessLevel struct {
	Read                  []*string
	Write                 []*string
	Tagging               []*string
	List                  []*string
	PermissionsManagement []*string
}

type ActionsForResourcesWithoutResourceConstraints struct {
	Read                  []*string
	Write                 []*string
	Tagging               []*string
	List                  []*string
	PermissionsManagement []*string
	SingleActions         []*string
}

type PolicyDocumentInput struct {
	ExcludeActions      []*string
	ActionsForResources *ActionsForResourcesWithoutResourceConstraints
	ActionsForServices  *ActionsForServicesAtAccessLevel
	Overrides           *Overrides
}

type Overrides struct {
	SkipResourceConstraints []*string
}

type PolicyDocument struct {
	Statement []struct {
		Action   []string `json:"Action"`
		Effect   string   `json:"Effect"`
		Resource []string `json:"Resource"`
		Sid      string   `json:"Sid"`
	} `json:"Statement"`
	Version string `json:"Version"`
}