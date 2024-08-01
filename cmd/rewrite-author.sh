#!/bin/bash

git filter-repo --commit-callback '
if commit.committer_email == b"liuxiaowen@baimaohui.net":
    commit.committer_name = b"Mitchell Ryrie"
    commit.committer_email = b"liuxiaowen66621@gmail.com"
if commit.author_email == b"liuxiaowen@baimaohui.net":
    commit.author_name = b"Mitchell Ryrie"
    commit.author_email = b"liuxiaowen66621@gmail.com"
'
