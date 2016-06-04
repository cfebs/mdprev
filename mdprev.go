package main

import (
    "bytes"
    "fmt"
    "net/http"
    "io"
    "path/filepath"
    "io/ioutil"
    "html/template"
    "github.com/shurcooL/github_flavored_markdown"
    "flag"
    "log"
)

type Config struct {
    BasePath string
}

type TemplateData struct {
    Content template.HTML
}

type FileView struct {
    Markdown template.HTML
    Filepath string
}

type IndexView struct {
    BasePath string
    FilePaths []string
}

func WriteTemplate(w io.Writer, tpldata *TemplateData) {
    t, _ := template.ParseFiles("template.html")
    t.Execute(w, tpldata)
}

func MakeCfgHandler(fn func(w http.ResponseWriter, r *http.Request, cfg *Config), cfg *Config) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        fn(w, r, cfg)
    }
}

func RootHandler(w http.ResponseWriter, r *http.Request, cfg *Config) {
    filenames, err := filepath.Glob(cfg.BasePath + "/**.md")

    if err != nil {
        fmt.Fprintf(w, "File path %s could not be rendered", err)
        return
    }

    for i, name := range filenames {
        filenames[i] = name[len(cfg.BasePath + "/"):]
    }

    indexview := new(IndexView)
    indexview.FilePaths = filenames
    indexview.BasePath = cfg.BasePath

    t, _ := template.ParseFiles("index.html")
    var view bytes.Buffer
    t.Execute(&view, indexview);
    WriteTemplate(w, &TemplateData{Content: template.HTML(view.String())})
}

func FileHandler(w http.ResponseWriter, r *http.Request, cfg *Config) {
    // @todo test relative paths for security
    file_part := r.URL.Path[len("/file/"):]
    file_path := cfg.BasePath + "/" + file_part
    content, err := ioutil.ReadFile(file_path)

    if err != nil {
        fmt.Fprintf(w, "File path %s could not be rendered", file_path)
        return
    }

    output := github_flavored_markdown.Markdown(content)
    fileview := new(FileView)
    fileview.Markdown = template.HTML(output);
    fileview.Filepath = file_part

    // render view
    // write template
    t, _ := template.ParseFiles("file.html")
    var view bytes.Buffer
    t.Execute(&view, fileview);
    WriteTemplate(w, &TemplateData{Content: template.HTML(view.String())})
}

func main() {
    cfg := new(Config)

    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        log.Fatal("First arg must be path")
        return
    }

    basepath, err := filepath.Abs(args[0])

    if err != nil {
        log.Fatal("Abs path fail")
    }

    cfg.BasePath = basepath

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/",  http.StripPrefix("/static/", fs))

    http.HandleFunc("/", MakeCfgHandler(RootHandler, cfg))
    http.HandleFunc("/file/", MakeCfgHandler(FileHandler, cfg))
    fmt.Println("Listening on 8080")
    http.ListenAndServe(":8080", nil);
}
