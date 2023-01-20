# **Contributing**

## **Permissions**

In GitHub you must change the repository settings: `settings -> actions[general] -> Allow GitHub Actions to create and approve pull requests reviews`

## **Tag Management**

While unsable (major version < 1.0.0) expect the following commit patterns:

- `BREAKING CHANGE` only bumps semver minor if version < 1.0.0
- `feat` commits bump semver patch instead of minor if version < 1.0.0

When major is stable (major version >= 1.0.0):

- `fix`: which represents bug fixes, and correlates to a SemVer patch.
- `feat`: which represents a new feature, and correlates to a SemVer minor.
- `feat!:`, `fix!:`, `refactor!:`, etc., which represent a breaking change (indicated by the !) and will result in a SemVer major.

Manual bump to 1.0.0:

```bash
git commit --allow-empty -m "chore: release 1.0.0" -m "release-as: 1.0.0"
```

## **Release Please**

[release-please](https://github.com/googleapis/release-please)
[release-please-action](https://github.com/google-github-actions/release-please-action)
[release-please-action-manifest](https://github.com/googleapis/release-please/blob/main/docs/manifest-releaser.md)
