# Contributing guidelines

## Sending patches

You can send patches into this project via:

* **GitLab**: <https://mau.dev/andreijiroh-dev/golang-pkgs/-/merge_requests>
* **GitHub**: <https://github.com/andreijiroh-dev/golang-pkgs/pulls>
* **Mailing list (via `git send-email`)**: `~ajhalili2006/public-inbox@lists.sr.ht` (preferred) or `ajhalili2006@googlegroups.com`

### Setting up `git-email`

See <https://git-send-email.io/> for instructions on how to install `git-email` package and setting up email authenication.

After that:

```bash
# send to mailing list (using the lists.sr.ht address is recommended)
git config sendemail.to "~ajhalili2006/public-inbox@lists.sr.ht"
# alternative in case of downtime, but not recommended
git config sendemail.to "ajhalili2006@googlegroups.com"

# add subject prefix for project
git config format.subjectPrefix "PATCH golang-pkgs"

# sign-off your patches per https://dco.andreijiroh.xyz
git config format.signOff yes

# use --annotate by default
git config --global sendemail.annotate yes
```

Once fully configured, just `git send-email v<versionNo> HEAD^` and wait for a reply from ~ajhalili2006 for feedback.
