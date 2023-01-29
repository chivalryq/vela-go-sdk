# KubeVela Go SDK [WIP]

This is a Go SDK for KubeVela. 

## Features:

- 🔧Application manipulating
  - [x] Add Components/Traits/Workflow Steps/Policies
  - [x] Setting Workflow Mode
  - [x] Convert to K8s Application Object
  - [ ] Convert from K8s Application Object
  - [ ] Referring to external Workflow object.
- 🔍Application client
  - [x] Create/Delete/Patch/Update Application
  - [ ] List/Get Application
 
## Example

See [example](./example) for more details.

## Future Work

There is some proper features for this SDK and possible to be added in the future. If you are interested in any of them, please feel free to contact us.

- Part of vela CLI functions
  - Application logs/exec/port-forward
  - Application resource in tree structure
  - VelaQL
  - ...
- Standalone workflow functions
  - CRUD of workflow

## Known issues

There are some known issues in this SDK, please be aware of them before using it. 

1. `labels` and `annotations` trait is not working as expected.
2. `notification` workflow-step's parameter is not exactly the same as the one in KubeVela.