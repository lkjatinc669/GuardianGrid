package scanfolder

// PUBLIC
func ScanFolder(root string) (map[string]interface{}, error) {
	files := make([]map[string]interface{}, 0)

	err := walkFolder(root, &files)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"folder_scan": map[string]interface{}{
			"path":  root,
			"files": files,
			"count": len(files),
		},
	}, nil
}
