TAG := "hiro-latest"

#AWS_ACCOUNT_ID := "282045701300"
AZ_REPO_LINK := "bzinfdevqacr.azurecr.io"
#AWS_REGION := "ap-south-1"
TARGET_REPOSITORY := "user-svc"

docker-build:
	docker build -f Dockerfile --tag ${TARGET_REPOSITORY}:${TAG} .
push:
	az acr login --name  bzinfdevqacr
	docker tag ${TARGET_REPOSITORY}:${TAG} ${AZ_REPO_LINK}/${TARGET_REPOSITORY}:${TAG}
	docker push ${AZ_REPO_LINK}/${TARGET_REPOSITORY}:${TAG}
images-push: docker-build push

deploy:
	make images-push
	ssh -i "C:/Users/Hiro/.ssh/github.com-hiroBzinga"  nashaath.mohamed@20.219.153.186
	kubectl set image deployment/user-svc -n bzinga user-svc=hiro --record
	kubectl set image deployment/user-svc -n bzinga user-svc=hiro-latest --record
	kubectl rollout restart deployment/user-svc -n bzinga