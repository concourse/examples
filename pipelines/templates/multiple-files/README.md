This pipeline has been split over multiple files. This is an example so the only resource used is the [mock resource](https://github.com/concourse/mock-resource).

Render the template using [ytt](https://github.com/vmware-tanzu/carvel-ytt/releases) and running the following command inside this directory:

```
ytt -f ./
```
