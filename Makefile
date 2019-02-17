
all: dep build

dep:
	dep init

build:
	go build -o target/charityreports/bin .
	cp -r scripts target/charityreports/
	mkdir target/charityreports/out/
	mkdir target/charityreports/out/reports


run:
	cd target/charityreports && ./bin

docker-build:
	docker build . -t charityreports -t fjmendes1994/charityreports

docker-run:
	docker run -it --cpus="4" -v "$(pwd)"/reports:/charityreports/out/reports charityreports


clean:
	rm -rf target
	rm -rf vendor
	rm Gopkg.lock
	rm Gopkg.toml