package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "nq3k718vs1k7zcb",
				"created": "2023-08-21 21:44:07.981Z",
				"updated": "2023-08-21 21:44:07.981Z",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "npiueszc",
						"name": "avatar",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/png",
								"image/jpeg",
								"image/gif",
								"image/webp",
								"image/svg+xml"
							],
							"thumbs": null,
							"protected": false
						}
					}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
					"allowEmailAuth": true,
					"allowOAuth2Auth": true,
					"allowUsernameAuth": true,
					"exceptEmailDomains": null,
					"manageRule": null,
					"minPasswordLength": 8,
					"onlyEmailDomains": null,
					"requireEmail": true
				}
			},
			{
				"id": "8nghjxl8fzma74n",
				"created": "2023-08-21 21:44:07.983Z",
				"updated": "2023-08-21 21:44:07.983Z",
				"name": "saves",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "00vwpc0r",
						"name": "name",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": 5,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "rfkop8pb",
						"name": "description",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "a4fwlk8y",
						"name": "savefile",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"application/zip"
							],
							"thumbs": [],
							"protected": true
						}
					},
					{
						"system": false,
						"id": "axc3uc5r",
						"name": "author",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "_pb_users_auth_",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					},
					{
						"system": false,
						"id": "9c98fpih",
						"name": "isPrivate",
						"type": "bool",
						"required": false,
						"unique": false,
						"options": {}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_KN68WeV` + "`" + ` ON ` + "`" + `saves` + "`" + ` (` + "`" + `name` + "`" + `)",
					"CREATE INDEX ` + "`" + `idx_mfeNrom` + "`" + ` ON ` + "`" + `saves` + "`" + ` (` + "`" + `created` + "`" + `)"
				],
				"listRule": "isPrivate = false || (@request.auth.id != \"\" && author = @request.auth.id)",
				"viewRule": "isPrivate = false || (@request.auth.id != \"\" && author = @request.auth.id)",
				"createRule": "@request.auth.id != \"\" && author = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && author = @request.auth.id",
				"deleteRule": "@request.auth.id != \"\" && author = @request.auth.id",
				"options": {}
			},
			{
				"id": "c10wvx4sffp9pti",
				"created": "2023-08-21 21:47:55.707Z",
				"updated": "2023-08-21 21:48:42.205Z",
				"name": "mods",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "pugc0lnc",
						"name": "name",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": 10,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "zp7xbzs3",
						"name": "description",
						"type": "text",
						"required": true,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "eenif6rw",
						"name": "screenshots",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 10,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/png",
								"image/jpeg",
								"image/gif",
								"image/webp"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "z9qwxhos",
						"name": "icon",
						"type": "file",
						"required": false,
						"unique": false,
						"options": {
							"maxSelect": 1,
							"maxSize": 5242880,
							"mimeTypes": [
								"image/jpeg",
								"image/gif",
								"image/png",
								"image/svg+xml",
								"image/webp"
							],
							"thumbs": [],
							"protected": false
						}
					},
					{
						"system": false,
						"id": "7zlypxfi",
						"name": "author",
						"type": "relation",
						"required": true,
						"unique": false,
						"options": {
							"collectionId": "nq3k718vs1k7zcb",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_YuR3e56` + "`" + ` ON ` + "`" + `mods` + "`" + ` (` + "`" + `name` + "`" + `)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@request.auth.id != \"\" && author = @request.auth.id",
				"updateRule": "@request.auth.id != \"\" && author = @request.auth.id",
				"deleteRule": "@request.auth.id != \"\" && author = @request.auth.id",
				"options": {}
			},
			{
				"id": "vex2jdq00q3qk4j",
				"created": "2023-08-21 21:53:05.615Z",
				"updated": "2023-08-21 21:53:05.615Z",
				"name": "mod_versions",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "zatotdnj",
						"name": "version",
						"type": "number",
						"required": true,
						"unique": false,
						"options": {
							"min": 0,
							"max": null
						}
					},
					{
						"system": false,
						"id": "zrror3vd",
						"name": "changelog",
						"type": "text",
						"required": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "yw66g2pa",
						"name": "mod",
						"type": "relation",
						"required": false,
						"unique": false,
						"options": {
							"collectionId": "c10wvx4sffp9pti",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": []
						}
					}
				],
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_4Vkm0W7` + "`" + ` ON ` + "`" + `mod_versions` + "`" + ` (` + "`" + `mod` + "`" + `)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_sjv8I8w` + "`" + ` ON ` + "`" + `mod_versions` + "`" + ` (\n  ` + "`" + `version` + "`" + `,\n  ` + "`" + `mod` + "`" + `\n)"
				],
				"listRule": "",
				"viewRule": "",
				"createRule": "@collection.mods.id ?= mod.id && @collection.mods.author ?= @request.auth.id",
				"updateRule": "@collection.mods.author ?= @request.auth.id",
				"deleteRule": "@collection.mods.author ?= @request.auth.id",
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
