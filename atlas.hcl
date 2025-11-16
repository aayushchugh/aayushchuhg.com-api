data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/models",
    "--dialect", "postgres",
  ]
}

env "local" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/ayushchugh.com?search_path=public"
  url = getenv("DATABASE_URL")  # Changed from env() to getenv()
  
  migration {
    dir = "file://migrations"
  }
  
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}