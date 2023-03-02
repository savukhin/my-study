SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
GOLANG_DIR=.

cmd="protoc \
--go-grpc_out=$GOLANG_DIR --go_out=$GOLANG_DIR \
\
-I. \
*.proto "

if ! eval "$cmd" ; then
    echo "Fail generating grpc"
    exit
fi

echo "Success generating grpc"
# eval "$cmd"

# echo "${cmd}"
