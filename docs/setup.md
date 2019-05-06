# Setup Guide 


### Install go-slalom 
Install go-slalom CLI:

```bash
brew tap tredfield/homebrew-tap
brew install tredfield/tap/go-slalom

# verify, should see cli help 
go-slalom
```

Create a new github repository called `demo-app`. Go to https://github.com/new and use `dempo-app` as repository name.

Clone your private repository (preferable in your `$GOPATH`) and initialize podinfo.

```bash
# if GOPATH not set, use $HOME/go
cd $GOPATH/src

# $GITHUB_USER should be your github user
mkdir -p github.com/$GITHUB_USER
cd github.com/$GITHUB_USER
git clone https://github.com/$GITHUB_USER/demo-app
cd demo-app

go-slalom code init demo-app --git-user=stefanprodan --version=master
```

The above command does the following:
* downloads go-slalom source code from GitHub 
* replaces golang imports with your git username and project name
* creates a Dockerfile and Makefile customized for GitHub actions
* creates the main workflow for GitHub actions
* commits and pushes the code to GitHub
