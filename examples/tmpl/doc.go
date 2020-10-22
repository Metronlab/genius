package tmpl

//go:generate genius tmpl -data=types.tmpldata.yml -v authors=nyhu,contributors core.gen.go.tmpl
//go:generate genius tmpl -data=types.tmpldata.yml -v authors=nyhu,contributors types.gen.json.tmpl
//go:generate bash -c "echo dummy | genius tmpl -e txt -p dummy. stdin.gen.txt.tmpl"
