docker build -t hellogo .
docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD 
docker tag hellogo $DOCKER_USERNAME/hellogo:$TRAVIS_BUILD_NUMBER
docker push $DOCKER_USERNAME/hellogo:$TRAVIS_BUILD_NUMBER
docker tag hellogo $DOCKER_USERNAME/hellogo:latest
docker push $DOCKER_USERNAME/hellogo:latest
