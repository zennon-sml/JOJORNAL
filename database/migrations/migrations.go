package migrations

import(
    "github.com/zennon-sml/JOJORNAL/models"
    "gorm.io/gorm"
)

func fazerMigrations(bd *gorm.DB){
    db.AutoMigrate(models.Artigo)
}
