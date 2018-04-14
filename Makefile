
install:
	go get -v github.com/Masterminds/glide
	cd ${GOPATH}/src/github.com/Masterminds/glide && git checkout tags/v0.13.1 && go install && cd -
	glide install

test: 
	./ci.sh
