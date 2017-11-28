
#   instead of the $GOPATH directly. For normal projects this can be dropped.
vendor/k8s.io/code-generator/generate-internal-groups.sh all \
github.com/enablecloud/kulbe/client/application \ github.com/enablecloud/kulbe/apis/cr github.com/enablecloud/kulbe/apis/cr \
application:v1
vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/enablecloud/kulbe/client/application \ github.com/enablecloud/kulbe/apis/cr \
application:v1
#   instead of the $GOPATH directly. For normal projects this can be dropped.
vendor/k8s.io/code-generator/generate-internal-groups.sh all \
github.com/enablecloud/kulbe/client/component \ github.com/enablecloud/kulbe/apis/cr github.com/enablecloud/kulbe/apis/cr \
component:v1
vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/enablecloud/kulbe/client/component \ github.com/enablecloud/kulbe/apis/cr \
component:v1