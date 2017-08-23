package store

type SqlStorer interface {
	Select(map[string]interface{}) StoreChannel
	Create() StoreChannel
	Update() StoreChannel
	Delete() StoreChannel
}
