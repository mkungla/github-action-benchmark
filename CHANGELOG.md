<a name="v1.6.1"></a>
# [v1.6.1](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.6.1) - 07 Dec 2019

- **Fix:** Time unit of mean time in `pytest` benchmark results were always `sec`. Now time units are converted to `msec`, `usec` and `nsec` if necessary
- **Fix:** Detecting rejection by remote on `git push` was not sufficient
- **Improve:** Add a small link at right bottom of dashboard page to show this action provided the page
- **Improve:** Showed at least 1 significant digit for threshold float values like `2.0`
- **Improve:** Updated dependencies


[Changes][v1.6.1]


<a name="v1.6.0"></a>
# [v1.6.0](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.6.0) - 04 Dec 2019

- **New:** `fail-threshold` input was added. Format is the same as `alert-threshold`, but you can give different thresholds to sending a commit comment and making the workflow fail by giving different value to `fail-threshold` from `alert-threshold`. This value is optional. If omitted, `fail-threshold` value is the same as `alert-threshold`
- **Improve:** Retry logic was improved on `git push` failed due to remote branch updates after `git pull`. Now this action retries entire process to update `gh-pages` branch when the remote rejected automatic `git push`. Previously this action tried to rebase the local onto the remote but it sometimes failed due to conflicts

[Changes][v1.6.0]


<a name="v1.5.0"></a>
# [v1.5.0](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.5.0) - 30 Nov 2019

- **New:** Added `max-items-in-chart` input was added to limit the number of data points in a graph chart.
- **New:** Supported [Google C++ Benchmark Framework](https://github.com/google/benchmark) for C++ projects. Please check [the example project](https://github.com/rhysd/github-action-benchmark/tree/master/examples/cpp) and [the example workflow](https://github.com/rhysd/github-action-benchmark/blob/master/.github/workflows/cpp.yml) to know the setup
- **Fix:** Fix the order of graphs in the default `index.html`. To apply this fix, please remove `index.html` in your GitHub Pages branch and run your benchmark workflow again
- **Improve:** Use the actions marketplace URL for the link to this action in commit comment
- **Improve:** Updated dependencies
- **Dev:** Added Many tests for checking the updates on a new benchmark result
- **Dev:** Changed directory structure. Sources are now put in `src/` directory

[Changes][v1.5.0]


<a name="v1.4.0"></a>
# [v1.4.0](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.4.0) - 23 Nov 2019

- **New:** `external-data-json-path` input was added to support to put benchmark data externally rather than Git branch
  - By using this input and [actions/cache](https://github.com/actions/cache), you no longer need to use Git branch for this action if you only want performance alerts. Benchmark data is stored as workflow cache.
  - By this input, minimal setup for this action is much easier. Please read ['How to use' section](https://github.com/rhysd/github-action-benchmark#minimal-setup) in README.md.

[Changes][v1.4.0]


<a name="v1.3.2"></a>
# [v1.3.2](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.3.2) - 23 Nov 2019

- **Improve:** Styles in alert commit comment were improved
- **Fix:** When benchmark name (with `name` input) contained spaces, URL for the workflow which detected performance regression was broken

[Changes][v1.3.2]


<a name="v1.3.1"></a>
# [v1.3.1](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.3.1) - 21 Nov 2019

- **Fix:** `git push` sometimes failed in the situation where `prepush` hook is set and runs unexpectedly. Now `git push` is run with `--no-verify` for pushing auto generated commit to remote.

[Changes][v1.3.1]


<a name="v1.3.0"></a>
# [v1.3.0](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.3.0) - 21 Nov 2019

- **New:** Alert feature was added :tada:
  - With this feature enabled, you can get alert commit comment or make workflow fail when possible performance regression is detected [like this](https://github.com/rhysd/github-action-benchmark/commit/077dde1c236baba9244caad4d9e82ea8399dae20#commitcomment-36047186)
  - `comment-on-alert` input was added to enable commit comment on alert. `github-token` input is necessary as well to use GitHub API. Unlike deploying GitHub Pages, `secrets.GITHUB_TOKEN` is sufficient for this purpose (if you don't use GitHub Pages). The input is set to `false` by default.
  - `fail-on-alert` input was added to mark running workflow fail on alert. The input is set to `false` by default.
  - `alert-threshold` input was added to specify the threshold to check alerts. When current result gets worse than previous exceeding the threshold. Value is ratio such as `"200%"`. For example, when benchmark gets result 230 ns/iter and previous one was 100ns/iter, it means 230% worse and an alert will happen.
  - Please read [documentation](https://github.com/rhysd/github-action-benchmark#use-this-action-with-alert-commit-comment) for setup
- **New:** `alert-comment-cc-users` input was added to specify users mentioned in an alert commit comment so that they can easily notice it via GitHub notification
- **New:** `skip-fetch-gh-pages` input was added to skip `git pull` which is automatically executed on public repo or when you set `github-token` on private repo.
- **Improve:** E2E checks on CI were added
- **Improve:** Updated dependencies

[Changes][v1.3.0]


<a name="v1.2.0"></a>
# [v1.2.0](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.2.0) - 17 Nov 2019

- **New:** Support [pytest-benchmark](https://pypi.org/project/pytest-benchmark/) for Python projects which use pytest
  - Benchmark value is how long one iteration takes (seconds/iter)
- **Improve:** Show more extra data in tooltip which are specific to tools
  - Go
    - Iterations
    - Number of CPUs used
  - Benchmark.js
    - Number of samples
  - pytest-benchmark
    - Mean time
    - Number of rounds

For reflecting the extra data improvement, please refresh your `index.html`. Remove current `index.html` in GitHub Pages branch and push the change to remote, then re-run your benchmark workflow.

[Changes][v1.2.0]


<a name="v1.1.4"></a>
# [v1.1.4](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.1.4) - 16 Nov 2019

- **Improve:** Title styles in default `index.html` which is generated when no `index.html` is in your GitHub Pages branch. If you want to update your `index.html` to the latest, please remove it and push to remote at first then re-run your workflow which will invoke github-action-benchmark
- **Improve:** More metadata in `action.yml`. Now icon and its color are set.

[Changes][v1.1.4]


<a name="v1.1.3"></a>
# [v1.1.3](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.1.3) - 16 Nov 2019

- **Fix:** Retry failed when no Git user config is provided. Ensure to give bot user info to each `git` command invocations

[Changes][v1.1.3]


<a name="v1.1.2"></a>
# [v1.1.2](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.1.2) - 16 Nov 2019

- **Improve:** Added retry for `git push`. When remote GitHub Pages branch is updated after the current workflow had fetched the branch, `git push` will fail because the remote branch is not up-to-date. In the case this action will try to rebase onto the latest remote by `git pull --rebase` and `git push` again. This is useful when your multiple workflows may be trying to push GitHub Pages branch at the same timing. `auto-push` input must be set to `true` for this.
- **Fix:** Description for `auto-push` was missing in `action.yml`

[Changes][v1.1.2]


<a name="v1.1.1"></a>
# [v1.1.1](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.1.1) - 14 Nov 2019

- **Improve:** More strict check for `auto-push` input. Now the value must be one of `true`, `false` (default value is `false`)

[Changes][v1.1.1]


<a name="v1.1.0"></a>
# [v1.1.0](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.1.0) - 14 Nov 2019

- **New:** Added `auto-push` input
  - If this value is set to `true`, this action pushes GitHub Pages branch to remote automatically. You no longer need to push the branch by yourself.
  - Below `github-token` input must be set for this
  - This input is optional. You can still push the branch by yourself if you want
  - Please read [documentation](https://github.com/rhysd/github-action-benchmark#how-to-use) for more details
- **New:** Added `github-token` input
  - For doing some operations which requires GitHub API token, this input is necessary
    - pull from remote branch when your repository is private
    - push to remote branch
    - deploy and trigger GitHub Pages build
  - This input is optional. When you do none of above operations, this input is not necessary
- `README.md` was updated to avoid [the issue on public repository](https://github.community/t5/GitHub-Actions/Github-action-not-triggering-gh-pages-upon-push/td-p/26869) (#1)

e.g.

```yaml
- name: Store benchmark result
  uses: rhysd/github-action-benchmark@v1
  with:
    name: My Project Go Benchmark
    tool: 'go'
    output-file-path: output.txt
    github-token: ${{ secrets.PERSONAL_GITHUB_TOKEN }}
    auto-push: true
```

Note that you need to make a personal access token for deploying GitHub Pages from GitHub Action workflow. Please read `RADME.md` for more details.

[Changes][v1.1.0]


<a name="v1.0.2"></a>
# [v1.0.2](https://github.com/rhysd/github-action-benchmark/releases/tag/v1.0.2) - 10 Nov 2019

First release :tada:

Please read documentation for getting started:

https://github.com/rhysd/github-action-benchmark#readme

[Changes][v1.0.2]


[v1.6.1]: https://github.com/rhysd/github-action-benchmark/compare/v1.6.0...v1.6.1
[v1.6.0]: https://github.com/rhysd/github-action-benchmark/compare/v1.5.0...v1.6.0
[v1.5.0]: https://github.com/rhysd/github-action-benchmark/compare/v1.4.0...v1.5.0
[v1.4.0]: https://github.com/rhysd/github-action-benchmark/compare/v1.3.2...v1.4.0
[v1.3.2]: https://github.com/rhysd/github-action-benchmark/compare/v1.3.1...v1.3.2
[v1.3.1]: https://github.com/rhysd/github-action-benchmark/compare/v1.3.0...v1.3.1
[v1.3.0]: https://github.com/rhysd/github-action-benchmark/compare/v1.2.0...v1.3.0
[v1.2.0]: https://github.com/rhysd/github-action-benchmark/compare/v1.1.4...v1.2.0
[v1.1.4]: https://github.com/rhysd/github-action-benchmark/compare/v1.1.3...v1.1.4
[v1.1.3]: https://github.com/rhysd/github-action-benchmark/compare/v1.1.2...v1.1.3
[v1.1.2]: https://github.com/rhysd/github-action-benchmark/compare/v1.1.1...v1.1.2
[v1.1.1]: https://github.com/rhysd/github-action-benchmark/compare/v1.1.0...v1.1.1
[v1.1.0]: https://github.com/rhysd/github-action-benchmark/compare/v1.0.2...v1.1.0
[v1.0.2]: https://github.com/rhysd/github-action-benchmark/tree/v1.0.2

 <!-- Generated by changelog-from-release -->