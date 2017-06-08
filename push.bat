cls 
git status 
pause
git add .
git commit -m %*
git push
cls
@echo off
echo "Code pushed with comment =>" %*
