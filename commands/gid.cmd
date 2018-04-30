#######################################################
## How to modify a specified commit in git?

# From https://stackoverflow.com/questions/8824971/how-to-amend-older-git-commit
# and https://stackoverflow.com/questions/1186535/how-to-modify-a-specified-commit-in-git

git rebase -i HEAD^^^
# Now mark the ones you want to amend with edit or e (replace pick). Now save and exit.

# Now make your changes, then
git add -A
git commit --amend --no-edit
git rebase --continue

#######################################################
##
