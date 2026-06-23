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
};

type Settings struct {
    Timeout      int    `yaml:"timeout"`;
    AlertWebhook string `yaml:"alert_webhook"`;
};

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
    webhook_url string;
};


func loadConfig(filePath string) AppConfig {
   
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

    yamlErr := yaml.Unmarshal(data, &config);
    if yamlErr != nil {
        log.Fatalf("Error parsing YAML: %v", yamlErr);
    }

    var sites []SiteConfig;

    for _, v := range config.Sites {
        siteConfig := SiteConfig {
            name: v.Name,
            url: v.URL,
        };

        sites = append(sites, siteConfig);
    }

    appConfig := AppConfig {
        sites: sites,
        timeout: config.Settings.Timeout,
        webhook_url: config.Settings.AlertWebhook,

    }

    return appConfig;

}

func main(){

    fmt.Println("Hello, World!");
    config := loadConfig("targets.yaml");
    fmt.Printf("Retrieved sites are: %+v\n", config.sites);
}

