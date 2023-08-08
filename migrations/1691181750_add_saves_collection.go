package migrations

import (
	"app.lets-school-central/api/utils"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		collection := models.Collection{
			BaseModel: models.BaseModel{
				Id: "8nghjxl8fzma74n",
			},
			Name:   "saves",
			Type:   "base",
			System: false,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Id:   "00vwpc0r",
					Name: "name",
					Type: "text",
					Options: map[string]interface{}{
						"min":     5,
						"max":     nil,
						"pattern": "",
					},
				},
				&schema.SchemaField{
					Id:   "rfkop8pb",
					Name: "description",
					Type: "text",
					Options: map[string]interface{}{
						"min":     nil,
						"max":     nil,
						"pattern": "",
					},
				},
				&schema.SchemaField{
					Id:   "a4fwlk8y",
					Name: "savefile",
					Type: "file",
					Options: map[string]interface{}{
						"maxSelect":  1,
						"maxSize":    5242880,
						"mimeTypes":  []string{"application/zip"},
						"thumbs":     []string{},
						"protected":  true,
						"publicRead": true,
					},
				},
				&schema.SchemaField{
					Id:   "axc3uc5r",
					Name: "author",
					Type: "relation",
					Options: map[string]interface{}{
						"collectionId":  "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect":     nil,
						"maxSelect":     1,
						"displayFields": []string{},
					},
				},
				&schema.SchemaField{
					Id:      "9c98fpih",
					Name:    "isPrivate",
					Type:    "bool",
					Options: map[string]interface{}{},
				}),
			Indexes: types.JsonArray[string]{
				"CREATE UNIQUE INDEX `idx_KN68WeV` ON `saves` (`name`)",
				"CREATE INDEX `idx_mfeNrom` ON `saves` (`created`)",
			},
			ListRule:   utils.String("isPrivate = false || (@request.auth.id != \"\" && author = @request.auth.id)"),
			ViewRule:   utils.String("isPrivate = false || (@request.auth.id != \"\" && author = @request.auth.id)"),
			CreateRule: utils.String("@request.auth.id != \"\" && author = @request.auth.id"),
			UpdateRule: utils.String("@request.auth.id != \"\" && author = @request.auth.id"),
			DeleteRule: utils.String("@request.auth.id != \"\" && author = @request.auth.id"),
		}

		return daos.New(db).SaveCollection(&collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)
		collection, err := dao.FindCollectionByNameOrId("saves")
		if err != nil {
			return err
		}
		if collection != nil {
			return dao.DeleteCollection(collection)
		}
		return nil
	})
}
