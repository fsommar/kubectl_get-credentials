# kubectl get-credentials

This `kubectl` plugin effectively replaces the `gcloud container clusters get-credentials` command, allowing for more
choice while avoiding annoyances such as changing the current context when adding new contexts. It also uses its own
authentication mechanism for `kubectl`, removing the dependency on `gcloud` altogether.

_Why would I use this plugin?_

* Do you have a lot of GKE clusters in the same Google Cloud Platform project?
* Do you dislike the verbosity of the Kubernetes contexts generated by `gcloud`?
* Do you want to update your cluster details independent of your contexts?
* Do you switch between contexts often, and are encumbered by `gcloud` having to fetch new credentials each time?

If your answer is yes to either of those questions, `kubectl get-credentials` might be for you!

## Example usage

```sh
kubectl get-credentials gcp --project fredriksommar \
  --selector env=production,tier!=foo \
  --template '{{ .Cluster.Name }}' \
  --create-contexts
```

or equivalently

```sh
kubectl get-credentials gcp -p fredriksommar -lenv=production,tier!=foo -xt '{{ .Cluster.Name }}'
```

## Installation

Clone this repo and run `go install`. To verify that it's working as expected, run `kubectl get-credentials gcp --help`.