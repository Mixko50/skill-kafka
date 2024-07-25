docker build -t $SKILL_API_REGISTRIES:$CI_COMMIT_SHORT_SHA -f  ./api/Dockerfile ./api
docker tag $SKILL_API_REGISTRIES:$CI_COMMIT_SHORT_SHA $SKILL_API_REGISTRIES:latest
docker login $REGISTRIES -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD
docker image push --all-tags $SKILL_API_REGISTRIES