
all: dep build

dep:
	dep init

build:
	go build -o target/charityreports/bin .
	cp -r scripts target/charityreports/
	mkdir target/charityreports/out/
	mkdir target/charityreports/out/reports

k8s: build docker-build docker-push k8s-deploy

run:
	cd target/charityreports && ./bin

docker-build:
	docker build . -t charityreports -t fjmendes1994/charityreports:0.0.3

docker-push:
	docker push fjmendes1994/charityreports:0.0.3

docker-run:
	docker run -it --cpus="4" fjmendes1994/charityreports:0.0.3

k8s-deploy:
	envsubst < devstuff/k8s/deployment.yml | kubectl apply -f -

clean:
	rm -rf target
	rm -rf vendor
	rm Gopkg.lock
	rm Gopkg.toml

