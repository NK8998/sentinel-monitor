package main;

import ( 
    "fmt"
    "os"
    "io"  
    "log"
    "gopkg.in/yaml.v3"
);

type SiteField struct{ 
     Name string `yaml:"name"`;
     URL  string `yaml:"url"`;
}

type Settings struct {
    Timeout      int    `yaml:"timeout"`;
    AlertWebhook string `yaml:"alert_webhook"`;
} 

type Config struct {
    Settings Settings `yaml:"settings"`;
    Sites []SiteField `yaml:"sites"`;
};

type SiteConfig struct {
    name  string;
    url string;
};

type AppConfig struct {
    sites []SiteConfig;
    timeout int;
    webhook_url *string;
};


func loadConfig(filePath string){
   
    var config Config;

    file, err := os.Open(filePath);
    if err != nil {
        log.Fatalf("Failed to open file: %s", err);
    }

    defer file.Close();
    
    data, err := io.ReadAll(file);
    if err != nil {
        log.Fatal(err);
    }

//    for index, value := range data {
//      fmt.Printf("Index: %d, Byte: %d, Char: %c\n", index, value, value)
//}

    yamlErr := yaml.Unmarshal(data, &config);
    if yamlErr != nil {
        log.Fatalf("Error parsing YAML: %v", yamlErr);
    }

    fmt.Printf("Retrieved sites are: %+v\n", config.Sites);

}

func main(){

    fmt.Println("Hello, World!");
    loadConfig("targets.yaml");
}

