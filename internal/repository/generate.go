package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i ChatServerRepository -o ./mocks/ -s "_minimock.go"
