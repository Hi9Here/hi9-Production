echo "coping in to dist bower_components and app"
rsync -a bower_components/ dist/bower_components/ && rsync -a app/ dist/
echo "copied"

