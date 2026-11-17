// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package utils

import (
	"net/url"
	"strconv"
)

// NextOffset returns the offset of the next page from a pagination link, or 0 when there is none.
func NextOffset(href *string) int64 {
	if href == nil {
		return 0
	}
	link, err := url.Parse(*href)
	if err != nil {
		return 0
	}
	query := link.Query()
	value := query.Get("offset")
	if value == "" {
		value = query.Get("start")
	}
	offset, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return offset
}

// ListAll fetches the pages of a list call and returns the last page as a field map, whose `itemsKey`
// holds the items of all fetched pages. `offset` is the `Offset` field of the call options: the first
// page starts there and it is moved to each next page. Only the first page is fetched when the user
// set `limit`, or when `offset` is nil because the call has no paging.
func ListAll(c *Call, itemsKey string, offset **int64, fetch func() (any, error)) (map[string]any, error) {
	singlePage := offset == nil || ForRead(c.Data, Limit)
	var current int64
	if offset != nil && *offset != nil {
		current = **offset
	}
	var all []any
	for {
		page, err := fetch()
		if err != nil {
			return nil, err
		}
		fields, err := FieldMap(page)
		if err != nil {
			return nil, err
		}
		items, _ := fields[itemsKey].([]any)
		all = append(all, items...)
		fields[itemsKey] = all

		next := nextOffset(fields)
		// Stop when the next link is missing or does not move forward, to avoid looping forever.
		if singlePage || next == 0 || next <= current {
			return fields, nil
		}
		current = next
		*offset = &next
	}
}

func nextOffset(page map[string]any) int64 {
	link, _ := page[Next].(map[string]any)
	href, ok := link[Href].(string)
	if !ok {
		return 0
	}
	return NextOffset(&href)
}
