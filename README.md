## How to use this repo

1. Place the chart into `stable` folder
2. Add repo to your Helm repositories as:
```
helm repo add matic https://matic-insurance.github.io/helm-charts
```
3. Work with the charts
```
helm search matic/nginx
NAME        	CHART VERSION	APP VERSION	DESCRIPTION
matic/nginx	1.0.0        	1.0        	A Nginx Helm chart for Kubernetes
```

## Use hook for local dry-run tests

```bash
git config core.hooksPath .hooks
```

## Add tests and golden file

We use as an example application-mesh chart and redirects component.

1. Add values helm-charts/stable/application-mesh/tests/golden/components/redirects.values.yaml
2. Add new component to helm-charts/stable/application-mesh/tests/golden/golden_test.go
```
		{
			GoldenFileName: "components/redirects.golden.yaml",
			ValuesFiles:    []string{"components/redirects.values.yaml"},
			Templates:      []string{"templates/redirects.yaml"},
		},
```
3. Run
```
cd helm-charts/stable/application-mesh
go test ./... -update-golden
```
This command will execute your tests and update the golden files to match the current output.
4. Run
```
go test ./...
```
5. Commit changes
