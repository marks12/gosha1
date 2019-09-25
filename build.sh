#!/bin/bash

declare -a files
declare -a content

walk_dir () {

    for pathname in "$1"/*; do
        if [ -d "$pathname" ]; then
            walk_dir "$pathname"
        else
            add_file $pathname
        fi
    done
}

add_file () {
    shopt -s nullglob dotglob

#    printf '%s\n' adding "$pathname"

    value=`cat $pathname  | base64`

    files=( "${files[@]}" "$pathname" )
    content=( "${content[@]}" "$value" )

}

DOWNLOADING_DIR=./front/dist

walk_dir "$DOWNLOADING_DIR"


cat > ./webapp/html.go <<EOL
package webapp

type HtmlFile struct {
	Path string
	Content string
}

func GetFiles() (Files []HtmlFile) {

EOL



for i in "${!files[@]}"; do

cat >> ./webapp/html.go << EOL
    Files =  append(Files, HtmlFile{Path: "${files[i]}", Content: \`${content[i]}\`})
EOL

done

cat >> ./webapp/html.go <<EOL

    return Files
}
EOL


env GOOS=darwin GOARCH=amd64 go build -o ./gosha-mac -v ./main.go
go build -v

env GOOS=windows GOARCH=386 go build -o gosha.exe -v ./main.go
go build -v

