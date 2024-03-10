# Contributing

## Prerequisites

1. [Install Go][go-install].
2. Download the sources and switch the working directory:

    ```bash
    go get github.com/rezakhademix/validator
    ```

## Submitting a Pull Request

A typical workflow is:

1. [Fork the repository.][fork] [This tip maybe also helpful.][go-fork-tip]
2. [Create a topic branch.][branch]
    - always use dashes to separate branch name words just like kebab-case naming convention. e.g: feature-kebab-case-example [ReadMore About kebab-case][kebab-case]
3. Add tests for your change.
4. Run `go test` to make sure everything is ok.
5. Implement the change and ensure the steps from the previous step pass.
6. [Add, commit and push your changes.][git-help]
7. [Submit a pull request.][pull-req]

[go-install]: https://golang.org/doc/install
[go-fork-tip]: http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html
[fork]: https://help.github.com/articles/fork-a-repo
[branch]: http://learn.github.com/p/branching.html
[git-help]: https://guides.github.com
[pull-req]: https://help.github.com/articles/using-pull-requests
[kebab-case]: https://developer.mozilla.org/en-US/docs/Glossary/Kebab_case
