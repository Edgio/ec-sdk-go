# Onboarding: ec-go-sdk

## Purpose
This document is a collection of all onboarding-related information, tips & 
tricks, etc. for first-time SDK contributors.

## Learning Resources
The internet contains a plethora of learning resources for the Go language. 
Below are sources that we find useful:
- [Go Get Started Guide](https://go.dev/learn/)
    - Also lists additional learning resources.
- [Pluralsight (License Required)](https://www.pluralsight.com/)
    - We recommend the [Getting Started With Go](https://app.pluralsight.com/library/courses/getting-started-with-go/) course for an excellent guide to setting up your 
    local development environment.
- [Go by Example](https://gobyexample.com/)
    - Ideal for those who learn best by doing.
- [The Go Playground](https://go.dev/play/)
    - Sandbox for playing with Go code.
- [Go Docs](https://golang.org/doc/)
    - Landing page for many Go resources e.g. learning, concepts, etc.
- [Effective Go](https://go.dev/doc/effective_go)
    - Outlines community-accepted code style and best practices.
- [Our Go Guide](https://github.com/EdgeCast/ec-sdk-go/blob/main/Go.md)
    - Best practices and code style our team adheres to.

## Local Environment Setup
If you have a Pluralsight license, we recommend watching the Getting Started 
With Go course listed above in the Learning Resources.

Below is a breakdown of the setup:
1. Install Go - [download](https://go.dev/dl/)
2. Install Git - [download](https://git-scm.com/downloads)
3. Install Visual Studio Code - [download](https://code.visualstudio.com/download)
4. Set up Git access for your machine
    - Use a Git client like [Sourcetree](https://www.sourcetreeapp.com/) or
    [GitHub Desktop](https://desktop.github.com/)
    - [Create an SSH Key](https://docs.github.com/en/authentication/connecting-to-github-with-ssh) 
    for your GitHub account and add it to your local machine.
5. Install Visual Studio Code Extension
    - In VS Code, open the Extensions menu and search for "Go". The developer 
    should be "Go Team at Google".
6. Install any suggested Go tools as prompted by Visual Studio
7. Clone the [ec-sdk-go](https://github.com/EdgeCast/ec-sdk-go) repository.
8. Refer to the [README.md](README.md), [Contributing.md](Contributing.md), and 
[Architecture.md](Architecture.md) files.

## Development Workflow
Create a branch off of main and begin coding! 

### Default Environment Configuration
Please note that the SDK is configured to point to the EdgeCast production 
environment. This is fine if you own a test account. Developers employed at 
EdgeCast may wish to point to a different environment. You can do so globally by 
modifying the URLs in [config.go](edgecast/config.go) or during runtime by 
setting the appropriate URL fields in the `SDKConfig` struct returned by 
`edgecast.NewSDKConfig()`. 

For example:
```go
config := edgecast.NewSDKConfig()
apiURL, _ := url.Parse("https://api.vdms.io")
config.BaseAPIURL = *apiURL
```

### Testing
#### Unit Testing
Ensure that all unit tests pass before submitting a pull request. 

1. Open a terminal at the root of the repository.
2. Run `go test ./…` from the repository root. 

Please create or modify unit tests when modifying or adding to any of the code 
in `edgecast/internal`. When unit testing feature Services e.g. 
`waf.WafService`, use the `ecclient.MockAPIClient` mock struct to mock the API
functionality. Below is an example:

```go
func TestAddOrigin(t *testing.T) {
	config := edgecast.NewSDKConfig()
	originService, _ := New(config)

	originService.client = ecclient.MockAPIClient{
		// Mock data for the mock API client to return
		ResponseData: AddUpdateOriginOK{CustomerOriginID: 1},
	}

	// Add tests...
}
```

#### Regression Testing
Consider the scope of changes in your PR and whether it is necessary to run some 
or all of the example files located in the [example](example) folder as 
end-to-end tests to identify regressions.

### Release
Create a new release in GitHub with the appropriate vM.m.r semantic version e.g. 
v0.1.8 via the releases tab. Ensure to create a tag on the release screen using 
the same name. This process will be replaced with a GitHub action in the future.

In your release, be sure to include each section below. If there are no changes
for a section, omit it. Refer to existing 
[releases](https://github.com/EdgeCast/ec-sdk-go/releases).
- Breaking Changes
    - Alert consumers of the SDK of any changes that can break their code.
- New Features
- Bug Fixes and Enhancements
    - Enhancements can include performance improvements, code optimization, etc.