# Release go-slalom

If you are working on a cli or application that you want to make available to your team or otherwise then you will want
to create a github release and provide binaries. 

Or you may want to create a release for your service with binaries and use those when creating an image.

Below will show you how to use goreleaser to release a go application. 

### Use goreleaser

First install goreleaser

```bash
brew install goreleaser
```

goreleaser needs a github token to access your repository

```bash
export GITHUB_TOKEN=`YOUR_TOKEN`
```

It uses the latest tag applied to your repository. Create a tag if one does not exist

**Note**, the tag must adhere to [semantic versioning](https://goreleaser.com/semver). 
```bash
$ git tag -a v0.1.0 -m "First release"
$ git push origin v0.1.0
```

Now run goreleaser!
```bash
goreleaser
```

It will create a release in your repository with a changelog (containing commits for the release) and attached binaries

![release](screens/release.png)