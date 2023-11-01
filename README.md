# gh-alfredworkflow

## Why this repo?

有很多类似功能的repo，比如

- [gharlan/alfred-github-workflow: GitHub Workflow for Alfred 4](https://github.com/gharlan/alfred-github-workflow)
- [lox/alfred-github-jump: An alfred plugin to quickly jump to a github repository page](https://github.com/lox/alfred-github-jump)
- [edgarjs/alfred-github-repos: Alfred workflow to easily open Github repositories](https://github.com/edgarjs/alfred-github-repos)
- [giovannicoppola/alfred-hubHub: A hub for your GitHub repositories](https://github.com/giovannicoppola/alfred-hubHub)
- ...

但是，要不就是功能太多，要不就是功能太少，用起来都不算趁手，所以自己写一个。

所以，这是一个自用服务。如果你也喜欢，可以在 relsee 页面下载并使用。

如果有问题，请发 issues 告诉我。

## Features

基本上参考了 [gharlan/alfred-github-workflow: GitHub Workflow for Alfred 4](https://github.com/gharlan/alfred-github-workflow) 的功能，删除和修改了一些不需要的

比如

- 移除了 @user 命令，因为这个功能实际上已经被repo覆盖了
- 移除了 OAuth 登录，直接填写 token，解决用户隐私问题
- 增加了 `gh my topic` 命令，可以直接列出 starred 的 topics，方便查看（对我来说，这是个高频操作）
- 把搜索 repo 命令分成了专门用来搜索自己repo的 `gh repo <repo>` 和 在 github 搜索repo的 `gh repos <repo>`

### my 命令

```shell

gh my dashboard
gh my notification
gh my profile
gh my issues
gh my pull
gh my new
gh my setting
gh my gist
gh my star
# list all my starred topics, Allows you to reach the topics you want to view faster
gh my topic

```

### repo 命令

```shell

# search my repo, print LANGUAGE & repo describe & stars
gh repo <repo>

# directly search repo
gh repos <repo>

```

### actions 命令

```shell

# 更新 workflow
gh actions update

```
