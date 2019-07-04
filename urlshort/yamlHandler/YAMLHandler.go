package yamlHandler

import (
    "net/http"
    "os"

    "gopkg.in/yaml.v2"
)

type YAMLHandler struct {
    FilePath        string
    FallbackHandler http.Handler
}

type Mapping struct {
    Path     string `yaml:"path"`
    Redirect string `yaml:"redirect"`
}

func ReadYAMLData(filePath string) (res map[string]string) {
    reader, err := os.Open(filePath)
    CheckErr(err)
    decoder := yaml.NewDecoder(reader)
    var mp []Mapping
    err = decoder.Decode(&mp)
    CheckErr(err)
    res = make(map[string]string)
    for _, v := range mp {
        res[v.Path] = v.Redirect
    }
    return
}

func (y YAMLHandler) MakeHandler() *http.ServeMux {
    yamlData := ReadYAMLData(y.FilePath)
    m := MapHandler{yamlData, y.FallbackHandler}
    return m.MakeHandler()
    return nil
}
