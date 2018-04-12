govm() {
    export GOVMOS="darwin"
    export GOVMARCH="amd64"
    if [[ $1 == "use" ]]; then
        shift;
        d=`govm path $@`
        if [ ${d:0:5} == "PATH:" ]; then
            gopath=${d:5}
            PATH=$gopath/bin:$PATH
            export GOTOOLDIR="$gopath/pkg/tool/$GOVMOS_$GOVMARCH"
            export GOROOT=$gopath
            echo "Using version $1"
        else
            echo "could not find version"
        fi
    else
        command govm "$@"
    fi
}