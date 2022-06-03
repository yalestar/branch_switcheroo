### Branch Switcheroo

---

I had some shell functions that checked whether a git repo used `main` or `master`
and changed to those branches using `gcom` and `gpom` (git checkout main/master and 
git pull origin main/master depending on what the default branch name was)

This seemed like a good thing to rewrite in Go, since after all, *it's a systems language*.

**To use:**

`make install`