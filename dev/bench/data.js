window.BENCHMARK_DATA = {
  "lastUpdate": 1651082410611,
  "repoUrl": "https://github.com/vmware-tanzu/carvel-kapp-controller",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "name": "vmware-tanzu",
            "username": "vmware-tanzu"
          },
          "committer": {
            "name": "vmware-tanzu",
            "username": "vmware-tanzu"
          },
          "id": "6200ea5083ca1db200b904390e7f2d958c8321ed",
          "message": "pkgr benchmark test",
          "timestamp": "2022-04-06T17:27:49Z",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/pull/629/commits/6200ea5083ca1db200b904390e7f2d958c8321ed"
        },
        "date": 1650324813749,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 93995132319,
            "unit": "ns/op\t        63.23 DeleteSeconds\t        30.72 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 23716216909,
            "unit": "ns/op\t        15.48 DeleteSeconds\t         8.197 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 14601909907,
            "unit": "ns/op\t         9.418 DeleteSeconds\t         5.148 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "name": "vmware-tanzu",
            "username": "vmware-tanzu"
          },
          "committer": {
            "name": "vmware-tanzu",
            "username": "vmware-tanzu"
          },
          "id": "3281853ec58c6c619bdcd0fe5ff2fe5f87db2848",
          "message": "pkgr benchmark test",
          "timestamp": "2022-04-20T16:56:42Z",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/pull/629/commits/3281853ec58c6c619bdcd0fe5ff2fe5f87db2848"
        },
        "date": 1650479550694,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 95229929395,
            "unit": "ns/op\t        64.36 DeleteSeconds\t        30.83 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 23808474566,
            "unit": "ns/op\t        15.56 DeleteSeconds\t         8.203 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 14629206009,
            "unit": "ns/op\t         9.420 DeleteSeconds\t         5.170 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "86852107+joe-kimmel-vmw@users.noreply.github.com",
            "name": "Joe Kimmel",
            "username": "joe-kimmel-vmw"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9df3455df3ef0e2779d47171b1dfeeb53160d8cf",
          "message": "pkgr benchmark test (#629)\n\n* pkgr benchmark test\r\n\r\n* benchmark action: do it on pullrequests even though maybe we wouldn't really want to\r\n\r\n* benchmarks with right argument and version of go\r\n\r\n* benchmark tests get thresholds so on any given test run we can do an absolute time comparison instead of just relative timings between benchmark runs.\r\n\r\n* working on benchmark graphs on github pages\r\n\r\n* write to github pages conditionally on branch name develop only\r\n\r\n* collapse benchmark storage into one block\r\n\r\n* only store benchmark results on develop branch\r\n\r\n* fix the benchmark storage\r\n\r\n* add docs w link to benchmark tests to devmd\r\n\r\n* pin action to sha instead of tag and comment out conditional so we can test\r\n\r\n* restore conditionals",
          "timestamp": "2022-04-20T15:58:40-04:00",
          "tree_id": "3b0a2791411a3236324ef448db255f1a0ae4fd2f",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/9df3455df3ef0e2779d47171b1dfeeb53160d8cf"
        },
        "date": 1650485485829,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 94443823631,
            "unit": "ns/op\t        63.52 DeleteSeconds\t        30.88 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 23855612927,
            "unit": "ns/op\t        15.56 DeleteSeconds\t         8.249 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 14642181389,
            "unit": "ns/op\t         9.423 DeleteSeconds\t         5.172 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "cppforlife@gmail.com",
            "name": "Dmitriy Kalinin",
            "username": "cppforlife"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2935b002334bfcf9515b414641bb0a0dabb088b7",
          "message": "Merge pull request #554 from vmware-tanzu/kctrl-app-commands\n\n`kctrl app` commands",
          "timestamp": "2022-04-21T08:35:22-04:00",
          "tree_id": "2e7834644fe0de9da446a6c764463d5d527b52be",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/2935b002334bfcf9515b414641bb0a0dabb088b7"
        },
        "date": 1650545215285,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 94062479801,
            "unit": "ns/op\t        63.23 DeleteSeconds\t        30.80 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 23757471228,
            "unit": "ns/op\t        15.52 DeleteSeconds\t         8.201 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 14610827937,
            "unit": "ns/op\t         9.433 DeleteSeconds\t         5.141 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "86852107+joe-kimmel-vmw@users.noreply.github.com",
            "name": "Joe Kimmel",
            "username": "joe-kimmel-vmw"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "3f6d2325258ac8a15e3fbfe8ff6959b44b13a0dd",
          "message": "apiserver: custom QPS and Burst to allow high throughput of packages (#635)\n\n* apiserver: custom QPS and Burst to allow high throughput of packages\r\n\r\n* updating benchmark test upper bounds to reflect our speedy new reality\r\n\r\n* smallest numbers that have the same effect",
          "timestamp": "2022-04-21T10:22:06-04:00",
          "tree_id": "1f247d50501b2429ae43ebba8639373800197c51",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/3f6d2325258ac8a15e3fbfe8ff6959b44b13a0dd"
        },
        "date": 1650551538244,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36448166003,
            "unit": "ns/op\t        18.93 DeleteSeconds\t        17.49 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9521619928,
            "unit": "ns/op\t         5.280 DeleteSeconds\t         4.207 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6424466071,
            "unit": "ns/op\t         4.247 DeleteSeconds\t         2.142 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "cppforlife@gmail.com",
            "name": "Dmitriy Kalinin",
            "username": "cppforlife"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "426dc0517b32e623dee854b20eebdb8f43e02c81",
          "message": "remove dead code (#636)\n\nCo-authored-by: Dmitriy Kalinin <dkalinin@vmware.com>",
          "timestamp": "2022-04-21T13:49:36-04:00",
          "tree_id": "06aeb89d7a851ef802f07348c46cc9adbff03346",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/426dc0517b32e623dee854b20eebdb8f43e02c81"
        },
        "date": 1650563984950,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36454126329,
            "unit": "ns/op\t        18.94 DeleteSeconds\t        17.48 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9432302068,
            "unit": "ns/op\t         5.254 DeleteSeconds\t         4.142 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6423077774,
            "unit": "ns/op\t         4.255 DeleteSeconds\t         2.133 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "benm@vmware.com",
            "name": "Ben Moss",
            "username": "benmoss"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "600f2a3f155a5afdf93ac9f7daeeb66b62435518",
          "message": "Merge pull request #643 from vmware-tanzu/dependabot/github_actions/reviewdog/action-misspell-1.12",
          "timestamp": "2022-04-26T08:57:25-04:00",
          "tree_id": "3c4e29723a2dfb5f16df7266c74b5c915dc3cc26",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/600f2a3f155a5afdf93ac9f7daeeb66b62435518"
        },
        "date": 1650978486983,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36571937043,
            "unit": "ns/op\t        18.98 DeleteSeconds\t        17.54 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9551408963,
            "unit": "ns/op\t         5.321 DeleteSeconds\t         4.189 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6525658597,
            "unit": "ns/op\t         4.331 DeleteSeconds\t         2.151 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "benm@vmware.com",
            "name": "Ben Moss",
            "username": "benmoss"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "708b92b0f4ce2673f9d234004d89b5968273cdd7",
          "message": "Merge pull request #641 from vmware-tanzu/dependabot/github_actions/actions/setup-go-3\n\nBump actions/setup-go from 1 to 3",
          "timestamp": "2022-04-26T08:59:31-04:00",
          "tree_id": "8e35206cfed101ab0fdac5c89dc31d1e05230b49",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/708b92b0f4ce2673f9d234004d89b5968273cdd7"
        },
        "date": 1650978707044,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36934301053,
            "unit": "ns/op\t        19.26 DeleteSeconds\t        17.63 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9638468695,
            "unit": "ns/op\t         5.366 DeleteSeconds\t         4.227 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6527526181,
            "unit": "ns/op\t         4.313 DeleteSeconds\t         2.166 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "benm@vmware.com",
            "name": "Ben Moss",
            "username": "benmoss"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "4c0a29912178cd9337e83c1e8aaba55f20cc3550",
          "message": "Merge pull request #640 from vmware-tanzu/dependabot/github_actions/actions/cache-3\n\nBump actions/cache from 1 to 3",
          "timestamp": "2022-04-26T08:59:59-04:00",
          "tree_id": "eb05d875f77ccc80e0a60672d6958855dbf98716",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/4c0a29912178cd9337e83c1e8aaba55f20cc3550"
        },
        "date": 1650978759868,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36836204793,
            "unit": "ns/op\t        19.15 DeleteSeconds\t        17.65 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9621329566,
            "unit": "ns/op\t         5.362 DeleteSeconds\t         4.215 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6583535836,
            "unit": "ns/op\t         4.342 DeleteSeconds\t         2.146 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "benm@vmware.com",
            "name": "Ben Moss",
            "username": "benmoss"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "c5f32fc1202eaa6e2def9191c8399d82a4e4f238",
          "message": "Merge pull request #642 from vmware-tanzu/dependabot/github_actions/github/codeql-action-2",
          "timestamp": "2022-04-27T13:19:30-04:00",
          "tree_id": "0f755f9843056d19cc89e4727acfaa7a5c14d197",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/c5f32fc1202eaa6e2def9191c8399d82a4e4f238"
        },
        "date": 1651080576169,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36395157526,
            "unit": "ns/op\t        18.91 DeleteSeconds\t        17.45 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9492121825,
            "unit": "ns/op\t         5.296 DeleteSeconds\t         4.162 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6413889991,
            "unit": "ns/op\t         4.250 DeleteSeconds\t         2.122 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "benm@vmware.com",
            "name": "Ben Moss",
            "username": "benmoss"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "42274fbe2fac7c99477516c93044417d464c0840",
          "message": "Merge pull request #645 from vmware-tanzu/dependabot-stops-ignoring-patch-1",
          "timestamp": "2022-04-27T13:20:31-04:00",
          "tree_id": "ac764e65af343900ee067396ce497fd9db69218e",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/42274fbe2fac7c99477516c93044417d464c0840"
        },
        "date": 1651080737833,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36899294340,
            "unit": "ns/op\t        19.18 DeleteSeconds\t        17.66 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9643019379,
            "unit": "ns/op\t         5.410 DeleteSeconds\t         4.181 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6524449138,
            "unit": "ns/op\t         4.316 DeleteSeconds\t         2.163 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "distinct": true,
          "id": "12b66f37ab1d348001acd0253157296c3159c850",
          "message": "Bump k8s.io/apiserver from 0.22.4 to 0.22.9\n\nBumps [k8s.io/apiserver](https://github.com/kubernetes/apiserver) from 0.22.4 to 0.22.9.\n- [Release notes](https://github.com/kubernetes/apiserver/releases)\n- [Commits](https://github.com/kubernetes/apiserver/compare/v0.22.4...v0.22.9)\n\n---\nupdated-dependencies:\n- dependency-name: k8s.io/apiserver\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>",
          "timestamp": "2022-04-27T17:49:32Z",
          "tree_id": "9ce1d2629f8ae762b34111141524331585fc4fbb",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/12b66f37ab1d348001acd0253157296c3159c850"
        },
        "date": 1651082396719,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36647962144,
            "unit": "ns/op\t        19.05 DeleteSeconds\t        17.56 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9594400778,
            "unit": "ns/op\t         5.374 DeleteSeconds\t         4.177 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6531480074,
            "unit": "ns/op\t         4.351 DeleteSeconds\t         2.135 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "committer": {
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]",
            "username": "dependabot[bot]"
          },
          "distinct": true,
          "id": "3b02b556f3de6ef9fa794d37768321559f99a918",
          "message": "Bump github.com/stretchr/testify from 1.7.0 to 1.7.1\n\nBumps [github.com/stretchr/testify](https://github.com/stretchr/testify) from 1.7.0 to 1.7.1.\n- [Release notes](https://github.com/stretchr/testify/releases)\n- [Commits](https://github.com/stretchr/testify/compare/v1.7.0...v1.7.1)\n\n---\nupdated-dependencies:\n- dependency-name: github.com/stretchr/testify\n  dependency-type: direct:production\n  update-type: version-update:semver-patch\n...\n\nSigned-off-by: dependabot[bot] <support@github.com>",
          "timestamp": "2022-04-27T17:49:43Z",
          "tree_id": "7aa3cf4117765d721d29ece644bfdf9ed7abf61f",
          "url": "https://github.com/vmware-tanzu/carvel-kapp-controller/commit/3b02b556f3de6ef9fa794d37768321559f99a918"
        },
        "date": 1651082409903,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_pkgr_with_500_packages",
            "value": 36606529854,
            "unit": "ns/op\t        19.00 DeleteSeconds\t        17.55 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_100_packages",
            "value": 9577952833,
            "unit": "ns/op\t         5.336 DeleteSeconds\t         4.198 DeploySeconds",
            "extra": "1 times\n2 procs"
          },
          {
            "name": "Benchmark_pkgr_with_50_packages",
            "value": 6475583865,
            "unit": "ns/op\t         4.276 DeleteSeconds\t         2.157 DeploySeconds",
            "extra": "1 times\n2 procs"
          }
        ]
      }
    ]
  }
}