package migrations

import(
    "github.com/zennon-sml/JOJORNAL/models"
    "gorm.io/gorm"
)

func FazerMigrations(bd *gorm.DB){
    bd.Debug().AutoMigrate(models.Artigo{})
    bd.Debug().AutoMigrate(models.Administrador{})
    bd.Debug().AutoMigrate(models.Usuario{})
    bd.Debug().AutoMigrate(models.Comentarios{})
}
