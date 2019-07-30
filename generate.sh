gen_dir()
{
    for file in *.proto
    do
        protoc -I=$1 --go_out=$1 $1/$file
    done
}

for dir in "$@"
do
    gen_dir $dir
done
