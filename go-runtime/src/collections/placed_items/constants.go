package collections_placed_items

const (
	COLLECTION_NAME = "PlacedItems"
)

const (
	STORAGE_INDEX_BY_REFERENCE_KEY = "PlacedItemsStorageIndexByReferenceKey"
	STORAGE_INDEX_BY_FILTERS_1     = "PlacedItemsStorageIndexByFilters1"
	STORAGE_INDEX_BY_FILTERS_2     = "PlacedItemsStorageIndexByFilters2"
	//tiles owned by user, with referenceKey
	STORAGE_INDEX_BY_FILTERS_3     = "PlacedItemsStorageIndexByFilters3"
	STORAGE_INDEX_BY_INVENTORY_KEY = "PlacedItemsStorageIndexByInventoryKey"
)

const (
	TYPE_TILE     = 0
	TYPE_BUILDING = 1
	TYPE_ANIMAL   = 2
)

const (
	SIZE_0 = 0
	SIZE_1 = 1
	SIZE_2 = 1
	SIZE_3 = 1
)

const (
	CROP_CURRENT_STATE_NORMAL      = 0
	CROP_CURRENT_STATE_NEED_WATER  = 1
	CROP_CURRENT_STATE_IS_WEEDY    = 2
	CROP_CURRENT_STATE_IS_INFESTED = 3
)

const (
	ANIMAL_CURRENT_STATE_NORMAL = 0
	ANIMAL_CURRENT_STATE_HUNGRY = 1
	ANIMAL_CURRENT_STATE_SICK   = 2
)
