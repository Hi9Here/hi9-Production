echo "coping in to dist bower_components and app"
 rsync -a app/ dist/ && rsync -a bower_components/ dist/bower_components/ && rsync -a gitRepos/ dist/bower_components/ 
echo "copied"

