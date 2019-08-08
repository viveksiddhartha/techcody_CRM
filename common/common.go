package common

import (
	"honnef.co/go/js/dom"

	"SV_CRM/common/datastore"
)

type Env struct {
	DB             datastore.Datastore
	Window         dom.Window
	Document       dom.Document
	PrimaryContent dom.Element
}
