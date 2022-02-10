package rulesengine

import "time"

// // UpdateDeployPolicyStateResponse -
// type UpdateDeployPolicyStateResponse struct {
// 	ID    string `json:"id,omitempty"`
// 	State string `json:"state,omitempty"`
// }

type Policy struct {
	Type        string `json:"@type,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	Platform    string `json:"platform,omitempty"`
	Rules       []Rule `json:"rules,omitempty"`
}

// PolicyResponse -
type PolicyResponse struct {
	Policy
	AtID       string    `json:"@id,omitempty"`
	ID         string    `json:"id,omitempty"`
	PolicyType string    `json:"policy_type,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	History    []History `json:"history,omitempty"`
}

type History struct {
	ID        int       `json:"id,omitempty"`
	Type      string    `json:"type,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	User      User      `json:"user,omitempty"`
}

// Rule -
type Rule struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Matches     []Match `json:"matches,omitempty"`
}

type RuleResponse struct {
	Rule
	Ordinal   int       `json:"ordinal,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Match -
type Match struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	Ordinal    int       `json:"ordinal,omitempty"`
	Value      string    `json:"value,omitempty"`
	Codes      string    `json:"codes,omitempty"`
	Compare    string    `json:"compare,omitempty"`
	Encoded    bool      `json:"encoded,omitempty"`
	Hostnames  string    `json:"hostnames,omitempty"`
	IgnoreCase bool      `json:"ignore-case,omitempty"`
	Name       string    `json:"name,omitempty"`
	RelativeTo string    `json:"relative-to,omitempty"`
	Result     string    `json:"result,omitempty"`
	Matches    []Match   `json:"matches,omitempty"`
	Features   []Feature `json:"features,omitempty"`
}

// Feature -
type Feature struct {
	Action          string   `json:"action,omitempty"`
	Code            string   `json:"code,omitempty"`
	Destination     string   `json:"destination,omitempty"`
	Enabled         bool     `json:"enabled,omitempty"`
	Expires         int      `json:"expires,omitempty"`
	Extensions      string   `json:"extensions,omitempty"`
	Format          string   `json:"format,omitempty"`
	HeaderName      string   `json:"header-name,omitempty"`
	HeaderValue     string   `json:"header-value,omitempty"`
	Instance        string   `json:"instance,omitempty"`
	KbytesPerSecond int      `json:"kbytes-per-second,omitempty"`
	MediaTypes      []string `json:"mediaTypes,omitempty"`
	Methods         string   `json:"methods,omitempty"`
	Milliseconds    int      `json:"milliseconds,omitempty"`
	Mode            string   `json:"mode,omitempty"`
	Name            string   `json:"name,omitempty"`
	Names           []string `json:"names,omitempty"`
	Parameters      string   `json:"parameters,omitempty"`
	PrebufSeconds   int      `json:"prebuf-seconds,omitempty"`
	Requests        int      `json:"requests,omitempty"`
	Seconds         int      `json:"seconds,omitempty"`
	SeekEnd         string   `json:"seekEnd,omitempty"`
	SeekStart       string   `json:"seekStart,omitempty"`
	Site            string   `json:"site,omitempty"`
	Source          string   `json:"source,omitempty"`
	Status          string   `json:"status,omitempty"`
	Type            string   `json:"type,omitempty"`
	Tags            string   `json:"tags,omitempty"`
	Treatment       string   `json:"treatment,omitempty"`
	Units           string   `json:"units,omitempty"`
	Value           string   `json:"value,omitempty"`
}

// SubmitDeployRequest -
type SubmitDeployRequest struct {
	PolicyID    int    `json:"policy_id"`
	Environment string `json:"environment,omitempty"`
	Message     string `json:"message"`
}

// DeployRequestOK -
type DeployRequestOK struct {
	ID          string                   `json:"id,omitempty"`
	AtID        string                   `json:"@id,omitempty"`
	Type        string                   `json:"@type,omitempty"`
	Links       []map[string]interface{} `json:"@links,omitempty"`
	State       string                   `json:"state,omitempty"`
	Environment string                   `json:"environment,omitempty"`
	CustomerID  string                   `json:"customer_id"`
	CreatedAt   time.Time                `json:"created_at,omitempty"`
	UpdatedAt   time.Time                `json:"updated_at,omitempty"`
	IsVisible   bool                     `json:"is_visible,omitempty"`
	Policies    PolicyResponse           `json:"policies,omitempty"`
	History     []map[string]interface{} `json:"history,omitempty"`
	User        User                     `json:"user,omitempty"`
}

// User -
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type GetPolicyParams struct {
	AccountNumber  string
	CustomerUserID string
	PortalTypeID   string
	PolicyID       int
}

func NewGetPolicyParams() *GetPolicyParams {
	return &GetPolicyParams{}
}

type AddPolicyParams struct {
	AccountNumber  string
	CustomerUserID string
	PortalTypeID   string
	PolicyAsString *string
	Policy         *Policy
}

func NewAddPolicyParams() *AddPolicyParams {
	return &AddPolicyParams{}
}

type SubmitDeployRequestParams struct {
	AccountNumber  string
	CustomerUserID string
	PortalTypeID   string
	DeployRequest  SubmitDeployRequest
}

func NewSubmitDeployRequestParams() *SubmitDeployRequestParams {
	return &SubmitDeployRequestParams{}
}
