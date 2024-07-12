package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i ApiService -o ./mocks/ -s "_minimock.go"
