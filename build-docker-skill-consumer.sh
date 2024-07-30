docker build -t $SKILL_CONSUMER_REGISTRIES:$CI_COMMIT_SHORT_SHA -f  ./consumer/Dockerfile ./consumer
docker tag $SKILL_CONSUMER_REGISTRIES:$CI_COMMIT_SHORT_SHA $SKILL_CONSUMER_REGISTRIES:latest
docker login $REGISTRIES -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD
docker image push --all-tags $SKILL_CONSUMER_REGISTRIES