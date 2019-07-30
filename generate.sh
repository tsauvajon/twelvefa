gen_dir()
{
    for file in $1/*.proto
    do
        protoc -I=$1 --go_out=plugins=grpc:$1 $file
    done
}

for dir in "$@"
do
    gen_dir $dir
done
