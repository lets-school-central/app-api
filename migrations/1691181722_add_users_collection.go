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
				Id: "nq3k718vs1k7zcb",
			},
			Name:   "users",
			Type:   "auth",
			System: false,
			Schema: schema.NewSchema(
				&schema.SchemaField{
					Id:       "npiueszc",
					Name:     "avatar",
					Type:     "file",
					System:   false,
					Required: false,
					Options: map[string]any{
						"maxSelect": 1,
						"maxSize":   5242880,
						"mimeTypes": []string{
							"image/png",
							"image/jpeg",
							"image/gif",
							"image/webp",
							"image/svg+xml",
						},
						"protected": false,
					},
				},
			),
			Indexes:    types.JsonArray[string]{},
			ListRule:   utils.String("id = @request.auth.id"),
			ViewRule:   utils.String("id = @request.auth.id"),
			CreateRule: utils.String(""),
			UpdateRule: utils.String("id = @request.auth.id"),
			DeleteRule: utils.String("id = @request.auth.id"),
			Options: types.JsonMap{
				"allowEmailAuth":     true,
				"allowOAuth2Auth":    true,
				"allowUsernameAuth":  true,
				"exceptEmailDomains": nil,
				"manageRule":         nil,
				"minPasswordLength":  8,
				"onlyEmailDomains":   nil,
				"requireEmail":       true,
			},
		}

		return daos.New(db).ImportCollections([]*models.Collection{&collection}, true, nil)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)
		collection, err := dao.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		if collection != nil {
			return dao.DeleteCollection(collection)
		}
		return nil
	})
}
