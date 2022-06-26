### Branch Switcheroo

---

I had some shell functions that checked whether a git repo used `main` or `master`
and changed to those branches using `gcom` and `gpom` (git checkout main/master and 
git pull origin main/master depending on what the default branch name was).


**To use:**

`make install`

Then in your terminal:
* `gcom`: to check out main or master, depending on which this repo uses
* `gpom`: to pull main or master, depending on which this repo uses
