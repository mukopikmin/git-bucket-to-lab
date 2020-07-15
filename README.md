# Migration app from GitBucket to GitLab

Migrate repositories on GItBucket to projects on GitLab.

## Working environment

- GitBucket: 4.33.0
- GitLab: 12.10.6

## Features

- [x] Repository (Including code, commits, tags, branches)
- [ ] Release
- [x] Issue
- [x] Pull request
- [ ] Wiki

## Install

Build client application.

    npm install
    npm run build

Set environtment variable and start server.

    GITBUCKET_URL=http://gitbucket.example.com GITLAB_URL=http://gitlab.example.com go run main.go

### Run with docker

    docker run \
        -e GITBUCKET_URL=http://gitbucket.example.com \
        -e GITLAB_URL=http://gitlab.example.com \
        mukopikmin/git-bucket-to-lab
