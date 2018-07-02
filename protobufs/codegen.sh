protowrap -I $(pwd)/client \
    --go_out=$(pwd)/../protocol \
    --proto_path $(pwd)/client \
    --print_structure \
    --only_specified_files \
    $(pwd)/client/*.proto

sed -i 's/^package.*/package protocol/' ../protocol/*.pb.go
