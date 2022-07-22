package migrations

import(
    "github.com/zennon-sml/JOJORNAL/models"
    "gorm.io/gorm"
)

func FazerMigrations(bd *gorm.DB){
    bd.AutoMigrate(models.Artigo{})
}
