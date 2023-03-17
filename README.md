# cli 

Template CLI project including: 

* config - defined by Go struct
* data - stored in a an embedded sqlite DB

> Both customizable and stored in the user home directory.

The repo also includes pipelines to create distributable binaries in multiple OS/Architecture combinations with SBOM and SLSA provenance. 

> See [releases](https://github.com/mchmarny/cli/releases/latest) for example.

## Repo Usage 

Use this template to create a new repo (click the green button and follow the wizard)

![](images/template.png)

When done, clone your new repo locally, and navigate into it

```shell
git clone git@github.com:$GIT_HUB_USERNAME/$REPO_NAME.git
cd $REPO_NAME
```

Initialize your new repo. This will update all the references to your newly clone GitHub repository.

```shell
tools/init
```

When completed, commit and push the updates to your repository: 

```shell
git add --all
git commit -m 'repo init'
git push --all
```

> The above push will trigger the `on-push` flow. You can navigate to the `/actions` in your repo to see the status of that pipeline. 

![](images/push.png)

### Trigger release pipeline

The canonical version of the entire repo is stored in [.version](.version) file. Feel free to edit it (by default: `v0.0.1`). When done, trigger the release pipeline:

> If you did edit the version, make sure to commit and push that change to the repo first. You can also use `make tag` to automate the entire process.

```shell
export VERSION=$(cat .version)
git tag -s -m "initial release" $VERSION
git push origin $VERSION
```

### Monitor the pipeline 

Navigate to `/actions` in your repo to see the status of that release pipeline. Wait until all steps (aka jobs) have completed (green). 

> If any steps fail, click on them to see the cause. Fix it, commit/push changes to the repo, and tag a new release to re-trigger the pipeline again.

## Disclaimer

This is my personal project and it does not represent my employer. While I do my best to ensure that everything works, I take no responsibility for issues caused by this code.
