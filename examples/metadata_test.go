package examples

import (
	"testing"

	tms "github.com/testit-tms/adapters-go"
)

func TestMetadata_without_metadata_success(t *testing.T) {
	tms.Test(t, tms.TestMetadata{
		WorkItemIds: []string{"662"},
	},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_without_metadata_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			WorkItemIds: []string{"664"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_external_id_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with external id success",
			ExternalId:  "with_external_id_success",
			WorkItemIds: []string{"666"},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_external_id_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with external id failed",
			ExternalId:  "with_external_id_failed",
			WorkItemIds: []string{"668"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_title_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with title success",
			Title:       "with title",
			WorkItemIds: []string{"670"},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_title_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with title failed",
			Title:       "with title",
			WorkItemIds: []string{"672"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_work_item_ids_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with work item ids success",
			WorkItemIds: []string{"674"},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_work_item_ids_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with work item ids failed",
			WorkItemIds: []string{"676"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_description_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with description success",
			Description: "with description",
			WorkItemIds: []string{"678"},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_description_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with description failed",
			Description: "with description",
			WorkItemIds: []string{"680"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_labels_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with labels success",
			Labels:      []string{"label01", "label02"},
			WorkItemIds: []string{"682"},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_labels_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with labels failed",
			Labels:      []string{"label01", "label02"},
			WorkItemIds: []string{"684"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_class_name_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with class name success",
			ClassName:   "CustomClassName",
			WorkItemIds: []string{"686"},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_class_name_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with class name failed",
			ClassName:   "CustomClassName",
			WorkItemIds: []string{"688"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_name_space_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with name space success",
			NameSpace:   "CustomNameSpace",
			WorkItemIds: []string{"690"},
		},
		func() {
			tms.True(t, true)
		})

}

func TestMetadata_with_name_space_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with name space failed",
			NameSpace:   "CustomNameSpace",
			WorkItemIds: []string{"692"},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_links_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with links success",
			WorkItemIds: []string{"694"},
			Links: []tms.Link{
				{
					Url: "http://google.com",
				},
				{
					Url:         "http://google.com",
					Title:       "Google",
					Description: "Google search engine",
					LinkType:    tms.LINKTYPE_RELATED,
				},
			},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_links_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with links failed",
			WorkItemIds: []string{"696"},
			Links: []tms.Link{
				{
					Url: "http://google.com",
				},
				{
					Url:         "http://google.com",
					Title:       "Google",
					Description: "Google search engine",
					LinkType:    tms.LINKTYPE_RELATED,
				},
			},
		},
		func() {
			tms.True(t, false)
		})
}

func TestMetadata_with_all_success(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with all success",
			ExternalId:  "with_all_success",
			Title:       "with title",
			WorkItemIds: []string{"698"},
			Description: "with description",
			Labels:      []string{"label01", "label02"},
			ClassName:   "CustomClassName",
			NameSpace:   "CustomNameSpace",
			Links: []tms.Link{
				{
					Url: "http://google.com",
				},
				{
					Url:         "http://google.com",
					Title:       "Google",
					Description: "Google search engine",
					LinkType:    tms.LINKTYPE_RELATED,
				},
			},
		},
		func() {
			tms.True(t, true)
		})
}

func TestMetadata_with_all_failed(t *testing.T) {
	tms.Test(t,
		tms.TestMetadata{
			DisplayName: "with all failed",
			ExternalId:  "with_all_failed",
			Title:       "with title",
			WorkItemIds: []string{"700"},
			Description: "with description",
			Labels:      []string{"label01", "label02"},
			ClassName:   "CustomClassName",
			NameSpace:   "CustomNameSpace",
			Links: []tms.Link{
				{
					Url: "http://google.com",
				},
				{
					Url:         "http://google.com",
					Title:       "Google",
					Description: "Google search engine",
					LinkType:    tms.LINKTYPE_RELATED,
				},
			},
		},
		func() {
			tms.True(t, false)
		})
}
