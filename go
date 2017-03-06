#!/bin/bash
echo "Tried on $(date)" >> log
git add -u
git add -A
git commit --amend -m "Test"
git push --force origin test-branch-for-test-pr

