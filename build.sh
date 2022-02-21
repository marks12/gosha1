#!/bin/bash

cd front
npm i

npm run build

cd ../

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

    value=`cat $pathname  | base64`

    files=( "${files[@]}" "$pathname" )
    content=( "${content[@]}" "$value" )

}

DOWNLOADING_DIR=./front/dist

walk_dir "$DOWNLOADING_DIR"

cat > ./webapp/webapp/html.go <<EOL
package webapp

type HtmlFile struct {
	Path string
	Content string
}

func GetFiles() (Files []HtmlFile) {

EOL



for i in "${!files[@]}"; do

cat >> ./webapp/webapp/html.go << EOL
    Files =  append(Files, HtmlFile{Path: "${files[i]}", Content: \`${content[i]}\`})
EOL

done

cat >> ./webapp/webapp/html.go <<EOL

    return Files
}
EOL

go build -o ./gosha -v ./main.go

env GOOS=darwin GOARCH=amd64 go build -o ./gosha-mac -v ./main-mac.go
